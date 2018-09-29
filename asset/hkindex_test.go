package asset

import (
	"fmt"
)

func ExampleHkindex() {
	obj := NewHkIndex("159920")
	yield := obj.YieldIndicator()
	trend := obj.TrendIndicator()
	if yield > 0 {
		fmt.Println("ok")
	}
	if trend > 0 {
		fmt.Println("ok")
	}
	//output:
	//ok
	//ok
}
