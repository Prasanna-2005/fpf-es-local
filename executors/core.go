package executors

// Form1bData holds the structure of the section 1b in Form PF.
// It will be dynamically increased based on the needs to answer other questions
type Form1bData struct {
	GrossAsset         float64 `fieldNo:"11" formLineNo:"a" date:"09/30"`
	NetAsset           float64 `fieldNo:"12" formLineNo:"a" date:"09/30"`
	EquityInOtherFunds float64 `fieldNo:"15" formLineNo:"a"`
	Level3Asset        float64 `fieldNo:"20" formLineNo:"level3-asset" date:"09/30"`
	CostBasedAsset     float64 `fieldNo:"20" formLineNo:"cb-asset" date:"09/30"`
	CostBasedLiability float64 `fieldNo:"20" formLineNo:"cb-liability" date:"09/30"`
	Top5Percent        float64 `fieldNo:"21" formLineNo:"21"`
	UsPercent          float64 `fieldNo:"22" formLineNo:"a"`
	NonProfitPercent   float64 `fieldNo:"22" formLineNo:"j"`
	PensionPlanPercent float64 `fieldNo:"22" formLineNo:"l"`
	OtherPercent       float64 `fieldNo:"22" formLineNo:"t"`
	Q1GrossValue       float64 `fieldNo:"23" formLineNo:"xiii-gross" date:"03/31"`
	Q1NetValue         float64 `fieldNo:"23" formLineNo:"xiii-net" date:"03/31"`
	Q2GrossValue       float64 `fieldNo:"23" formLineNo:"xiv-gross" date:"06/30"`
	Q2NetValue         float64 `fieldNo:"23" formLineNo:"xiv-net" date:"06/30"`
	Q3GrossValue       float64 `fieldNo:"23" formLineNo:"xv-gross" date:"09/30"`
	Q3NetValue         float64 `fieldNo:"23" formLineNo:"xv-net" date:"09/30"`
}
