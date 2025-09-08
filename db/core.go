package db

import (
	"database/sql"
	"time"
)

// FileMapper represents the reduced mapper with respect to DB that only holds the neccessary information for execution
type FileMapper struct {
	FilingPath    string `db:"filingPath"`
	FundShortCode string `db:"fundShortCode"`
	QuarterNo     int    `db:"quarterNo"`
}

// ExtractionEntity represents the structure of extraction table that stores the extracted Entity
type ExtractionEntity struct {
	FilingId  int       `db:"filingId"`
	FundId    string    `db:"fundId"`
	Entity    string    `db:"entity"`
	Value     string    `db:"value"`
	CreatedAt time.Time `db:"createdAt"`
	CreatedBy int       `db:"createdBy"`
	UpdatedAt time.Time `db:"updatedAt"`
	UpdatedBy int       `db:"updatedBy"`
}

// FormData represents the structure of form_data table that stores the answers/form data from the computations
type FormData struct {
	FilingId   int            `db:"filingId"`
	FegFundId  string         `db:"fegFundId"`
	FieldNo    string         `db:"fieldNo"`
	FormLineNo sql.NullString `db:"formLineNo"`
	Value      string         `db:"value"`
	Date       sql.NullTime   `db:"date"`
	CreatedAt  time.Time      `db:"createdAt"`
	CreatedBy  int            `db:"createdBy"`
	UpdatedAt  time.Time      `db:"updatedAt"`
	UpdatedBy  int            `db:"updatedBy"`
}
