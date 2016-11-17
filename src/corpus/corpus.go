package corpus

import (
	"bufio"
	_ "errors"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	VOCAB_HASH_SIZE int = 30000000
)

func (p *TCorpusImpl) GetWordIdx(word string) (idx int32, ok bool) {
	idx, ok = p.Word2Idx[word]
	return idx, ok
}

func (p *TCorpusImpl) GetWordsCnt() int {
	return p.WordsCnt
}

func (p *TCorpusImpl) createBinaryTree() {
	vocab_size := p.GetVocabCnt()
	size := vocab_size*2 + 1
	cnt := make([]int64, size, size)
	parent := make([]int32, size, size)
	binary := make([]bool, size, size)
	//p.GetAllWords 一定要降序排列 p.sortVocab
	for i, item := range p.GetAllWords() {
		cnt[i] = int64(item.Cnt)
	}
	for i := vocab_size; i < vocab_size*2; i++ {
		cnt[i] = math.MaxInt64
	}
	pos1 := vocab_size - 1 //初始化为min
	pos2 := vocab_size
	min1, min2 := 0, 0                  //最小和次小
	for i := 0; i < vocab_size-1; i++ { //vocab_size 个非叶结点
		if pos1 >= 0 {
			if cnt[pos1] < cnt[pos2] {
				min1 = pos1
				pos1--
			} else {
				min1 = pos2
				pos2++
			}
		} else {
			min1 = pos2
			pos2++
		}

		if pos1 >= 0 {
			if cnt[pos1] < cnt[pos2] {
				min2 = pos1
				pos1--
			} else {
				min2 = pos2
				pos2++
			}
		} else {
			min2 = pos2
			pos2++
		}
		cnt[vocab_size+i] = cnt[min1] + cnt[min2]
		parent[min1] = int32(vocab_size + i)
		parent[min2] = int32(vocab_size + i)
		binary[min2] = true
	}
	root := int32(vocab_size + vocab_size - 1 - 1)
	for i := 0; i < vocab_size; i++ {
		code := make([]bool, 0, 40)
		point := make([]int32, 0, 40)
		ii := int32(i)
		for true {
			point = append(point, int32(ii))
			code = append(code, binary[ii])
			p := parent[ii]
			if p == root {
				break
			}
			ii = p
		}
		//code point reverse
		reverse_code := make([]bool, 0, len(code))
		for j := len(code) - 1; j >= 0; j-- {
			reverse_code = append(reverse_code, code[j])
		}
		reverse_point := make([]int32, 0, len(point))
		//root node index
		reverse_point = append(reverse_point, root-int32(vocab_size)) // 加上root, 减去vocab_size后直接表示syn1中的下标了
		for j := len(point) - 1; j > 0; j-- {                         //NOTE  j != 0, j == 0是叶子节点、HS训练的时候不需要了, HS的Point只需要[root, leaf)
			syn1idx := point[j] - int32(vocab_size)
			reverse_point = append(reverse_point, syn1idx)
		}
		//(root->leaf]
		p.Words[i].Code = reverse_code
		p.Words[i].Point = reverse_point
	}
}

func (p TWordItemSlice) Len() int {
	return len(p)
}

func (p TWordItemSlice) Less(i, j int) bool {
	return p[i].Cnt < p[j].Cnt
}

func (p TWordItemSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *TCorpusImpl) GetDocCnt() int {
	return len(p.Doc2Idx)
}

func (p *TCorpusImpl) GetVocabCnt() int {
	return len(p.GetAllWords())
}

func (p *TCorpusImpl) GetAllWords() (words TWordItemSlice) {
	return p.Words
}

func (p *TCorpusImpl) GetDocWordsByDocid(id string) (doc []*TWordItem) {
	idx, ok := p.Doc2Idx[id]
	if ok {
		return p.GetDocWordsByIdx(int(idx))
	}
	return nil
}

func (p *TCorpusImpl) GetAllDocWordsIdx() [][]int32 {
	return p.Doc2WordsIdx
}

func (p *TCorpusImpl) GetAllDocWords() (docs [][]*TWordItem) {
	for i := 0; i < len(p.Doc2WordsIdx); i++ {
		docs = append(docs, p.GetDocWordsByIdx(i))
	}
	return docs
}

func (p *TCorpusImpl) GetWordItemByIdx(i int) (item *TWordItem) {
	if i < 0 || i >= len(p.Words) {
		log.Fatal("index out of range")
	}
	return &p.Words[i]
}

