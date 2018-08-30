package crawler

import (
)

type Strategy interface {
	Process(id string) error
}

type Crawler struct {
	strategy Strategy	
}

func (obj *Crawler) Process(id string) (err error) {
	err = obj.strategy.Process(id)
	return	
}

func New(strategyInf Strategy) *Crawler {
	obj := new(Crawler)
	obj.strategy = strategyInf
	return obj
}