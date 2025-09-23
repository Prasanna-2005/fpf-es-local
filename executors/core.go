package executors

// Form1bData holds the structure of the section 1b in Form PF.
// It will be dynamically increased based on the needs to answer other questions
type Form1bData struct {
	FundName string `fieldNo:"5" formLineNo:"5_a"`
	FundID   string `fieldNo:"5" formLineNo:"5_b"`
	FundNFA  string `fieldNo:"5" formLineNo:"5_c"`
	FundLEI  string `fieldNo:"5" formLineNo:"5_d"`

	GrossAsset         float64 `fieldNo:"11" formLineNo:"11_a" date:"09/30"`
	NetAsset           float64 `fieldNo:"12" formLineNo:"12_a" date:"09/30"`
	EquityInOtherFunds float64 `fieldNo:"15" formLineNo:"15_a"`
	Level3Asset        float64 `fieldNo:"20" formLineNo:"20_level3-asset" date:"09/30"`
	CostBasedAsset     float64 `fieldNo:"20" formLineNo:"20_cb-asset" date:"09/30"`
	CostBasedLiability float64 `fieldNo:"20" formLineNo:"20_cb-liability" date:"09/30"`
	Top5Percent        float64 `fieldNo:"21" formLineNo:"21"`
	UsPercent          float64 `fieldNo:"22" formLineNo:"22_a"`
	NonProfitPercent   float64 `fieldNo:"22" formLineNo:"22_j"`
	PensionPlanPercent float64 `fieldNo:"22" formLineNo:"22_l"`
	OtherPercent       float64 `fieldNo:"22" formLineNo:"22_t"`
	Q1GrossValue       float64 `fieldNo:"23" formLineNo:"23_xiii-gross" date:"03/31"`
	Q1NetValue         float64 `fieldNo:"23" formLineNo:"23_xiii-net" date:"03/31"`
	Q2GrossValue       float64 `fieldNo:"23" formLineNo:"23_xiv-gross" date:"06/30"`
	Q2NetValue         float64 `fieldNo:"23" formLineNo:"23_xiv-net" date:"06/30"`
	Q3GrossValue       float64 `fieldNo:"23" formLineNo:"23_xv-gross" date:"09/30"`
	Q3NetValue         float64 `fieldNo:"23" formLineNo:"23_xv-net" date:"09/30"`
}
