package db

import (
	"database/sql"
	"time"
)

// FileMapper represents the reduced mapper with respect to DB that only holds the neccessary information for execution
type FileMapper struct {
	FilingPath    string `db:"filingPath"`
	FundShortCode string `db:"fundShortCode"`
	QuarterNo     string `db:"quarterNo"`
}

// ExtractionEntity represents the structure of extraction table that stores the extracted Entity
type ExtractionEntity struct {
	FilingId  string    `db:"filingId"`
	FundId    string    `db:"fundSeriesId"`
	Entity    string    `db:"entity"`
	Value     string    `db:"value"`
	CreatedAt time.Time `db:"createdAt"`
	CreatedBy string    `db:"createdBy"`
	UpdatedAt time.Time `db:"updatedAt"`
	UpdatedBy string    `db:"updatedBy"`
}

// FundInfo stores the static identification info for the funds
type FundInfo struct {
	FundName string         `db:"fundName"`
	FundID   string         `db:"fegFundId"`
	FundNFA  sql.NullString `db:"nfaId"`
	FundLEI  sql.NullString `db:"leiId"`
	FundType string         `db:"fundType"`
}

// FormData represents the structure of form_data table that stores the answers/form data from the computations
type FormData struct {
	FilingId   string       `db:"filingId"`
	FegFundId  string       `db:"fundSeriesId"`
	FieldNo    string       `db:"fieldNo"`
	FormLineNo string       `db:"formLineNo"`
	Value      string       `db:"value"`
	Date       sql.NullTime `db:"date"`
	CreatedAt  time.Time    `db:"createdAt"`
	CreatedBy  string       `db:"createdBy"`
	UpdatedAt  time.Time    `db:"updatedAt"`
	UpdatedBy  string       `db:"updatedBy"`
}

var FundStorer = make(map[string]FundInfo)
