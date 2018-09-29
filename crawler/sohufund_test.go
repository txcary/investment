package crawler

import (
	"fmt"
)

var (
	sohuids = [...]string{
		"159919",
	}
)

func ExampleSohufund() {
	obj := NewSohufund()
	for _, id := range sohuids {
		obj.Crawl(id)
		if obj.Trend > 0 {
			fmt.Println("OK")
		}
	}
	//output:
	//OK
}
