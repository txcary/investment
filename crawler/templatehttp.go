package crawler

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
)

type TemplateHttp struct {
	TemplateBase
}

func (obj *TemplateHttp) Crawl(id string) (err error) {
	obj.mutex.Lock()
	defer obj.mutex.Unlock()

	if !obj.strategy.CrawlNeeded(id) {
		return
	}
	url := obj.strategy.GetUrl()
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
