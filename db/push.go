package db

import (
	"fmt"
	"quantrium-tech/form_pf_extractor_service/config"
	"strconv"
	"strings"
	"time"
)

// PushEntities takes the entity-value mapped data and does a batched query to push the extracted entities.
func PushEntities(filingId string, fundCode string, userId int, data map[string]string) error {

	fundId := FundCodeMapper[fundCode]
	loc, _ := time.LoadLocation("Asia/Kolkata")
	filingIdConv, _ := strconv.Atoi(filingId)

	var batch []ExtractionEntity

	for key, value := range data {
		entity := ExtractionEntity{
			FilingId:  filingIdConv,
			FundId:    fundId,
			Entity:    key,
			Value:     value,
			CreatedAt: time.Now().In(loc),
			CreatedBy: userId,
		}
		batch = append(batch, entity)
	}

	query := fmt.Sprintf("INSERT INTO %s (filingId, fundSeriesId, entity, value, createdAt, createdBy) VALUES (:filingId, :fundSeriesId, :entity, :value, :createdAt, :createdBy)", config.TableCfg.ExtractionTable)

	_, err := db.NamedExec(query, batch)
	if err != nil {
		return fmt.Errorf("pushing extraction data to db: " + err.Error())
	}

	return nil
}

// PushFormData takes the tagged-value mapped data and does a batched query to push the form data
func PushFormData(filingId string, fundCode string, userId int, formdata map[[3]string]string) error {

	fundId := FundCodeMapper[fundCode]
	loc, _ := time.LoadLocation("Asia/Kolkata")
	filingIdConv, _ := strconv.Atoi(filingId)

	var batch []FormData

	for key, value := range formdata {

		fd := FormData{
			FilingId:  filingIdConv,
			FegFundId: fundId,
			Value:     value,
			CreatedAt: time.Now().In(loc),
			CreatedBy: userId,
		}

		fd.FieldNo = key[0]
		fd.FormLineNo = key[1]

		if key[2] != "" {
			parts := strings.Split(key[2], "/")
			month, _ := strconv.Atoi(parts[0])
			day, _ := strconv.Atoi(parts[1])
			year := time.Now().Year() - 1

			date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
			fd.Date = ToNullTime(&date)
		} else {
			fd.Date = ToNullTime(nil)
		}

		batch = append(batch, fd)
	}

	query := fmt.Sprintf("INSERT INTO %s (filingId, fundSeriesId, fieldNo, formLineNo, value, date, createdAt, createdBy) VALUES (:filingId, :fundSeriesId, :fieldNo, :formLineNo, :value, :date, :createdAt, :createdBy)", config.TableCfg.FormDataTable)

	_, err := db.NamedExec(query, batch)
	if err != nil {
		return fmt.Errorf("pushing form data to db: " + err.Error())
	}

	return nil
}
