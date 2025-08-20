package executors

import (
	"fmt"
	"strings"
	"unicode"
)

func IndexToExcel(row, col int) string {
	colName := ""
	for col >= 0 {
		colName = string(rune(col%26+'A')) + colName
		col = col/26 - 1
	}
	return fmt.Sprintf("%s%d", colName, row+1)
}

func CleanNonNumeric(s string) string {
	var result strings.Builder
	for _, r := range s {
		if unicode.IsDigit(r) {
			result.WriteRune(r)
		}
	}
	return result.String()
}
