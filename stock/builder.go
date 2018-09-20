package stock 

import (
	"sync"
	"fmt"
	"github.com/txcary/investment/db"
	"github.com/txcary/lixinger"
	"github.com/txcary/investment/config"
)

type Builder struct {
	stockMap sync.Map
	stockDb *db.Db
	lixingerObj *lixinger.Lixinger
}

var obj *Builder

func (obj *Builder) getInfo(id string) (res Info, err error) {
	res.Name, err = obj.lixingerObj.GetString(id, "latest", "stockCnName")
	res.Id = id
	return
}

func (obj *Builder) getMarket(id string) (res Market, err error) {
	res.Price, err = obj.lixingerObj.GetFloat64(id, "latest", "stock_price")
	res.Pb, err = obj.lixingerObj.GetFloat64(id, "latest", "pb")
	res.Pe, err = obj.lixingerObj.GetFloat64(id, "latest", "pe_ttm")
	res.DividendRate, err = obj.lixingerObj.GetFloat64(id, "latest", "dividend_r")
	return
}

func (obj *Builder) getFinance(id string) (res Finance, err error) {
	res.Roe, err = obj.lixingerObj.GetFloat64(id, "latest", "q.metrics.roe.ttm")
	res.Glowth, err = obj.lixingerObj.GetFloat64(id, "latest", "q.profitStatement.np.ttm_y2y")

	res.RoeList, err = obj.lixingerObj.FilterFloat64(id, "12-31", "q.metrics.roe.ttm")
	res.ProfitList, err = obj.lixingerObj.FilterFloat64(id, "12-31", "q.profitStatement.np.ttm")
	res.TotalAssetsGlowthList, err = obj.lixingerObj.FilterFloat64(id, "12-31", "q.balanceSheet.ta.t_y2y")
	res.ProfitGlowthList, err = obj.lixingerObj.FilterFloat64(id, "12-31", "q.profitStatement.np.ttm_y2y")
	res.CostGlowthList, err = obj.lixingerObj.FilterFloat64(id, "12-31", "q.profitStatement.oc.ttm_y2y")
	res.OperatingIncomeGlowthList, err = obj.lixingerObj.FilterFloat64(id, "12-31", "q.profitStatement.toi.ttm_y2y")
	res.OperatingCaseFlowList, err = obj.lixingerObj.FilterFloat64(id, "12-31", "q.cashFlow.ncffoa.ttm")
	res.CurrentAssetList, err = obj.lixingerObj.FilterFloat64(id, "12-31", "q.balanceSheet.tca.t")
	res.CurrentLiabilitiesList, err = obj.lixingerObj.FilterFloat64(id, "12-31", "q.balanceSheet.tcl.t")
	
	// TODO: Got recent 3 years
	res.ProfitDividendRateList = make([]float64,1)
	pe, err := obj.lixingerObj.GetFloat64(id, "latest", "pe_ttm")
	div, err := obj.lixingerObj.GetFloat64(id, "latest", "dividend_r")
	res.ProfitDividendRateList[0] = pe*div
		
	return
}

func (obj *Builder) newStockFromCatch(id string) *Stock {
	info, _ := obj.getInfo(id)
	market, _ := obj.getMarket(id)
	finance, _ := obj.getFinance(id)
	fmt.Println("Start Catch")
	stockobj := New(info, market, finance)
	fmt.Println("Catch Done")
	return stockobj
}

func (obj *Builder) newStockFromDb(id string) *Stock {
	stockobj := new(Stock) 
	err := obj.stockDb.Get(id, stockobj)
	if err == nil {
		return stockobj
	}
	return nil
}

func (obj *Builder) newStock(id string) *Stock {
	var stockobj *Stock
	stockobj = obj.newStockFromDb(id)
	if stockobj == nil {
		stockobj = obj.newStockFromCatch(id)
		obj.putDb(id, stockobj)
	}
	return stockobj
}

func (obj *Builder) putDb(id string, stockobj *Stock) {
	obj.stockDb.Put(id, stockobj)
}

func (obj *Builder) UpdateIfExpired(id string) (updated bool) {
	var v interface{}
	var ok bool
	updated = false
	if v, ok = obj.stockMap.Load(id); !ok {
		return
	}
	stockobj := v.(*Stock)
	if stockobj.IsFinanceExpired() {
		finance, _ := obj.getFinance(id)
		stockobj.SetFinance(finance)
		updated = true
	}
	if stockobj.IsMarketExpired() {
		market, _ := obj.getMarket(id)
		stockobj.SetMarket(market)
		updated = true
	}
	if updated {
		obj.putDb(id, stockobj)
	}
	return
}

func (obj *Builder) Build(id string) *Stock {
	var v interface{}
	var ok bool
	if v, ok = obj.stockMap.Load(id); !ok {
		v, ok = obj.stockMap.LoadOrStore(id, obj.newStock(id))
	}
	stockobj := v.(*Stock)
	obj.UpdateIfExpired(id)
	return stockobj
}

func BuilderInstance() *Builder {
	if obj == nil {
		token := config.Instance().GetString("lixinger", "token")

		obj = new(Builder)
		obj.lixingerObj = lixinger.New(token)

		dbpath := config.Instance().GetString("stock", "db")
		obj.stockDb = db.New(dbpath)
	}

	return obj
}
