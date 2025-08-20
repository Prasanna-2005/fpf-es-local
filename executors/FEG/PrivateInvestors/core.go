package privateinvestors

type WorkPaperData struct {
	totalAssets                   int
	unfundedCommitments           int
	totalPartnerCaptial           int
	totalLiabilities              int
	investments                   int
	investmentsUnrealisedGainLoss int
	cashSweep                     int
	cashGeneralAccount            int
}

type LiquidityData struct {
	top5OwnedPercentage        int
	nonProfitOwnedPercentage   int
	pensionPlanOwnedPercentage int
	otherPlanOwnedPercentage   int
	usPersonOwnedPercentage    int
}

type QuarterPerformanceData struct {
	q1GrossPercentage int
	q2GrossPercentage int
	q3GrossPercentage int
	q1NetPercentage   int
	q2NetPercentage   int
	q3NetPercentage   int
}
