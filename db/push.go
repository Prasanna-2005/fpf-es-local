package db

import (
	"fmt"
	"quantrium-tech/form_pf_extractor_service/config"
	"time"
)

// PushEntities takes the entity-value mapped data and does a batched query to push the extracted entities.
func PushEntities(filingId string, fundCode string, userId string, data map[string]string) error {

	fundId := FundCodeMapper[fundCode]
	loc, _ := time.LoadLocation("Asia/Kolkata")
	filingIdConv := filingId

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
func PushFormData(filingId string, fundCode string, userId string, formdata map[[2]string]string) error {

	fundId := FundCodeMapper[fundCode]
	loc, _ := time.LoadLocation("Asia/Kolkata")
	filingIdConv := filingId

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

		batch = append(batch, fd)
	}

	query := fmt.Sprintf("INSERT INTO %s (filingId, fundSeriesId, fieldNo, formLineNo, value, createdAt, createdBy) VALUES (:filingId, :fundSeriesId, :fieldNo, :formLineNo, :value, :createdAt, :createdBy)", config.TableCfg.FormDataTable)

	_, err := db.NamedExec(query, batch)
	if err != nil {
		return fmt.Errorf("pushing form data to db: " + err.Error())
	}

	return nil
}
