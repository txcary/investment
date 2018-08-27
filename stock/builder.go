package stock 

import (
	"github.com/txcary/lixinger"
	"github.com/txcary/investment/config"
)

type Builder struct {
	stockMap map[string]*Stock
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

func (obj *Builder) newStock(id string) *Stock {
	info, _ := obj.getInfo(id)
	market, _ := obj.getMarket(id)
	finance, _ := obj.getFinance(id)
	stockObj := New(info, market, finance)
	return stockObj
}

func (obj *Builder) Build(id string) *Stock {
	if _,ok := obj.stockMap[id]; !ok {
		obj.stockMap[id] = obj.newStock(id)
	}
	return obj.stockMap[id]	
}

func BuilderInstance() *Builder {
	if obj == nil {
		token := config.Instance().GetString("lixinger", "token")

		obj = new(Builder)
		obj.lixingerObj = lixinger.New(token)
		obj.stockMap = make(map[string]*Stock)
	}

	return obj
}
