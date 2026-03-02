package executors

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
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
		if unicode.IsDigit(r) || r == '.' || r == '-' || r == '(' || r == ')' {
			result.WriteRune(r)
		}
	}

	res := result.String()
	if res != "" {
		if res[0] == '(' && res[len(res)-1] == ')' {
			res = "-" + res[1:len(res)-1]
		}
	}

	return res
}

// CollectEntities takes a reflect.Value and a store (referenced) and collect the tags and values of a workpaperdata
func CollectEntities(v reflect.Value, store map[string]string) {
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		if value.Kind() == reflect.Struct {
			CollectEntities(value, store)
		}

		entity := field.Tag.Get("entity")
		switch value.Kind() {
		case reflect.Float64:
			store[entity] = strconv.FormatFloat(value.Float(), 'f', -1, 64)
		case reflect.String:
			store[entity] = v.String()
		default:
			continue
		}
	}

}

// CollectAnswers takes a reflect.Value and a store (referenced) and collect the tags and values of formdata

func CollectAnswers(v reflect.Value, store map[[2]string]string) {
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		// dive into nested structs
		if value.Kind() == reflect.Struct &&
			value.Type() != reflect.TypeOf(sql.NullString{}) &&
			value.Type() != reflect.TypeOf(sql.NullFloat64{}) {
			CollectAnswers(value, store)
			continue
		}

		fieldNo := field.Tag.Get("fieldNo")
		lineNo := field.Tag.Get("formLineNo")
		l := [2]string{fieldNo, lineNo}

		switch {
		case value.Type() == reflect.TypeOf(sql.NullString{}):
			ns := value.Interface().(sql.NullString)
			if ns.Valid {
				store[l] = ns.String
			} else {
				store[l] = ""
			}

		case value.Type() == reflect.TypeOf(sql.NullFloat64{}):
			nf := value.Interface().(sql.NullFloat64)
			if nf.Valid {
				store[l] = strconv.FormatFloat(nf.Float64, 'f', -1, 64)
			} else {
				store[l] = ""
			}

		case value.Kind() == reflect.String:
			store[l] = value.String()

		case value.Kind() == reflect.Float64:
			tagValue, ok := field.Tag.Lookup("default")
			if ok && value.Float() == 0 {
				store[l] = tagValue
				continue
			}
			store[l] = strconv.FormatFloat(value.Float(), 'f', -1, 64)

		case value.Kind() == reflect.Bool:
			store[l] = strconv.FormatBool(value.Bool())

		default:
			continue
		}
	}
}
