package db

import (
	"encoding/json"
	"fmt"
	"quantrium-tech/form_pf_extractor_service/config"
	"time"

	"github.com/jmoiron/sqlx"
)

// SourceResolver maps a form line to the files that produced its extracted value.
type SourceResolver interface {
	SourceFiles(formLineNo string) []string
}

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

// PushFormData takes the tagged-value mapped data and does a batched query to push the form data.
// resolver is optional; when set, each row stores the source file names for that form line.
func PushFormData(filingId string, fundCode string, userId string, formdata map[[2]string]string, resolver SourceResolver) error {
	fundId := FundCodeMapper[fundCode]
	loc, _ := time.LoadLocation("America/New_York")

	var batch []FormData
	var fieldNos []string
	var lineNos []string

	for key, value := range formdata {
		batch = append(batch, FormData{
			FilingId:    filingId,
			FegFundId:   fundId,
			FieldNo:     key[0],
			FormLineNo:  key[1],
			Value:       value,
			CreatedAt:   time.Now().In(loc),
			CreatedBy:   userId,
			SourceFiles: encodeSourceFiles(resolveSourceFiles(resolver, key[1])),
		})
		fieldNos = append(fieldNos, key[0])
		lineNos = append(lineNos, key[1])
	}

	// safe batch update using sqlx.In to safely expand slices in go
	// This generates: UPDATE table SET isCurrentValue=0 WHERE ... AND fieldNo IN (?,?) ...
	updateQuery, args, err := sqlx.In(fmt.Sprintf(`
		UPDATE %s 
		SET isCurrentValue = 0 
		WHERE filingId = ? 
		  AND fundSeriesId = ? 
		  AND fieldNo IN (?) 
		  AND formLineNo IN (?)`, config.TableCfg.FormDataTable),
		filingId, fundId, fieldNos, lineNos)

	if err != nil {
		return fmt.Errorf("preparing update query: %w", err)
	}

	// sqlx.In returns '?' placeholders, but MySQL requires '?'
	// db.Rebind ensures the query is compatible with your driver
	updateQuery = db.Rebind(updateQuery)
	_, err = db.Exec(updateQuery, args...)
	if err != nil {
		return fmt.Errorf("executing batch update: %w", err)
	}

	//batch insert using NamedExec for the struct slice
	insertQuery := fmt.Sprintf(`
		INSERT INTO %s (filingId, fundSeriesId, fieldNo, formLineNo, value, sourceFiles, createdAt, createdBy, isCurrentValue) 
		VALUES (:filingId, :fundSeriesId, :fieldNo, :formLineNo, :value, :sourceFiles, :createdAt, :createdBy, 1)`,
		config.TableCfg.FormDataTable)

	_, err = db.NamedExec(insertQuery, batch)
	if err != nil {
		return fmt.Errorf("executing batch insert: %w", err)
	}

	return nil
}

func resolveSourceFiles(resolver SourceResolver, formLineNo string) []string {
	if resolver == nil {
		return nil
	}
	return resolver.SourceFiles(formLineNo)
}

func encodeSourceFiles(names []string) string {
	if names == nil {
		names = []string{}
	}
	raw, err := json.Marshal(names)
	if err != nil {
		return "[]"
	}
	return string(raw)
}

// func PushFormData(filingId string, fundCode string, userId string, formdata map[[2]string]string) error {

// 	fundId := FundCodeMapper[fundCode]
// 	loc, _ := time.LoadLocation("Asia/Kolkata")
// 	filingIdConv := filingId

// 	var batch []FormData

// 	for key, value := range formdata {

// 		fd := FormData{
// 			FilingId:  filingIdConv,
// 			FegFundId: fundId,
// 			Value:     value,
// 			CreatedAt: time.Now().In(loc),
// 			CreatedBy: userId,
// 		}

// 		fd.FieldNo = key[0]
// 		fd.FormLineNo = key[1]

// 		batch = append(batch, fd)
// 	}

// 	query := fmt.Sprintf("INSERT INTO %s (filingId, fundSeriesId, fieldNo, formLineNo, value, createdAt, createdBy) VALUES (:filingId, :fundSeriesId, :fieldNo, :formLineNo, :value, :createdAt, :createdBy)", config.TableCfg.FormDataTable)

// 	_, err := db.NamedExec(query, batch)
// 	if err != nil {
// 		return fmt.Errorf("pushing form data to db: " + err.Error())
// 	}

// 	return nil
// }
