package db

import (
	"encoding/json"
	"fmt"
	"quantrium-tech/form_pf_extractor_service/config"
	"quantrium-tech/form_pf_extractor_service/logging"
	"time"
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

// currentSnapshot is the version-x row the smart check compares against.
// DecideVersionValue chooses the single row written for one line.
// A nil prior, or a prior that was itself extracted, stores val_extracted.
// An edited prior is kept only when the new extraction matches the latest extracted string.
func DecideVersionValue(extracted string, prior *currentSnapshot, latestExtracted *string) (value string, isEdit bool, preserved bool) {
	if prior == nil || !prior.IsEdit {
		return extracted, false, false
	}
	if latestExtracted != nil && *latestExtracted == extracted {
		return prior.Value, true, true
	}
	return extracted, false, false
}

// PushFormData writes one versioned row per emitted line.
// Older versions are left untouched. targetVersion is chosen once per extraction run.
func PushFormData(filingId string, fundCode string, userId string, formdata map[[2]string]string, resolver SourceResolver, targetVersion int) error {
	if len(formdata) == 0 {
		return nil
	}
	if targetVersion < 1 {
		targetVersion = 1
	}
	fundId := FundCodeMapper[fundCode]
	loc, _ := time.LoadLocation("America/New_York")
	now := time.Now().In(loc)

	var priors map[lineKey]currentSnapshot
	var extracted map[lineKey]string
	if targetVersion > 1 {
		var err error
		priors, err = loadVersionCurrents(filingId, fundId, targetVersion-1)
		if err != nil {
			return err
		}
		for _, row := range priors {
			if row.IsEdit {
				extracted, err = loadLatestExtracted(filingId, fundId, targetVersion-1)
				if err != nil {
					return err
				}
				break
			}
		}
	}

	batch := make([]FormData, 0, len(formdata))
	for key, value := range formdata {
		k := lineKey{FieldNo: key[0], FormLineNo: key[1]}
		var prior *currentSnapshot
		if row, ok := priors[k]; ok {
			copied := row
			prior = &copied
		}
		var latest *string
		if raw, ok := extracted[k]; ok {
			copied := raw
			latest = &copied
		}
		chosen, isEdit, preserved := DecideVersionValue(value, prior, latest)
		sourceFiles := "[]"
		createdBy := userId
		if preserved && prior != nil && prior.CreatedBy != "" {
			createdBy = prior.CreatedBy
		}
		if !preserved {
			encoded, encErr := encodeSourceFiles(resolveSourceFiles(resolver, key[1]))
			if encErr != nil {
				logging.Error(fmt.Sprintf("PushFormData: user_id=%s filing=%s fund=%s line=%s: %v", userId, filingId, fundCode, key[1], encErr))
				encoded = "[]"
			}
			sourceFiles = encoded
		}
		batch = append(batch, FormData{
			FilingId:    filingId,
			FegFundId:   fundId,
			FieldNo:     key[0],
			FormLineNo:  key[1],
			Value:       chosen,
			CreatedAt:   now,
			CreatedBy:   createdBy,
			SourceFiles: sourceFiles,
			Version:     targetVersion,
			IsEdit:      isEdit,
		})
	}

	insertQuery := fmt.Sprintf(`
		INSERT INTO %s (filingId, fundSeriesId, fieldNo, formLineNo, value, sourceFiles, createdAt, createdBy, isCurrentValue, version, isEdit)
		VALUES (:filingId, :fundSeriesId, :fieldNo, :formLineNo, :value, :sourceFiles, :createdAt, :createdBy, 1, :version, :isEdit)`,
		config.TableCfg.FormDataTable)

	if _, err := db.NamedExec(insertQuery, batch); err != nil {
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

func encodeSourceFiles(names []string) (string, error) {
	if names == nil {
		names = []string{}
	}
	raw, err := json.Marshal(names)
	if err != nil {
		return "[]", err
	}
	return string(raw), nil
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