func (p *TCorpusImpl) GetDocWordsByIdx(i int) (doc []*TWordItem) {
	if i >= 0 && i < len(p.Doc2WordsIdx) {
		doc = make([]*TWordItem, len(p.Doc2WordsIdx[i]), len(p.Doc2WordsIdx[i]))
		for j, idx := range p.Doc2WordsIdx[i] {
			if idx >= 0 && idx < int32(len(p.Words)) {
				doc[j] = &p.Words[idx]
			}
		}
	}
	return doc
}

func (p *TCorpusImpl) reduceVocabulary() {
	var idx int32 = 0
	actual_size := len(p.Word2Idx)
	p.Word2Idx = map[string]int32{}
	for i, item := range p.Words {
		if i == actual_size {
			break
		}
		if item.Cnt > p.MinCnt {
			p.Words[idx] = item
			p.Word2Idx[item.Word] = idx
			idx++
		}
	}
	p.Words = p.Words[:len(p.Word2Idx)]
	p.MinReduce++
}

func (p *TCorpusImpl) addWord(word string) (err error) {
	idx, ok := p.Word2Idx[word]
	if !ok {
		item := TWordItem{Word: word}
		idx = int32(len(p.Words))
		p.Words = append(p.Words, item)
		p.Word2Idx[word] = idx
	}
	p.Words[idx].Cnt++
	return err
}

func (p *TCorpusImpl) loadAsWords(docid string, content string) int {
	items := strings.Split(content, " ")
	for _, word := range items {
		p.addWord(word)
		if len(p.Word2Idx) > int(0.7*float32(VOCAB_HASH_SIZE)) {
			p.reduceVocabulary()
		}
	}
	return len(items)
}

func (p *TCorpusImpl) loadAsDoc(docid string, content string) int {
	items := strings.Split(content, " ")
	wordsIdx := []int32{}
	for _, word := range items {
		idx, ok := p.Word2Idx[word]
		if ok {
			wordsIdx = append(wordsIdx, int32(idx))
		}
	}
	p.Doc2WordsIdx = append(p.Doc2WordsIdx, wordsIdx)
	p.Doc2Idx[docid] = int32(len(p.Doc2WordsIdx) - 1)
	return 1
}

func (p *TCorpusImpl) String() string {
	words_cnt := strconv.Itoa(len(p.Words))
	words_map_cnt := strconv.Itoa(len(p.Word2Idx))
	docs_cnt := strconv.Itoa(len(p.Doc2WordsIdx))
	return fmt.Sprintf("words_cnt:%v,words_map_cnt:%v,docs_cnt:%v\n", words_cnt, words_map_cnt, docs_cnt)
}

func (p *TCorpusImpl) buildVocabulary(fname string) (err error) {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Buffer([]byte{}, bufio.MaxScanTokenSize*10)
	train_words := 0
	batch := 0
	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, "\t")
		if len(items) != 2 {
			log.Printf("len(items)=%d\n", len(items))
			continue
		}
		docid, content := items[0], items[1]
		cnt := p.loadAsWords(docid, content)
		train_words += cnt
		batch += cnt
		if batch >= 10000000 {
			batch = 0
			fmt.Printf("%s train %d words\n", time.Now(), train_words)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	p.sortVocab()
	p.createBinaryTree()
	return err
}

func (p *TCorpusImpl) sortVocab() {
	//先排序
	// Words occuring less than min_count times will be discarded from the vocab
	p.Word2Idx = map[string]int32{}
	var cnt int32 = 0
	sort.Sort(sort.Reverse(p.Words))

	p.WordsCnt = 0
	for _, item := range p.Words {
		if item.Cnt > p.MinCnt {
			p.Words[cnt] = item
			p.Word2Idx[item.Word] = cnt
			p.WordsCnt += int(cnt)
			cnt++
		}
	}
	p.Words = p.Words[:cnt]
}

func (p *TCorpusImpl) buildDocument(fname string) (err error) {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Buffer([]byte{}, bufio.MaxScanTokenSize*10)
	train_docs := 0
	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, "\t")
		if len(items) != 2 {
			log.Printf("len(items)=%d\n", len(items))
			continue
		}
		docid, content := items[0], items[1]
		cnt := p.loadAsDoc(docid, content)
		train_docs += cnt
		if train_docs%10000 == 0 {
			fmt.Printf("%s train %d docs\n", time.Now(), train_docs)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return err
}

func (p *TCorpusImpl) Build(fname string) (err error) {
	err = p.buildVocabulary(fname)
	if err != nil {
		return err
	}
	return p.buildDocument(fname)
}

func NewCorpus() ICorpus {
	self := &TCorpusImpl{
		Word2Idx: make(map[string]int32),
		Doc2Idx:  make(map[string]int32),
		MinCnt:  10}
	return ICorpus(self)
}