package db

import (
	"fmt"
	"quantrium-tech/form_pf_extractor_service/config"
	"strconv"
	"time"
)

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

	query := fmt.Sprintf("INSERT INTO %s (filingId, fundId, entity, value, createdAt, createdBy) VALUES (:filingId, :fundId, :entity, :value, :createdAt, :createdBy)", config.TableCfg.ExtractionTable)

	_, err := db.NamedExec(query, batch)
	if err != nil {
		return fmt.Errorf("pushing extraction data to db: " + err.Error())
	}

	return nil
}
