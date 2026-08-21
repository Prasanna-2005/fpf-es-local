package invesmentservices

import (
	"fmt"
	"quantrium-tech/form_pf_extractor_service/executors"
	"quantrium-tech/form_pf_extractor_service/logging"
	"slices"
)

// Compute act over a private fund to calculate all the required compuations and returns a 1bData since PV only uses it in computations
func (F *Fund) Compute() *executors.Form1bData {

	wd := F.workPaperData
	ld := F.workPaperData.liquidityData
	pd := F.workPaperData.fundPerformanceData

	data := executors.NewForm1bData()
	tracker := executors.NewFormStateTracker(*data)
	// log the field-level diff no matter how the function returns
	defer func() {
		changes := tracker.DiffForm(data)
		if len(changes) == 0 {
			logging.Info("No data changed")
			return
		}
		F.workPaperAudit = make(executors.FieldChangeAuditor, len(changes))
		logging.Info(fmt.Sprintf("WorkPaperData changed for fund %s", F.fundID))
		for _, c := range changes {
			F.FormDataAudit[c.Entity] = c.Change
			fmt.Printf("%v: %q -> %q\n", c.Entity, c.Old, c.New)
		}

	}()

	// workpaper data
	data.GrossAsset = executors.CalculateGrossAssets(wd.totalAssets, wd.unfundedCommitments)
	data.NetAsset = executors.CalculateNetAsset(wd.totalPartnerCaptial, wd.unfundedCommitments)
	data.EquityInOtherFunds = wd.sumLevelExtData
	if slices.Contains(NotSame, F.fundID) {
		data.EquityInOtherFunds = 0
	}

	data.BorrowedTotalAmount = wd.commericaLineOfCreditPayable
	if data.BorrowedTotalAmount != 0 {
		data.BorrowedUSFinInstiAmt = 100
	}

	data.Level1Asset = wd.l1
	data.Level2Asset = wd.l2
	data.Level3Asset = wd.l3
	data.CostBasedAsset = calculateCBAsset(data.GrossAsset, wd.sumLevelExtData, wd.unfundedCommitments)
	data.Level2Liability = wd.liabilityL2
	data.CostBasedLiability = executors.CalcCBLiabilities(data.GrossAsset, data.NetAsset, data.Level2Liability)

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
