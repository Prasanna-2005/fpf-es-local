package db

import "time"

// FileMapper represents the reduced mapper with respect to DB that only holds the neccessary information for execution
type FileMapper struct {
	FilingPath    string `db:"filingPath"`
	FundShortCode string `db:"fundShortCode"`
	QuarterNo     int    `db:"quarterNo"`
}

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

type Fund struct {
}

type FundData struct {
}
