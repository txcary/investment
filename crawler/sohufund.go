package crawler

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

type Sohufund struct {
	TemplateHttp
	id      string
	Current float64
	Average float64
	Trend   float64
}

func (obj *Sohufund) trToFloat(trSel *goquery.Selection) (value float64) {
	valueStr := trSel.Find("td").Eq(2).Text()
	fmt.Sscanf(valueStr, "%f", &value)
	return
}

func (obj *Sohufund) CrawlNeeded(id string) bool {
	obj.id = id
	return true
}

func (obj *Sohufund) GetUrl() (url string) {
	url = `http://q.fund.sohu.com/q/vl.php?code=` + obj.id
	return
}

func (obj *Sohufund) Process(intf interface{}) error {
	doc := intf.(*goquery.Document)
	var total float64
	var count float64
	doc.Find("table").Eq(1).Find("tr").Each(func(trIdx int, trSel *goquery.Selection) {
		if trIdx >= 100 {
			return
		}
		if trIdx == 0 {
			return
		}
		if trIdx == 1 {
			obj.Current = obj.trToFloat(trSel)
		}
		total += obj.trToFloat(trSel)
		count++
	})
	if count > 0 {
		obj.Average = total / count
		obj.Trend = obj.Current / obj.Average
	}
	return nil
}

func NewSohufund() (obj *Sohufund) {
	obj = new(Sohufund)
	obj.SetStrategyToTemplate(obj)
	return
}
