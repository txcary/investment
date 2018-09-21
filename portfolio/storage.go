package portfolio

import (
	"github.com/txcary/investment/db"
)

type Storage struct {
	database *db.Db
}

func (obj *Storage) Put(key string, value []byte) error {
	return obj.database.PutBytes(key, value)
}

func (obj *Storage) Get(key string) ([]byte, error) {
	return obj.database.GetBytes(key)
}

func (obj *Storage) Init(database *db.Db) {
	obj.database = database
}

func NewStorage(database *db.Db) *Storage {
	obj := new(Storage)
	obj.Init(database)
	return obj
}
