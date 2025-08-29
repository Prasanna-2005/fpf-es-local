package executors

// Form1bData holds the structure of the section 1b in Form PF.
// It will be dynamically increased based on the needs to answer other questions
type Form1bData struct {
	GrossAsset         float64
	NetAsset           float64
	EquityInOtherFunds float64
	Level3Asset        float64
	CostBasedAsset     float64
	CostBasedLiability float64
	Top5Percent        float64
	UsPercent          float64
	NonProfitPercent   float64
	PensionPlanPercent float64
	OtherPercent       float64
	Q1GrossValue       float64
	Q1NetValue         float64
	Q2GrossValue       float64
	Q2NetValue         float64
	Q3GrossValue       float64
	Q3NetValue         float64
}
