package crawler

import (
	"fmt"
)

var (
	cnnids = [...]string{
		"dow",
		"sandp",
		"nasdaq",
	}
)

func ExampleCnn() {
	cnnObj := NewCnn()
	for _, id := range cnnids {
		cnnObj.Crawl(id)
		fmt.Println(cnnObj.Pe)
	}
	//output:
	//TODO
}