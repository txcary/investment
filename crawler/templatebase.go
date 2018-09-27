package crawler
import (
	"sync"
)
type TemplateBase struct {
	mutex sync.Mutex
	strategy Strategy	
}
func (obj *TemplateBase) SetStrategyToTemplate(strategyInf Strategy) {
	obj.strategy = strategyInf
}
