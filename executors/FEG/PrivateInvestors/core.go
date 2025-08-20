package privateinvestors

// PrivateFund represents a single fund in private investment series. This act as the parent for underlying structures.
type PrivateFund struct {
	fundID        string
	workPaperData WorkPaperData
	holder        Holder
}

// WorkPaperData represents the core structure that holds all the data for a particular Fund-Series combination
type WorkPaperData struct {
	totalAssets                   int
	unfundedCommitments           int
	totalPartnerCaptial           int
	totalLiabilities              int
	investments                   int
	investmentsUnrealisedGainLoss int
	cashSweep                     int
	cashGeneralAccount            int
	liquidityData                 LiquidityData
	fundPerformanceData           FundPerformanceData
}

// LiquidityData represents the structure that holds all the data retrieved from the investorLiquidity File
type LiquidityData struct {
	top5OwnedPercentage        int
	nonProfitOwnedPercentage   int
	pensionPlanOwnedPercentage int
	otherPlanOwnedPercentage   int
	usPersonOwnedPercentage    int
}

// FundPerformanceData represents the strucure that holds the fund performance on three quarters
type FundPerformanceData struct {
	q1GrossPercentage int
	q2GrossPercentage int
	q3GrossPercentage int
	q1NetPercentage   int
	q2NetPercentage   int
	q3NetPercentage   int
}

// Holder holds the path of the files needed for a extraction of a single Private Fund
type Holder struct {
	q1FilePath                string
	q2FilePath                string
	q3FilePath                string
	investorLiquidityFilePath string
}
