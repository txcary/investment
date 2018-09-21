package builder

import (
)

type PrototypeBuilder struct{
	BuilderBase
	newPrototype func(id string) interface{}
}

func (obj *PrototypeBuilder) Build(id string) (interface{}, bool) {
	return obj.newPrototype(id), true
}

func (obj *PrototypeBuilder)Init(newPrototype func(id string) interface{}) {
	obj.newPrototype = newPrototype	
}
func NewPrototypeBuilder(newPrototype func(id string) interface{}) *PrototypeBuilder {
	obj := new(PrototypeBuilder)
	obj.Init(newPrototype)
	return obj
}