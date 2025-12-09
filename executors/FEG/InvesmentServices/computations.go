package invesmentservices

import (
	"fmt"
	"quantrium-tech/form_pf_extractor_service/executors"
	"slices"
)

// Compute act over a private fund to calculate all the required compuations and returns a 1bData since PV only uses it in computations
func (F *Fund) Compute() *executors.Form1bData {

	wd := F.workPaperData
	ld := F.workPaperData.liquidityData
	pd := F.workPaperData.fundPerformanceData

	data := executors.NewForm1bData()

	// workpaper data
	fmt.Printf("\nfund:%v  computation:%v\n", F.fundID, F.workPaperData.unfundedCommitments)
	data.GrossAsset = executors.CalculateGrossAssets(wd.totalAssets, wd.unfundedCommitments)
	data.NetAsset = executors.CalculateNetAsset(wd.totalPartnerCaptial, wd.unfundedCommitments)
	data.EquityInOtherFunds = wd.sumLevelExtData
	if slices.Contains(NotSame, F.fundID) {
		fmt.Printf("\nInside\n %v", F.fundID)
		data.EquityInOtherFunds = 0
	}
	data.Level1Asset = wd.l1
	data.Level2Asset = wd.l2
	data.Level3Asset = wd.l3
	data.CostBasedAsset = calculateCBAsset(data.GrossAsset, wd.sumLevelExtData)
	data.CostBasedLiability = wd.totalLiabilities

	// investor liquidity data
	data.UsPercent = ld.usPersonOwnedPercentage
	data.OtherPercent = ld.otherPlanOwnedPercentage
	data.PensionPlanPercent = ld.pensionPlanOwnedPercentage
	data.NonProfitPercent = ld.nonProfitOwnedPercentage
	data.Top5Percent = ld.top5OwnedPercentage
	data.SovereignPercent = ld.sovereignFundPercentage
	data.PrivateFundPercent = ld.pvtFundPercentage

	/* ---------- monthly gross ---------- */
	data.JanGrossValue = pd.janGrossPercentage
	data.FebGrossValue = pd.febGrossPercentage
	data.MarGrossValue = pd.marGrossPercentage
	data.AprGrossValue = pd.aprGrossPercentage
	data.MayGrossValue = pd.mayGrossPercentage
	data.JunGrossValue = pd.junGrossPercentage
	data.JulGrossValue = pd.julGrossPercentage
	data.AugGrossValue = pd.augGrossPercentage
	data.SepGrossValue = pd.sepGrossPercentage
	data.OctGrossValue = pd.octGrossPercentage
	data.NovGrossValue = pd.novGrossPercentage
	data.DecGrossValue = pd.decGrossPercentage

	/* ---------- monthly net ---------- */
	data.JanNetValue = pd.janNetPercentage
	data.FebNetValue = pd.febNetPercentage
	data.MarNetValue = pd.marNetPercentage
	data.AprNetValue = pd.aprNetPercentage
	data.MayNetValue = pd.mayNetPercentage
	data.JunNetValue = pd.junNetPercentage
	data.JulNetValue = pd.julNetPercentage
	data.AugNetValue = pd.augNetPercentage
	data.SepNetValue = pd.sepNetPercentage
	data.OctNetValue = pd.octNetPercentage
	data.NovNetValue = pd.novNetPercentage
	data.DecNetValue = pd.decNetPercentage

	/* ---------- quarterly ---------- */
	data.Q1GrossValue = pd.q1GrossPercentage
	data.Q2GrossValue = pd.q2GrossPercentage
	data.Q3GrossValue = pd.q3GrossPercentage
	data.Q4GrossValue = pd.q4GrossPercentage

	data.Q1NetValue = pd.q1NetPercentage
	data.Q2NetValue = pd.q2NetPercentage
	data.Q3NetValue = pd.q3NetPercentage
	data.Q4NetValue = pd.q4NetPercentage

	/* ---------- annual ---------- */
	data.YearGrossValue = pd.yearGrossPercentage
	data.YearNetValue = pd.yearNetPercentage
	fmt.Printf("\npd inside comp:%v\n", pd)

	return data

}

func calculateCBAsset(GrossAsset, LevelingSum float64) float64 {
	return GrossAsset - LevelingSum
}
