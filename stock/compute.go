package stock

import (
	"github.com/txcary/investment/common/utils"
)

const (
	StockDividentTaxRate float64 = 0.2
)

func minAverage(list []float64) float64 {
	length := int(len(list) / 2)
	avgAll := utils.Average(list...)
	avgRecent := utils.Average(list[:length]...)
	return utils.Min(avgAll, avgRecent)
}

func (obj *Stock) computeAvgGlowth() {
	profitGlowth := minAverage(obj.ProfitGlowthList)
	assetGlowth := minAverage(obj.TotalAssetsGlowthList)
	incomeGlowth := minAverage(obj.OperatingIncomeGlowthList)
	costGlowth := minAverage(obj.CostGlowthList)
	obj.AvgGlowth = utils.Min(costGlowth, profitGlowth, assetGlowth, incomeGlowth)
}

func (obj *Stock) computeAvgRoe() {
	obj.AvgRoe = minAverage(obj.RoeList)
}

func (obj *Stock) computeAjustedPe() {
	obj.AdjustedPe = obj.Pb / obj.AvgRoe
}

func (obj *Stock) computeAvgProfitDividendRate() {
	obj.AvgProfitDividendRate = utils.Average(obj.ProfitDividendRateList...)
}

func (obj *Stock) computeSaftyFactorFromAdjustedPe() {
	obj.SaftyFactor = utils.Min(1, 20/obj.AdjustedPe)
}

func (obj *Stock) computeExpectionFactor() {
	obj.ExpectionFactor = utils.Min(1, obj.Roe/obj.AvgRoe)
}

func (obj *Stock) computeCashFactor() {
	cashFactorList := make([]float64, 0)
	for i, _ := range obj.ProfitList {
		if obj.ProfitList[i] == 0 {
			utils.Error("Profit==0! Ignored")
			continue
		}
		cashFactorList = append(cashFactorList, obj.OperatingCaseFlowList[i]/obj.ProfitList[i])
	}
	obj.CashFactor = utils.Min(1, minAverage(cashFactorList))
}

func (obj *Stock) computeFlowFactor() {
	flowFactorList := make([]float64, 0)
	for i, _ := range obj.CurrentLiabilitiesList {
		var factor float64 = 1
		if obj.CurrentLiabilitiesList[i] != 0 {
			factor = obj.CurrentAssetList[i] / obj.CurrentLiabilitiesList[i]
			flowFactorList = append(flowFactorList, factor)
		}
	}
	obj.FlowFactor = utils.Min(1, minAverage(flowFactorList))
}

func (obj *Stock) computeReturnLow() {
}

func (obj *Stock) computeReturn() {
	minReturn := obj.AvgRoe / obj.Pb
	minDividendAfterTex := minReturn * obj.AvgProfitDividendRate * (1 - StockDividentTaxRate)
	endogenousIncrement := utils.Min(obj.AvgRoe*(1-obj.AvgProfitDividendRate), obj.AvgGlowth)
	repoIncrement := utils.Max(0, obj.AvgRoe*(1-obj.AvgProfitDividendRate)-obj.AvgGlowth) / obj.Pb
	financingRate := utils.Max(0, obj.AvgGlowth-obj.AvgRoe*(1-obj.AvgProfitDividendRate))
	financingIncrement := (obj.Pb - 1) * financingRate / (obj.Pb + financingRate)

	obj.ReturnStd = minDividendAfterTex + endogenousIncrement + repoIncrement + financingIncrement
	obj.ReturnLow = minReturn
	obj.ReturnHigh = utils.Max(minReturn+obj.AvgGlowth, obj.Roe)

}

func (obj *Stock) computeAntiEconomicCycle() {
	obj.AntiEconomicCycle = utils.Min(obj.RoeList...) / obj.AvgRoe
}

func (obj *Stock) computeMarket() {
	obj.computeAjustedPe()
	obj.computeSaftyFactorFromAdjustedPe()
	obj.computeReturn()
}

func (obj *Stock) computeFinance() {
	obj.computeAvgRoe()
	obj.computeAvgGlowth()
	obj.computeAvgProfitDividendRate()
	obj.computeExpectionFactor()
	obj.computeCashFactor()
	obj.computeFlowFactor()
	obj.computeAntiEconomicCycle()
}
