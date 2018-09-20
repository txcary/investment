package builder

import (
	"encoding/json"
	"github.com/txcary/investment/utils"
	"sync"
)

type MemDecorator struct {
	objMap    sync.Map
	decorator Decorator
}

func (obj *MemDecorator) get(id string) (object interface{}, ok bool) {
	object, ok = obj.objMap.Load(id)
	return
}

func (obj *MemDecorator) put(id string, object interface{}) {
	existObj, ok := obj.get(id)
	if !ok {
		obj.objMap.Store(id, object)
	} else {
		objectBytes, err := json.Marshal(object)
		if err == nil {
			json.Unmarshal(objectBytes, existObj)
		}
	}
}

func (obj *MemDecorator) Update(id string, object interface{}) {
	obj.put(id, object)
	if !utils.CheckNil(obj.decorator) {
		obj.decorator.Update(id, object)
	}
}

func (obj *MemDecorator) Build(id string) (object interface{}, ok bool) {
	object, ok = obj.get(id)
	if ok {
		return
	}

	if !utils.CheckNil(obj.decorator) {
		object, ok = obj.decorator.Build(id)
		if ok {
			obj.put(id, object)
			return
		}
	}
	return nil, false
}

func NewMemDecorator(decorator Decorator) *MemDecorator {
	obj := new(MemDecorator)
	obj.decorator = decorator
	return obj
}
