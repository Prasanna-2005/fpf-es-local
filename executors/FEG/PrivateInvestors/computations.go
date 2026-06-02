package privateinvestors

import (
	"math"
	"quantrium-tech/form_pf_extractor_service/executors"
)

// POF Specific Formulae

// calculateLevel3Asset calculates the Level3 Asset for PV based on backtracked formula
func calculateLevel3Asset(investment1900 float64, unrealisedGainLoss1910 float64) float64 {
	return investment1900 + unrealisedGainLoss1910
}

// below one is Depriciated formula to calculate CB assets . This Version 1 formula
// calculateCBAsset calculates the CB Asset for PV based on backtracked formula
// func calculateCBAsset(cashChecking1000 float64, cashSweep1002 float64, cashGeneral1005 float64) float64 {
// 	return cashChecking1000 + cashSweep1002 + cashGeneral1005
// }

// This is V2 formula shared for calculating CB assets
func calculateCBAsset(GAV, UnfundedCommitments, FairValueAssets float64) float64 {
	return GAV - UnfundedCommitments - FairValueAssets
}

// Compute act over a private fund to calculate all the required compuations and returns a 1bData since PV only uses it in computations
func (F *PrivateFund) Compute() *executors.Form1bData {

	wd := F.workPaperData
	ld := F.workPaperData.liquidityData
	pd := F.workPaperData.fundPerformanceData

	data := executors.NewForm1bData()

	// workpaper data
	data.GrossAsset = executors.CalculateGrossAssets(wd.totalAssets, wd.unfundedCommitments)
	data.NetAsset = executors.CalculateNetAsset(wd.totalPartnerCaptial, wd.unfundedCommitments)
	data.BorrowedTotalAmount = wd.commericaLineOfCreditPayable
	if data.BorrowedTotalAmount != 0 {
		data.BorrowedUSFinInstiAmt = 100
	}
	data.EquityInOtherFunds = calculateLevel3Asset(wd.investments, wd.investmentsUnrealisedGainLoss)
	data.Level3Asset = data.EquityInOtherFunds
	data.CostBasedAsset = calculateCBAsset(data.GrossAsset, wd.unfundedCommitments, data.Level3Asset)
	data.CostBasedLiability = executors.CalcCBLiabilities(data.GrossAsset, data.NetAsset, 0)

	// investor liquidity data
	data.UsPercent = ld.UsPersonOwnedPercentage
	data.OtherPercent = ld.OtherPlanOwnedPercentage
	data.PensionPlanPercent = ld.PensionPlanOwnedPercentage
	data.NonProfitPercent = ld.NonProfitOwnedPercentage
	data.Top5Percent = ld.Top5OwnedPercentage
	F.calculateFundPerformance()
	// fund performance data
	assignFundPerf(data, pd)

	return data

}

func calculateQuarterPerformance(begBal, capCall, cashDistrib, netIncome, mgmtFee float64) (net, gross float64) {
	if begBal+capCall+cashDistrib != 0 {
		dr := begBal + capCall + cashDistrib
		net = math.Round(((netIncome-mgmtFee)/dr)*10000) / 100
		gross = math.Round((netIncome/dr)*10000) / 100
	}
	// net and gross default to 0 if condition is false
	return
}

func (F *PrivateFund) calculateFundPerformance() {
	pd := F.workPaperData.fundPerformanceData

	pd.q1NetPercentage, pd.q1GrossPercentage = calculateQuarterPerformance(
		pd.begBalq1, pd.capCallq1, pd.cashDistribq1, pd.netIncomeq1, pd.mgmtFeeq1,
	)
	pd.q2NetPercentage, pd.q2GrossPercentage = calculateQuarterPerformance(
		pd.begBalq2, pd.capCallq2, pd.cashDistribq2, pd.netIncomeq2, pd.mgmtFeeq2,
	)
	pd.q3NetPercentage, pd.q3GrossPercentage = calculateQuarterPerformance(
		pd.begBalq3, pd.capCallq3, pd.cashDistribq3, pd.netIncomeq3, pd.mgmtFeeq3,
	)
}
