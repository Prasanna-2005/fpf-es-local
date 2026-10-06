package executors

import (
	"strings"

	"quantrium-tech/form_pf_extractor_service/config"
)

// UniqueFiles drops blanks and duplicates while keeping first-seen order.
func UniqueFiles(names ...string) []string {
	var out []string
	seen := map[string]struct{}{}
	for _, name := range names {
		name = strings.TrimSpace(name)
		if name == "" {
			continue
		}
		key := strings.ToLower(name)
		if _, ok := seen[key]; ok {
			continue
		}
		seen[key] = struct{}{}
		out = append(out, name)
	}
	return out
}

func WithStoragePath(path string) string {
	if config.Cfg == nil {
		return path
	}
	return config.Cfg.StoragePath + path
}

func IsLiquidityLine(line string) bool {
	return line == "15" || strings.HasPrefix(line, "16_")
}

func IsBalanceLine(line string, investmentServices bool) bool {
	switch line {
	case "8", "9", "10", "12_a", "12_b", "14_assets_level3", "14_assets_costbased", "14_liabilities_costbased":
		return true
	case "14_assets_level1", "14_assets_level2", "14_liabilities_level2":
		return investmentServices
	default:
		return false
	}
}

// Q17Term reads the Q17 row letter from form lines such as 17_(d)_gross_performance.
func Q17Term(line string) string {
	const prefix = "17_("
	if !strings.HasPrefix(line, prefix) {
		return ""
	}
	rest := line[len(prefix):]
	end := strings.Index(rest, ")")
	if end <= 0 {
		return ""
	}
	return q17Term[strings.ToLower(rest[:end])]
}

var q17Term = map[string]string{
	"a": "january",
	"b": "february",
	"c": "march",
	"d": "first",
	"e": "april",
	"f": "may",
	"g": "june",
	"h": "second",
	"i": "july",
	"j": "august",
	"k": "september",
	"l": "third",
	"m": "october",
	"n": "november",
	"o": "december",
	"p": "fourth",
	"q": "year",
}
