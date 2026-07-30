package curio

import (
	"quantrium-tech/form_pf_extractor_service/executors"
)

// Compute act over a private fund to calculate all the required compuations and returns a 1bData since PV only uses it in computations
func (F *Fund) Compute() *executors.Form1bData {

	wd := F.workPaperData
	ld := F.workPaperData.liquidityData
	pd := F.workPaperData.irr

	data := executors.NewForm1bData()

	// workpaper data
	data.GrossAsset = executors.CalculateGrossAssets(wd.totalAssets, wd.unfundedCommitments)
	data.NetAsset = executors.CalculateNetAsset(wd.totalPartnerCaptial, wd.unfundedCommitments)
	data.EquityInOtherFunds = wd.sumLevelExtData
	data.BorrowedTotalAmount = wd.commericaLineOfCreditPayable
	if data.BorrowedTotalAmount != 0 {
		data.BorrowedUSFinInstiAmt = 100
	}
	data.Level1Asset = wd.l1
	data.Level2Asset = wd.l2
	data.Level3Asset = wd.l3
	data.CostBasedAsset = calculateCBAsset(data.GrossAsset, wd.sumLevelExtData, wd.unfundedCommitments)
	data.CostBasedLiability = executors.CalcCBLiabilities(data.GrossAsset, data.NetAsset, 0)

	// investor liquidity data
	data.Top5Percent = ld.Top5OwnedPercentage                           // 15
	data.UsPercent = ld.UsPersonOwnedPercentage                         // 16_a
	data.NonUSPersonsPercent = ld.NotUsPersonOwnedPercentage            // 16_b
	data.BrokerDealersPercent = ld.BrokerOwnedPercentage                // 16_c
	data.InsuranceCompaniesPercent = ld.InsuranceOwnedPercentage        // 16_d
	data.InvestmentCompaniesPercent = ld.RiaOwnedPercentage             // 16_e
	data.PrivateFundPercent = ld.PvtFundPercentage                      // 16_f
	data.NonProfitPercent = ld.NonProfitOwnedPercentage                 // 16_g
	data.PensionPlanPercent = ld.PensionPlanOwnedPercentage             // 16_h
	data.BankingInstitutionsPercent = ld.PropOwnedPercentage            // 16_i
	data.GovernmentEntitiesPercent = ld.GovtOwnedPercentage             // 16_j
	data.GovernmentPensionsPercent = ld.GovtPensionPlansOwnedPercentage // 16_k
	data.SovereignPercent = ld.SovereignFundPercentage                  // 16_l
	data.NonUSIntermediariesPercent = ld.UnknownInstiOwnedPercentage    // 16_m
	data.OtherPercent = ld.OtherPlanOwnedPercentage                     // 16_n

	assignFundPerf(data, pd)

	return data

}

func calculateCBAsset(GrossAsset, LevelingSum, unfdCmts float64) float64 {
	return GrossAsset - LevelingSum - unfdCmts
}
