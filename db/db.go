package db

import (
	"encoding/json"
	"github.com/syndtr/goleveldb/leveldb"
	"sync"
)

type Db struct {
	db *leveldb.DB	
	mutex sync.Mutex
}

func (obj *Db) Get (key string, value interface{}) (err error) {
	obj.mutex.Lock()
	bytes, err := obj.db.Get([]byte(key), nil)
	obj.mutex.Unlock()
	json.Unmarshal(bytes, value)
	return
}

func (obj *Db) Put (key string, value interface{}) (err error) {
	bytes, err := json.Marshal(value)
	obj.mutex.Lock()
	err = obj.db.Put([]byte(key), bytes, nil)
	obj.mutex.Unlock()
	return
}

func (obj *Db) Delete(key string) (err error) {
	obj.mutex.Lock()
	err = obj.db.Delete([]byte(key), nil)
	obj.mutex.Unlock()
	return
}

func (obj *Db) Close() {
	obj.mutex.Lock()
	obj.db.Close()
	obj.mutex.Unlock()
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