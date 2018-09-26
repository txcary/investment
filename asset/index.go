package asset
import (
	
)

type Index struct {
	id string	
}

func (obj *Index)ValueIndicator() (res float64, err error) {

}
func (obj *Index)TrendIndicator() (res float64, err error) {

}

func (obj *Index)Init(id string) {
	obj.id = id
}

func NewIndex(id string) *Index {
	obj := new(Index)
	obj.Init(id)
	return obj
}