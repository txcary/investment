package crawler

import (
	"fmt"
)

func ExampleJisiluetf() {
	obj := NewJisiluetf()
	obj.Crawl("513600")
	obj.Crawl("159919")
	fmt.Println(obj.Name)
	//output:
	//300ETF
}
