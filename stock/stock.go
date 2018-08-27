package stock

import (
//  "goplatform/model"
)

type Finance struct {
	Roe                       float64
	Glowth                    float64
	RoeList                   []float64
	ProfitList                []float64
	TotalAssetsGlowthList     []float64
	ProfitGlowthList          []float64
	CostGlowthList            []float64
	OperatingIncomeGlowthList []float64
	OperatingCaseFlowList     []float64
	ProfitDividendRateList    []float64
	CurrentAssetList          []float64
	CurrentLiabilitiesList    []float64
}

type Market struct {
	Price        float64
	Pb           float64
	Pe           float64
	DividendRate float64
}

type Info struct {
	Id   string
	Name string
}

type Computed struct {
	AvgProfitDividendRate float64
	AdjustedPe            float64
	SaftyFactor           float64
	ReturnLow             float64
	ReturnStd             float64
	ReturnHigh            float64
	AvgRoe                float64
	AvgGlowth             float64
	ExpectionFactor       float64
	CashFactor            float64
	FlowFactor            float64
	AntiEconomicCycle     float64
}

type Stock struct {
	Info
	Market
	Finance
	Computed
}

func (obj *Stock) SetInfo(info Info) {
	obj.Info = info
}

func (obj *Stock) SetMarket(marketInfo Market) {
	obj.Market = marketInfo
	obj.computeMarket()
}

func (obj *Stock) SetFinance(financeInfo Finance) {
	obj.Finance = financeInfo
	obj.computeFinance()
}

func (obj *Stock) Init(stockInfo Info, marketInfo Market, financeInfo Finance) {
	obj.SetFinance(financeInfo)
	obj.SetMarket(marketInfo)
	obj.SetInfo(stockInfo)
}

func New(stockInfo Info, marketInfo Market, financeInfo Finance) *Stock {
	obj := new(Stock)
	obj.Init(stockInfo, marketInfo, financeInfo)
	return obj
}
