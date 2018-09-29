package asset

import ()

type Asset interface {
	YieldIndicator() float64
	TrendIndicator() float64
}

type AsyncData struct {
	Value    float64
	Chan     chan bool
	IsRead   bool
	IsWrited bool
}

type AssetBase struct {
	id    string
	yield *AsyncData
	trend *AsyncData
}

func (obj *AsyncData) Get() float64 {
	if !obj.IsRead {
		obj.IsRead = <-obj.Chan
	}
	return obj.Value
}

func (obj *AsyncData) Set(value float64) {
	obj.Value = value
	if !obj.IsWrited {
		obj.Chan <- true
		obj.IsWrited = true
	}
}

func NewAsyncData() *AsyncData {
	obj := new(AsyncData)
	obj.Chan = make(chan bool, 1)
	return obj
}

func (obj *AssetBase) setYield(value float64) {
	obj.yield.Set(value)
}

func (obj *AssetBase) setTrend(value float64) {
	obj.trend.Set(value)
}

func (obj *AssetBase) YieldIndicator() (res float64) {
	return obj.yield.Get()
}

func (obj *AssetBase) TrendIndicator() (res float64) {
	return obj.trend.Get()
}

func (obj *AssetBase) Init(id string) {
	obj.id = id
	obj.yield = NewAsyncData()
	obj.trend = NewAsyncData()
}

func New(id string) *AssetBase {
	obj := new(AssetBase)
	obj.Init(id)
	return obj
}
