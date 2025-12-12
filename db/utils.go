package db

import (
	"database/sql"
	"time"
)

// ToNullString takes up a string and converts it to DB supported NullString for Nullable Fields
func ToNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{String: s, Valid: true}
}

// ToNullTime takes up a time and converts it to DB supported NullTime for Nullable Fields
func ToNullTime(t *time.Time) sql.NullTime {
	if t == nil {
		return sql.NullTime{Valid: false}
	}
	return sql.NullTime{Time: *t, Valid: true}
}
