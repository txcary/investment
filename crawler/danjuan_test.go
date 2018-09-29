package crawler

import (
	"fmt"
)

func ExampleDanjuan() {
	obj := NewDanjuan()
	obj.Crawl("HKHSI")
	fmt.Println(obj.Name)
	if obj.Pe > 0 {
		fmt.Println("OK")
	}
	//output:
	//恒生指数
	//OK
}
