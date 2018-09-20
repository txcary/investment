package builder

import (
	"github.com/txcary/investment/db"
	"github.com/txcary/investment/utils"
)

type DbDecorator struct {
	decorator Decorator
	newObject func() interface{}
	database  *db.Db
}

func (obj *DbDecorator) get(id string) (object interface{}, ok bool) {
	ok = false
	object = obj.newObject()
	err := obj.database.Get(id, object)
	if err == nil {
		ok = true
	}
	return
}

func (obj *DbDecorator) put(id string, object interface{}) {
	obj.database.Put(id, object)
}

func (obj *DbDecorator) Update(id string, object interface{}) {
	obj.put(id, object)
	if !utils.CheckNil(obj.decorator) {
		obj.decorator.Update(id, object)
	}
}

func (obj *DbDecorator) Build(id string) (object interface{}, ok bool) {
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

func NewDbDecorator(decorator Decorator, database *db.Db, newObject func() interface{}) *DbDecorator {
	obj := new(DbDecorator)
	obj.decorator = decorator
	obj.newObject = newObject
	obj.database = database
	return obj
}
