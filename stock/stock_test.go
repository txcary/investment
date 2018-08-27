package stock

import (
	"fmt"
	"reflect"
)

func dumpData(data interface{}) {
	v := reflect.ValueOf(data)
	t := reflect.TypeOf(data)
	c := v.NumField()
	for i:=0; i<c; i++ {
		f := v.Field(i)
		switch f.Kind() {
			case reflect.Float64:
				fmt.Println(t.Field(i).Name,  f.Float())
				break
		}
	}
}

func ExampleBuilder() {
	builder := BuilderInstance()
	stock := builder.Build("00700")
	fmt.Println(stock.Id)
	fmt.Println(stock.Name)
	//dumpData(stock.Computed)
	//output:
	//00700
	//腾讯控股
}
