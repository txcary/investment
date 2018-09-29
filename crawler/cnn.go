package crawler

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

var (
	pattern = [...]string{
		"div#cnnBody",
		"div.mod-quoteinfo",
		"div#wsod_quoteDetail",
		"div#wsod_quoteRight",
		"table",
		"tbody",
		"tr",
	}
)

type Cnn struct {
	TemplateHttp
	id string
	Pe float64
}

func (obj *Cnn) CrawlNeeded(id string) bool {
	obj.id = id
	return true
}

func (obj *Cnn) GetUrl() (url string) {
	url = `https://money.cnn.com/data/markets/` + obj.id + `/`
	return
}

func (obj *Cnn) Process(intf interface{}) error {
	doc := intf.(*goquery.Document)
	tag := doc.Find("body")
	for _, item := range pattern {
		tag = tag.Find(item)
		//fmt.Println(item, tag.Text())
	}
	peStr := tag.Eq(2).Find("td").Eq(1).Text()
	fmt.Sscanf(peStr, "%f", &obj.Pe)
	return nil
}

func NewCnn() (obj *Cnn) {
	obj = new(Cnn)
	obj.SetStrategyToTemplate(obj)
	return
}
