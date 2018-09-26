package asset
import (
	
)

type Asset interface {
	ValueIndicator() (float64, error)
	TrendIndicator() (float64, error)
}

type AssetBase struct {
	
}

func (obj *AssetBase)Init() {
	
}

func New() *Asset {
	obj := new(Asset)
	obj.Init()
	return obj
}