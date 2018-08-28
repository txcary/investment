package db
import (
	"fmt"
	"github.com/txcary/investment/common/utils"
)

type Data struct {
	Name string
	Value float64
}

func ExampleNew() {
	obj := New(utils.Gopath()+utils.Slash()+"db"+utils.Slash()+"test.db")
	dataIn := new(Data)
	dataIn.Name = "MyName"
	dataIn.Value = 3.14159
	dataOut := new(Data)
	obj.Delete("TestData")
	obj.Put("TestData", dataIn)
	obj.Get("TestData", dataOut)
	fmt.Println(dataOut)
	//output:
	//&{MyName 3.14159}
}