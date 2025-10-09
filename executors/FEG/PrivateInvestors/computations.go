package privateinvestors

import "quantrium-tech/form_pf_extractor_service/executors"

// POF Specific Formulae

// calculateLevel3Asset calculates the Level3 Asset for PV based on backtracked formula
func calculateLevel3Asset(investment1900 float64, unrealisedGainLoss1910 float64) float64 {
	return investment1900 + unrealisedGainLoss1910
}

// calculateCBAsset calculates the CB Asset for PV based on backtracked formula
func calculateCBAsset(cash1000 float64, cashSweep1002 float64) float64 {
	return cash1000 + cashSweep1002
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
	data.EquityInOtherFunds = calculateLevel3Asset(wd.investments, wd.investmentsUnrealisedGainLoss)
	data.Level3Asset = data.EquityInOtherFunds
	data.CostBasedAsset = calculateCBAsset(wd.cashGeneralAccount, wd.cashSweep)
	data.CostBasedLiability = wd.totalLiabilities

	// investor liquidity data
	data.UsPercent = ld.usPersonOwnedPercentage
	data.OtherPercent = ld.otherPlanOwnedPercentage
	data.PensionPlanPercent = ld.pensionPlanOwnedPercentage
	data.NonProfitPercent = ld.nonProfitOwnedPercentage
	data.Top5Percent = ld.top5OwnedPercentage

	// fund performance data
	data.Q1GrossValue = pd.q1GrossPercentage
	data.Q2GrossValue = pd.q2GrossPercentage
	data.Q3GrossValue = pd.q3GrossPercentage
	data.Q1NetValue = pd.q1NetPercentage
	data.Q2NetValue = pd.q2NetPercentage
	data.Q3NetValue = pd.q3NetPercentage

	return data

}
