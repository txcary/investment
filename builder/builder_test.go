package builder

import (
	"fmt"
	"github.com/txcary/investment/db"
	"github.com/txcary/investment/utils"
)

type TestObject struct {
	Id    string
	Name  string
	Value string
}

type TestDecorator struct {
}

func newObj() interface{} {
	return new(TestObject)
}

func (obj *TestDecorator) Build(id string) (interface{}, bool) {
	object := newObj().(*TestObject)
	object.Id = id
	object.Name = "TestName"
	object.Value = "TestValue"
	return object, true
}

func (obj *TestDecorator) Update(id string, object interface{}) {
	fmt.Println(object.(*TestObject))
}

func ExampleNew() {
	dbObj := db.New(utils.Gopath() + utils.Slash() + "db/test.builder")
	testDecorator := new(TestDecorator)
	dbDecorator := NewDbDecorator(testDecorator, dbObj, newObj)
	memDecorator := NewMemDecorator(dbDecorator)
	obj := New(memDecorator)
	testObj, _ := obj.Build("00700")
	fmt.Println(testObj)

	testObj, _ = obj.Build("00700")
	fmt.Println(testObj)

	updateObj := newObj().(*TestObject)
	updateObj.Id = "00700"
	updateObj.Name = "UpdateName"
	updateObj.Value = "UpdateValue"
	obj.Update("00700", updateObj)

	testObj, _ = obj.Build("00700")
	fmt.Println(testObj)
	//output:
	//&{00700 TestName TestValue}
	//&{00700 TestName TestValue}
	//&{00700 UpdateName UpdateValue}
	//&{00700 UpdateName UpdateValue}
}
