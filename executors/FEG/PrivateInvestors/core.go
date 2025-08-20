package privateinvestors

// PrivateFund represents a single fund in private investment series. This act as the parent for underlying structures.
type PrivateFund struct {
	fundID        string
	referenceID   string
	workPaperData WorkPaperData
	holder        Holder
}

// WorkPaperData represents the core structure that holds all the data for a particular Fund-Series combination
type WorkPaperData struct {
	totalAssets                   float64
	unfundedCommitments           float64
	totalPartnerCaptial           float64
	totalLiabilities              float64
	investments                   float64
	investmentsUnrealisedGainLoss float64
	cashSweep                     float64
	cashGeneralAccount            float64
	liquidityData                 LiquidityData
	fundPerformanceData           FundPerformanceData
}

// LiquidityData represents the structure that holds all the data retrieved from the investorLiquidity File
type LiquidityData struct {
	top5OwnedPercentage        float64
	nonProfitOwnedPercentage   float64
	pensionPlanOwnedPercentage float64
	otherPlanOwnedPercentage   float64
	usPersonOwnedPercentage    float64
}

// FundPerformanceData represents the strucure that holds the fund performance on three quarters
type FundPerformanceData struct {
	q1GrossPercentage float64
	q2GrossPercentage float64
	q3GrossPercentage float64
	q1NetPercentage   float64
	q2NetPercentage   float64
	q3NetPercentage   float64
}

// Holder holds the path of the files needed for a extraction of a single Private Fund
type Holder struct {
	q1FilePath string
	q2FilePath string
	q3FilePath string
}

// liquidityDataHolder represents the map that holds the liquidityData for the reference fund
type liquidityDataHolder map[string]*LiquidityData
