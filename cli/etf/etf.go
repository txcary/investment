package main

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	//"math"
	"net/http"
	"strconv"
	"strings"
	"github.com/txcary/investment/cli/table"
)

func main() {
	table := table.New()
	table.AddTitle("ID", "id")
	table.AddTitle("Name", "name")
	table.AddTitle("Index", "index")
	table.AddTitle("Over", "over")
	table.AddTitle("PE", "pe")
	table.AddTitle("PB", "pb")
	table.AddTitle("ROE", "roe")
	table.AddTitle("Return", "ret")
	table.AddTitle("Total", "total")

	var titleMap = map[string]string{
		"id":    "fund_id",
		"name":  "fund_nm",
		"index": "index_nm",
		"pe":    "pe",
		"pb":    "pb",
		"over":  "discount_rt",
		"total": "unit_total",
	}
	var dataList = make([]map[string]string, 0)

	resp, err := http.Get("https://www.jisilu.cn/jisiludata/etf.php")
	if err != nil {
		fmt.Println("http.Get err")
	}
	body, err := ioutil.ReadAll(resp.Body)
	json, err := simplejson.NewJson([]byte(body))
	rows := json.Get("rows")
	rowsArray, err := rows.Array()
	for i, _ := range rowsArray {
		cell := rows.GetIndex(i).Get("cell")
		var item = make(map[string]string)
		for title, value := range titleMap {
			item[title] = cell.Get(value).MustString()
			item[title] = strings.Replace(item[title], "%", "", -1)
		}
		dataList = append(dataList, item)
	}

	for i, _ := range dataList {
		total, _ := strconv.ParseFloat(dataList[i]["total"], 32)
		item := dataList[i]
		if total >= 5 {
			/*
			   y, _ := strconv.ParseFloat(dataList[i]["over"],32)
			   y, _ := strconv.ParseFloat(dataList[i]["profitBuy"],32)
			   y, _ := strconv.ParseFloat(dataList[i]["profitSub"],32)
			   y, _ := strconv.ParseFloat(dataList[i]["profit7d"],32)
			   y, _ := strconv.ParseFloat(dataList[i]["profit1y"],32)
			*/
			pe, err := strconv.ParseFloat(item["pe"], 32)
			if err != nil {
				continue
			}
			pb, err := strconv.ParseFloat(item["pb"], 32)
			if err != nil {
				continue
			}
			roe := 100 * pb / pe
			ret := 100/pe
			if ret > roe {
			  ret = roe
			}
			if ret < 5 {
				continue
			}
			line := map[string]string{
				"id":    item["id"],
				"name":  item["name"],
				"index": item["index"],
				"over":  item["over"],
				"pe":    item["pe"],
				"pb":    item["pb"],
				"roe":   fmt.Sprintf("%.2f", 100*pb/pe),
				"ret":   fmt.Sprintf("%.2f", ret),
				"total": item["total"],
			}
			table.AddLine(line)
		}
	}
	table.Sort("ret", false)
	table.ShowAll()
}

