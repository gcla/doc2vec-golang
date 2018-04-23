package doc2vec

import (
	"context"
	"sync"

	"github.com/gcla/doc2vec-golang/common"
	"github.com/gcla/doc2vec-golang/corpus"
	"github.com/gcla/doc2vec-golang/neuralnet"
)

//go:generate msgp

type SortItem struct {
	Idx int32
	Dis float64
}

type TSortItemSlice []*SortItem

type IDoc2Vec interface {
	Train(ctx context.Context, model common.IModelDataProvider)
	GetCorpus() corpus.ICorpus
	GetNeuralNet() neuralnet.INeuralNet
	SaveModel(fname string) (err error)
	LoadModel(fname string) (err error)
	Word2Words(word string)
	Word2Docs(word string)
	Sen2Words(ctx context.Context, content string, iters int)
	Sen2Docs(ctx context.Context, content string, iters int)
	Doc2Docs(docidx int)
	Doc2Words(docidx int)
	GetLikelihood4Doc(context string) (likelihood float64)
	GetLeaveOneOutKwds(ctx context.Context, content string, iters int)
	DocSimCal(content1 string, content2 string) (dis float64)
}

type TDoc2VecImpl struct {
	Dim              int
	UseCbow          bool //true:Continuous Bag-of-Word Model false:skip-gram
	WindowSize       int  //cbow model的窗口大小
	UseHS            bool
	UseNEG           bool //UseHS / UseNEG两种求解优化算法必须选一个 也可以两种算法都选 详见google word2vec源代码
	Negative         int  //负采样词的个数
	StartAlpha       float64
	Iters            int
	TrainedWords     int
	Corpus           corpus.ICorpus
	NN               neuralnet.INeuralNet
	NegSamplingTable []int32
	pool             *sync.Pool
}
