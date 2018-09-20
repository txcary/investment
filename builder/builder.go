package builder

import (
	"github.com/txcary/investment/utils"
)

type Decorator interface {
	Build(id string) (interface{}, bool)
	Update(id string, object interface{})
}

type Builder struct {
	decorator Decorator
}

func (obj *Builder) Update(id string, object interface{}) {
	if !utils.CheckNil(obj.decorator) {
		obj.decorator.Update(id, object)
	}
}

func (obj *Builder) Build(id string) (object interface{}, ok bool) {
	if !utils.CheckNil(obj.decorator) {
		object, ok = obj.decorator.Build(id)
		return
	}
	return nil, false
}

func New(decorator Decorator) *Builder {
	obj := new(Builder)
	obj.decorator = decorator
	return obj
}
