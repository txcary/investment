package portfolio
import (
)

type Portfolio struct {
	userName string	
	jsonString string
}

func (obj *Portfolio) Get(jsonString string) {
	return obj.jsonString
}

func (obj *Portfolio) Put(jsonString string) {
	obj.jsonString = jsonString
}

func (obj *Portfolio) Init(userName string) {
	obj.userName = userName
}

func New(userName string) *Portfolio {
	obj := new(Portfolio)
	obj.Init(userName)
	return obj
}