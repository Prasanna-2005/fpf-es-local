package invesmentservices

// ShortCodeReverse maps the long form back to the short code.
var ShortCodeConvertor = map[string]string{
	"FEG_SAGREITSI":        "FSREIT",
	"FEG_SELAEMERG":        "FSAEM",
	"FEG_SELBIOTECH":       "FSBIO",
	"FEG_SELECTACTIVE":     "FSAIE",
	"FEG_SELECTCORE":       "FSCFI",
	"FEG_SELECTDOMESTICII": "FSADE2",
	"FEG_SELECTGLOBAL":     "FSAGE",
	"FEG_SELECTGLOBALII":   "FSAGE2",
	"FEG_SELECTLLC":        "N/A",
	"FEG_SELECTPLUS":       "FSCPB",
	"FEG_SELECTSTREME":     "FSSEM",
	"FEG_SELECTSTRINT":     "FSSIE",
	"FEG_SELECTSTRUS":      "FSSUS",
	"FEG_SAEMES_II":        "FSAEM2",
	"FEG_SELDSS":           "FSDS",
	"FEG_SSCGES":           "FSCGES",
	"FEG_MPES":             "FEGMP",
	"FEG_SSCGSP":           "FSCGSP",
	"FEG_SELECTACTIVETCI":  "FSAE",
	"FEGSAESII":            "FSAE2",
	"FEG_SELECTPRVCAP":     "FSOPO",
	"FEG_SELECTPRVEQ":      "FSTGXI", // first occurrence wins
	"FEGSPESTGXII":         "FSTGX2", // both FSTGXII and FSTGX2 map to same value
	"FEGSPESTGXIV":         "FSTGX4", // both FSTGXIV and FSTGX4 map to same value
}

var LevelMapper = map[string]map[string][]string{
	"FEGMP": {
		"l3": {"invt. in master fund"},
	},
	"FSAEM": {
		"l3": {"invt hedge funds @ cost", "invt in hedge funds u/r"},
	},
	"FSAEM2": {
		"l3": {"invt hedge funds @ cost", "invt in hedge funds u/r"},
	},
	"FSAE": {
		"l3": {"invt hedge funds @ cost", "invt in hedge funds u/r"},
	},
	"FSAE2": {
		"l3": {"invt hedge funds @ cost", "invt in hedge funds u/r"},
	},
	"FSAGE": {
		"l3": {"invt hedge funds @ cost", "invt in hedge funds u/r"},
	},
	"FSAGE2": {
		"l3": {"invt hedge funds @ cost", "invt in hedge funds u/r", "direct holding @ cost", "direct holding u/r g/l"},
	},
	"FSCGES": {
		"l3": {"invt hedge funds @ cost", "invt in hedge funds u/r", "direct holding @ cost", "direct holding u/r g/l"},
	},
	"FSAIE": {
		"l1": {"equities @ cost", "equities u/r g/l", "equities u/r fx g/l", "adr @ cost", "adr u/r g/l", "gdr @ cost", "gdr u/r g/l"},
	},
	"FSSEM": {
		"l3": {"invt hedge funds @ cost", "invt in hedge funds u/r"},
	},
	"FSDS": {
		"l3": {"invt hedge funds @ cost", "invt in hedge funds u/r"},
	},
	"FSCFI": {
		"l2": {
			"corp bonds @ cost",
			"corp bonds u/r g/l",
			"govt bonds @ cost",
			"govt bonds u/r g/l",
			"treasuries @ cost",
			"treasuries u/r g/l",
			"mortgage & cmo @ cost",
			"mortgage & cmo u/r g/l",
		},
	},
	"FSBIO": {
		"l3": {"invt hedge funds @ cost", "invt in hedge funds u/r"},
	},
	"FSOPO": {
		"l3": {"private eq @ cost", "private eq u/r g/l"},
	},
	"FSTGX1": {
		"l3": {"private eq @ cost", "private eq u/r g/l"},
	},
	"FSTGX2": {
		"l3": {"private eq @ cost", "private eq u/r g/l"},
	},
	"FSTGX4": {
		"l3": {"equities @ cost", "equities u/r g/l", "reit @ cost", "reit u/r g/l"},
	},
	"FSSUS": {
		"l2": {
			"equities @ cost",
			"equities u/r g/l",
			"reit @ cost",
			"reit u/r g/l",
		},
	},
	"FSSIE": {
		"l2": {
			"equities @ cost",
			"equities u/r g/l",
			"reit @ cost",
			"reit u/r g/l",
		},
	},
}

var LevelingEntityMapper = map[string]string{
	"adr @ cost":              "adrCost",
	"adr u/r g/l":             "adrURGL",
	"corp bonds @ cost":       "corpBondsCost",
	"corp bonds u/r g/l":      "corpBondsURGL",
	"direct holding @ cost":   "directHoldingCost",
	"direct holding u/r g/l":  "directHoldingURGL",
	"equities @ cost":         "equitiesCost",
	"equities dividend rec":   "equitiesDividendRec",
	"equities ft w/h payable": "equitiesFtWhPayable",
	"equities u/r fx g/l":     "equitiesURFXGL",
	"equities u/r g/l":        "equitiesURGL",
	"gdr @ cost":              "gdrCost",
	"gdr u/r g/l":             "gdrURGL",
	"govt bonds @ cost":       "govtBondsCost",
	"govt bonds u/r g/l":      "govtBondsURGL",
	"invt hedge funds @ cost": "invtHedgeFundsCost",
	"invt in hedge funds u/r": "invtInHedgeFundsUR",
	"mortgage & cmo @ cost":   "mortgageCmoCost",
	"mortgage & cmo u/r g/l":  "mortgageCmoURGL",
	"private eq @ cost":       "privateEqCost",
	"private eq u/r g/l":      "privateEqURGL",
	"reit @ cost":             "reitCost",
	"reit u/r g/l":            "reitURGL",
	"invt. in master fund":    "invtInMasterFund",
	"treasuries @ cost":       "treasuriesCost",
	"treasuries u/r g/l":      "treasuriesURGL",
}
