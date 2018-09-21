package builder

import (
	"fmt"
	"github.com/txcary/investment/db"
	"github.com/txcary/investment/utils"
)

const (
	testId string = "00700"
)

type TestObject struct {
	Id    string
	Name  string
	Value string
}

func newObj() interface{} {
	return new(TestObject)
}

func testPrototype(id string) interface{} {
	object := newObj().(*TestObject)
	object.Id = id
	object.Name = "TestName"
	object.Value = "TestValue"
	return object
}

func ExampleNew() {
	dbObj := db.New(utils.Gopath() + utils.Slash() + "db/test.builder")
	dbObj.Delete(testId)
	testBuilder := NewPrototypeBuilder(testPrototype)
	dbBuilder := NewDbBuilder(testBuilder, dbObj, newObj)
	memBuilder := NewMemBuilder(dbBuilder)
	obj := New(memBuilder)
	testObj, _ := obj.Build(testId)
	fmt.Println(testObj)

	testObj, _ = obj.Build(testId)
	fmt.Println(testObj)

	//output:
	//&{00700 TestName TestValue}
	//&{00700 TestName TestValue}
}
