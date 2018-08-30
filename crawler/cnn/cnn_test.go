package cnn 

import (
	"fmt"
	"github.com/txcary/investment/crawler"
)

var (
	ids = [...]string{
		"dow",
		"sandp",
		"nasdaq",
	}
)

func ExampleCnn() {
	cnnObj := New()
	crawlerObj := crawler.New(cnnObj)
	for _, id := range ids {
		crawlerObj.Process(id)
		fmt.Println(cnnObj.Pe)
	}
	//output:
	//TODO
}