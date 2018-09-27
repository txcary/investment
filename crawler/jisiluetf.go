package crawler

import (
	"strconv"
	"strings"
	"github.com/bitly/go-simplejson"
)

var titleMap = map[string]string{
	"id":    "fund_id",
	"name":  "fund_nm",
	"index": "index_nm",
	"pe":    "pe",
	"pb":    "pb",
	"over":  "discount_rt",
	"total": "unit_total",
}

type Jisiluetf struct {
	TemplateJson
	dataList []map[string]string
	id string
	Name string
	Index string
	Over float64
	Pe float64
	Pb float64
	Total float64
	Roe float64
	IsValid bool
}

func (obj *Jisiluetf) update() {
	obj.IsValid = false 
	for i, _ := range obj.dataList {
		item := obj.dataList[i]
		if item["id"] != obj.id {
			continue
		}
		obj.IsValid = true

		obj.Name = item["name"]
		obj.Index = item["index"]
		obj.Over, _ = strconv.ParseFloat(item["over"],32)
		obj.Pe, _ = strconv.ParseFloat(item["pe"], 32)
		obj.Pb, _ = strconv.ParseFloat(item["pb"], 32)
		obj.Total, _ = strconv.ParseFloat(item["total"], 32)
		if obj.Pe >0 {
			obj.Roe = 100 * obj.Pb / obj.Pe
		} else {
			obj.Roe = 0
		}
		break
	}
}

func (obj *Jisiluetf) CrawlNeeded(id string)(bool) {
	obj.id = id
	if obj.dataList != nil {
		obj.update()
		return false
	} else {
		obj.dataList = make([]map[string]string,0)
		return true
	}
}

func (obj *Jisiluetf) GetUrl() (url string) {
	url = `https://www.jisilu.cn/jisiludata/etf.php`
	return
}

func (obj *Jisiluetf) Process(intf interface{}) error {
	json := intf.(*simplejson.Json)
	rows := json.Get("rows")
	rowsArray, err := rows.Array()
	if err!=nil {
		return err
	}
	for i, _ := range rowsArray {
		cell := rows.GetIndex(i).Get("cell")
		var item = make(map[string]string)
		for title, value := range titleMap {
			item[title] = cell.Get(value).MustString()
			item[title] = strings.Replace(item[title], "%", "", -1)
		}
		obj.dataList = append(obj.dataList, item)
	}

	obj.update()

	return nil
}

func NewJisiluetf() (obj *Jisiluetf) {
	obj = new(Jisiluetf)
	obj.SetStrategyToTemplate(obj)
	return 
}