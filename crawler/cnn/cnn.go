package cnn

import (
	"net/http"
	"io/ioutil"
	"bytes"
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
	Pe float64		
}

func (obj *Cnn) getUrl(id string) (url string) {
	url = `https://money.cnn.com/data/markets/`+id+`/`
	return
}

func (obj *Cnn) Process(id string) error {
	url := obj.getUrl(id)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}	
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	ioReader := bytes.NewReader(body)
	doc, err := goquery.NewDocumentFromReader(ioReader)
	if err != nil {
		return err
	}
	
	tag := doc.Find("body")
	for _,item := range pattern {
		tag = tag.Find(item)
		//fmt.Println(item, tag.Text())
	}
	peStr := tag.Eq(2).Find("td").Eq(1).Text()
	fmt.Sscanf(peStr, "%f", &obj.Pe)
	return nil
}

func New() (obj *Cnn) {
	obj = new(Cnn)
	return
}