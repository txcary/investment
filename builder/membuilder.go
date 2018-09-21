package builder

import (
	"encoding/json"
	"sync"
)

type MemBuilder struct {
	BuilderBase
	objMap    sync.Map
}

func (obj *MemBuilder) get(id string) (object interface{}, ok bool) {
	object, ok = obj.objMap.Load(id)
	return
}

func (obj *MemBuilder) put(id string, object interface{}) {
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

func (obj *MemBuilder) Build(id string) (object interface{}, ok bool) {
	object, ok = obj.get(id)
	if ok {
		return
	}

	object, ok = obj.BuilderBase.Build(id)
	if ok {
		obj.put(id, object)
	}
	return
}

func (obj *MemBuilder) Init(decorator Builder) {
	obj.decorator = decorator
}

func NewMemBuilder(decorator Builder) *MemBuilder {
	obj := new(MemBuilder)
	obj.Init(decorator)
	return obj
}
