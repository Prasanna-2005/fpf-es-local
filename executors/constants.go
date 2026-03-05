package executors

var ShortCodeConvertor = map[string]string{
	"FEG_SAGREITSI":        "FSREIT",
	"FEG_SELAEMERG":        "FSAEM",
	"FEG_SELBIOTECH":       "FSBIO",
	"FEG_SELECTACTIVE":     "FSAIE",
	"FEG_SELECTCORE":       "FSCFI",
	"FEG_SELECTDOMESTICII": "FSADE2",
	"FEG_SELECTGLOBAL":     "FSAGE",
	"FEG_SELECTGLOBALII":   "FSAGE2",
	"FEG_SELECTLLC":        "FSLLC",
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
