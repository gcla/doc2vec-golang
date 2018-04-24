package main

import (
	"net/http"
	_ "net/http/pprof"

	"log"
	"os"

	"context"

	"github.com/gcla/doc2vec-golang/common"
	"github.com/gcla/doc2vec-golang/doc2vec"
)

func main() {
	fname := os.Args[1]

	// for pprof
	go func() {
		log.Println(http.ListenAndServe("localhost:16060", nil))
	}()

	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
	//d2v := doc2vec.NewDoc2Vec(useCbow = true, useHS = false, useNEG = true, windowSize = 5, dim = 50, iters = 50)
	d2v := doc2vec.NewDoc2Vec(false, false, true, 5, 50, 50)
	file, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	model := common.NewModelFileDataProvider(file)
	ctx := context.Background()
	d2v.Train(ctx, model)
	err = d2v.SaveModel("2.model")
	if err != nil {
		log.Fatal(err)
	}
}
