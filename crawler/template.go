package crawler

import (
	"bytes"
	"net/http"
	"io/ioutil"
	"github.com/PuerkitoBio/goquery"
)

type Strategy interface {
	GetUrl(id string) (string)
	Process(*goquery.Document) error
}

type Template struct {
	strategy Strategy	
}

func (obj *Template) Crawl(id string) (err error) {
	url := obj.strategy.GetUrl(id)
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

	err = obj.strategy.Process(doc)
	return	
}

func (obj *Template) Init(strategyInf Strategy) {
	obj.strategy = strategyInf
}
