package corpus

import (
	"github.com/gcla/doc2vec-golang/common"
	"github.com/tinylib/msgp/msgp"
)

//go:generate msgp
type TWordItem struct {
	Cnt   int32   //term frequency
	Point []int32 //Huffman tree(n leaf + n inner node, include root) path. [root, leaf), node index
	Code  []bool  //Huffman code. (root, leaf], 0/1 codes
	Word  string  //word desc
}
type TWordItemSlice []TWordItem

type ICorpus interface {
	Build(model common.IModelDataProvider) (err error)
	GetVocabCnt() int //排重后的词库大小 Translate: "Word size after weight"
	GetDocCnt() int   //doc个数, 按docid排重
	GetWordsCnt() int //排重前的词数 Translate: "The number of words before the weight"
	GetWordIdx(word string) (idx int32, ok bool)
	GetWordItemByIdx(i int) (item *TWordItem)
	GetAllWords() (words TWordItemSlice)
	GetAllDocWordsIdx() [][]int32
	GetAllDocWords() (doc [][]*TWordItem)
	GetDocWordsByDocid(id string) (doc []*TWordItem)
	GetDocWordsByIdx(i int) (doc []*TWordItem)
	Transform(content string) (wordsidx []int32)
	GetDocids() []string // all documents as read in
	msgp.Encodable
	msgp.Decodable
	msgp.Marshaler
	msgp.Unmarshaler
	msgp.Sizer
}

type TCorpusImpl struct {
	Words        TWordItemSlice   //Vocab
	Word2Idx     map[string]int32 //word -> words中的下标
	Doc2WordsIdx [][]int32        //
	Doc2Idx      map[string]int32 //docid -> Doc2WordsIdx中的下表
	Docids       []string         // document words, in order, when read in
	MinReduce    int32
	MinCnt       int32
	WordsCnt     int //未排重的词数 Translate: "Unranked words"
}
