package crawler

import (
	"net/http"
	"io/ioutil"
	"github.com/bitly/go-simplejson"
)

type TemplateJson struct {
	TemplateBase
}

func (obj *TemplateJson) Crawl(id string) (err error) {
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
	json, err := simplejson.NewJson([]byte(body))
	if err != nil {
		return err
	}

	err = obj.strategy.Process(json)
	return	
}
