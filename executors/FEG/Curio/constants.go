package curio

var LevelingEntityMapper = map[string]string{
	"invt hedge funds @ cost": "invtHedgeFundsCost",
	"invt in hedge funds u/r": "invtInHedgeFundsUR",
	"private eq @ cost":       "privateEqCost",
	"private eq u/r g/l":      "privateEqURGL",
}

// this l3 will be included accordingly into sec1b abscissa
var LevelMapper = map[string]map[string][]string{

	"CCPTE": {
		"l3": {"invt hedge funds @ cost", "invt in hedge funds u/r"},
	},
	"CADTE": {
		"l3": {"invt hedge funds @ cost", "invt in hedge funds u/r"},
	},

	"CURIV": {
		"l3": {"private eq @ cost", "private eq u/r g/l"},
	},
	"CURTE": {
		"l3": {"private eq @ cost", "private eq u/r g/l"},
	},
	"CUADX": {
		"l3": {"private eq @ cost", "private eq u/r g/l"},
	},
	"CURCCT": {
		"l3": {"private eq @ cost", "private eq u/r g/l"},
	},
	"CURCCO": {
		"l3": {"private eq @ cost", "private eq u/r g/l"},
	},
	"CUKBX": {
		"l3": {"private eq @ cost", "private eq u/r g/l"},
	},
	"CURCP": {
		"l3": {"private eq @ cost", "private eq u/r g/l", "invt hedge funds @ cost", "invt in hedge funds u/r", "invt in hedge funds u/r g/l"},
	},
	"COPTE": {
		"l3": {"private eq @ cost", "private eq u/r g/l"},
	},
	"CUROP": {
		"l3": {"private eq @ cost", "private eq u/r g/l"},
	},
	"CURTC": {
		"l3": {"private eq @ cost", "private eq u/r g/l"},
	},
	"CURTB": {
		"l3": {"private eq @ cost", "private eq u/r g/l"},
	},
	"CURTIM": {
		"l3": {"private eq @ cost", "private eq u/r g/l"},
	},
	"CUPDC": {
		"l3": {"private eq @ cost", "private eq u/r g/l"},
	},
	"CUELT": {
		"l3": {"private eq @ cost", "private eq u/r g/l"},
	},
	"CURHP": {
		"l3": {"private eq @ cost", "private eq u/r g/l"},
	},
	"CUTRV": {
		"l3": {"private eq @ cost", "private eq u/r g/l"},
	},
}
