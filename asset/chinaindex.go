package asset
import (
	"github.com/txcary/investment/crawler"
	
)

type ChinaIndex struct {
	AssetBase
	sohufund *crawler.Sohufund
	jisiluetf *crawler.Jisiluetf
}

func (obj *ChinaIndex)Init(id string) {
	obj.AssetBase.Init(id)
	go func() {
		sohufund := crawler.NewSohufund()
		sohufund.Crawl(obj.id)	
		obj.setTrend(sohufund.Trend)
	}()
	go func() {
		jisiluetf := crawler.NewJisiluetf()
		jisiluetf.Crawl(obj.id)
		obj.setYield(100/jisiluetf.Pe)	
	}()
}

func NewChinaIndex(id string) *ChinaIndex {
	obj := new(ChinaIndex)
	obj.Init(id)
	return obj
}