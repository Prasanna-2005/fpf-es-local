package executors

import (
	"fmt"
	"strings"
	"unicode"
)

// IndexToExcel converts (x,y) coordinate to a excel cell reference
// (0,1) will translates to B1
func IndexToExcel(row, col int) string {
	colName := ""
	for col >= 0 {
		colName = string(rune(col%26+'A')) + colName
		col = col/26 - 1
	}
	return fmt.Sprintf("%s%d", colName, row+1)
}

// CleanNonNumeric takes a string and removes non-numeric characters except "." to hold the decimals
func CleanNonNumeric(s string) string {
	var result strings.Builder
	for _, r := range s {
		if unicode.IsDigit(r) || r == '.' {
			result.WriteRune(r)
		}
	}
	return result.String()
}
