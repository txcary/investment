package builder

import (
	"github.com/txcary/investment/db"
)

type DbBuilder struct {
	BuilderBase
	newObject func() interface{}
	database  *db.Db
}

func (obj *DbBuilder) get(id string) (object interface{}, ok bool) {
	object = obj.newObject()
	err := obj.database.Get(id, object)
	if err == nil {
		ok = true
	}
	return
}

func (obj *DbBuilder) put(id string, object interface{}) {
	obj.database.Put(id, object)
}

func (obj *DbBuilder) Build(id string) (object interface{}, ok bool) {
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

func (obj *DbBuilder) Init(decorator Builder, database *db.Db, newObject func() interface{}) {
	obj.decorator = decorator
	obj.newObject = newObject
	obj.database = database
}

func NewDbBuilder(decorator Builder, database *db.Db, newObject func() interface{}) *DbBuilder {
	obj := new(DbBuilder)
	obj.Init(decorator,database,newObject)
	return obj
}
