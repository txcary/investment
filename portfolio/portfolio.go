package portfolio

import (
	"sync"
	"encoding/json"
	"github.com/txcary/securejson"
	"github.com/txcary/investment/db"
	"github.com/txcary/investment/config"
)

type Portfolio struct {
	database *db.Db
	secjson  *securejson.SecureJson
}

var (
	obj   *Portfolio
	mutex sync.Mutex
)

func (obj *Portfolio) DelJson(inputJson []byte) (err error) {
	ok, err := obj.secjson.VerifyJson(inputJson)
	if !ok || err!=nil {
		return err
	}
	var data securejson.Json
	err = json.Unmarshal(inputJson, &data)
	if err != nil {
		return
	}
	err = obj.database.Delete(data.UserName)
	return
}

func (obj *Portfolio) GetJson(inputJson []byte) (outputJson []byte, err error) {
	outputJson, err = obj.secjson.GetJson(inputJson)
	return
}

func (obj *Portfolio) PutJson(inputJson []byte) (err error) {
	err = obj.secjson.PutJson(inputJson)
	return
}

func (obj *Portfolio) initDb() {
	dbpath := config.Instance().GetString("portfolio", "db")
	obj.database = db.New(dbpath)
}

func (obj *Portfolio) initSecJson(database *db.Db) {
	storage := NewStorage(database)
	obj.secjson = securejson.New(storage)
}

func Instance() *Portfolio {
	mutex.Lock()
	if obj == nil {
		obj = new(Portfolio)
		obj.initDb()
		obj.initSecJson(obj.database)
	}
	mutex.Unlock()
	return obj
}
