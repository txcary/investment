package asset

import (
	"github.com/txcary/goutils"
	"github.com/txcary/investment/crawler"
)

type HkIndex struct {
	AssetBase
}

var fundMapToIndex = map[string]string{
	"159920": "HKHSI",
}

func (obj *HkIndex) Init(id string) {
	obj.AssetBase.Init(id)
	go func() {
		sohufund := crawler.NewSohufund()
		sohufund.Crawl(obj.id)
		obj.setTrend(sohufund.Trend)
	}()
	go func() {
		id, ok := fundMapToIndex[obj.id]
		if !ok {
			utils.Error(obj.id + " is not a valid Hong Kong index fund!")
			obj.setYield(0)
		}
		danjuan := crawler.NewDanjuan()
		danjuan.Crawl(id)
		obj.setYield(100 / danjuan.Pe)
	}()
}

func NewHkIndex(id string) *HkIndex {
	obj := new(HkIndex)
	obj.Init(id)
	return obj
}
