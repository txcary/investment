package builder

import (
	"github.com/txcary/goutils"
)

type Builder interface {
	Build(id string) (interface{}, bool)
}

type BuilderBase struct {
	decorator Builder
}

func (obj *BuilderBase) Build(id string) (object interface{}, ok bool) {
	if !utils.CheckNil(obj.decorator) {
		object, ok = obj.decorator.Build(id)
		return
	}
	return nil, false
}

func (obj *BuilderBase)Init(decorator Builder) {
	obj.decorator = decorator
}

func New(decorator Builder) *BuilderBase {
	obj := new(BuilderBase)
	obj.Init(decorator)
	return obj
}
