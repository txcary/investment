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
	Template	
	Pe float64		
}

func (obj *Cnn) GetUrl(id string) (url string) {
	url = `https://money.cnn.com/data/markets/`+id+`/`
	return
}

func (obj *Cnn) Process(doc *goquery.Document) error {
	tag := doc.Find("body")
	for _,item := range pattern {
		tag = tag.Find(item)
		//fmt.Println(item, tag.Text())
	}
	peStr := tag.Eq(2).Find("td").Eq(1).Text()
	fmt.Sscanf(peStr, "%f", &obj.Pe)
	return nil
}

func NewCnn() (obj *Cnn) {
	obj = new(Cnn)
	obj.Template.Init(obj)
	return 
}