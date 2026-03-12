package privateinvestors

// ShortCodeConvertor converts the shortcode from db to code used in files
var ShortCodeConvertor = map[string]string{

	// Roman Numeral Mappings
	"FEGPOF I":   "FEGOF_1",
	"FEGPOF II":  "FEGOF_2",
	"FEGPOF III": "FEGOF_3",
	"FEGPOF IV":  "FEGOF_4",
	"FEGPOF V":   "FEGOF_5",
	"FEGPOF VI":  "FEGOF_6",
	"FEGPOF VII": "FEGOF_7",

	// Normal Mappings
	"FEGPOF 1": "FEGOF_1",
	"FEGPOF 2": "FEGOF_2",
	"FEGPOF 3": "FEGOF_3",
	"FEGPOF 4": "FEGOF_4",
	"FEGPOF 5": "FEGOF_5",
	"FEGPOF 6": "FEGOF_6",
	"FEGPOF 7": "FEGOF_7",
}

// FundtoSeriesMapper maps fund to list of series attached to that particular fund
var FundtoSeriesMapper = map[string][]string{

	"FEGOF_1": {""},
	"FEGOF_2": {"A", "B", "C"},
	"FEGOF_3": {"A", "B", "C", "D"},
	"FEGOF_4": {"A", "B", "C", "D"},
	"FEGOF_5": {"A", "B", "C", "C AIV", "D"},
	"FEGOF_6": {"A", "A_ESG", "B", "C", "C_ESG", "D"},
	"FEGOF_7": {"B", "C", "C / ESG", "D", "E"},
}

var MaturedFunds = []string{
	"FEGOF_1",
	"FEGOF_2_A",
	"FEGOF_2_B",
	"FEGOF_2_C",
	"FEGOF_3_C",
	"FEGOF_3_D",
}

var combineFunds = map[string]string{
	"FEGOF_5_C AIV": "FEGOF_5_C",
}
