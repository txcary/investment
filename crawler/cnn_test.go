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
		if cnnObj.Pe != 0 {
			fmt.Println("OK")
		}
	}
	//output:
	//OK
	//OK
	//OK
}