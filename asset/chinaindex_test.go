package asset
import (
	"fmt"	
)

func ExampleChinaindex() {
	obj := NewChinaIndex("159919")
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