package crawler

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"strconv"
	"strings"
)

type Danjuan struct {
	TemplateJson
	dataList []map[string]string
	id       string
	Name     string
	Pe       float64
	Pb       float64
	Roe      float64
	Div      float64
	IsValid  bool
}

var danjuanTitles = map[string]string{
	"id":   "index_code",
	"name": "name",
	"pe":   "pe",
	"pb":   "pb",
	"roe":  "roe",
	"div":  "yeild",
}

func (obj *Danjuan) update() {
	obj.IsValid = false
	for i, _ := range obj.dataList {
		item := obj.dataList[i]
		if item["id"] != obj.id {
			continue
		}
		obj.IsValid = true

		obj.Name = item["name"]
		obj.Roe, _ = strconv.ParseFloat(item["roe"], 32)
		obj.Pe, _ = strconv.ParseFloat(item["pe"], 32)
		obj.Pb, _ = strconv.ParseFloat(item["pb"], 32)
		obj.Div, _ = strconv.ParseFloat(item["div"], 32)
		break
	}
}

func (obj *Danjuan) CrawlNeeded(id string) bool {
	obj.id = id
	if obj.dataList != nil {
		obj.update()
		return false
	} else {
		obj.dataList = make([]map[string]string, 0)
		return true
	}
}

func (obj *Danjuan) GetUrl() string {
	return `https://danjuanapp.com/djapi/index_eva/dj`
}

func (obj *Danjuan) Process(intf interface{}) error {
	json := intf.(*simplejson.Json)
	rows := json.Get("data").Get("items")
	rowsArray, err := rows.Array()
	if err != nil {
		return err
	}
	for i, _ := range rowsArray {
		cell := rows.GetIndex(i)
		var item = make(map[string]string)
		for title, value := range danjuanTitles {
			item[title] = cell.Get(value).MustString()
			if len(item[title]) == 0 {
				asFloat := cell.Get(value).MustFloat64()
				item[title] = fmt.Sprintf("%.3f", asFloat)
			}
			item[title] = strings.Replace(item[title], "%", "", -1)
		}
		obj.dataList = append(obj.dataList, item)
	}
	obj.update()
	return nil
}

func NewDanjuan() (obj *Danjuan) {
	obj = new(Danjuan)
	obj.SetStrategyToTemplate(obj)
	return
}
