package db

import (
	"encoding/json"
	"github.com/syndtr/goleveldb/leveldb"
)

type Db struct {
	db *leveldb.DB	
}

func (obj *Db) Get (key string, value interface{}) (err error) {
	bytes, err := obj.db.Get([]byte(key), nil)
	json.Unmarshal(bytes, value)
	return
}

func (obj *Db) Put (key string, value interface{}) (err error) {
	bytes, err := json.Marshal(value)
	err = obj.db.Put([]byte(key), bytes, nil)
	return
}

func (obj *Db) Delete(key string) (err error) {
	err = obj.db.Delete([]byte(key), nil)
	return
}

func (obj *Db) Close() {
	obj.db.Close()
}

func (obj *Db) Init(fileName string) {
	var err error
	obj.db, err = leveldb.OpenFile(fileName, nil)	
	if err != nil {
		panic("leveldb.OpenFile() fail!")
	}
}

func New(fileName string) *Db {
	obj := new(Db)
	obj.Init(fileName)
	return obj
} 