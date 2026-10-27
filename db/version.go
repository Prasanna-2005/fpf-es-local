package db

import (
	"database/sql"
	"fmt"
	"quantrium-tech/form_pf_extractor_service/config"
)

type lineKey struct {
	FieldNo    string
	FormLineNo string
}

type currentSnapshot struct {
	Value     string
	IsEdit    bool
	CreatedBy string
}

// ResolveTargetVersion picks the version this extraction run will write.
// An unfinished run (rows newer than filing_tracker.currentVersion) is deleted
// so a retry reuses the same version number instead of skipping ahead.
func ResolveTargetVersion(filingID string) (int, error) {
	head, err := readCurrentVersion(filingID)
	if err != nil {
		return 0, err
	}
	maxV, hasMax, err := maxFormVersion(filingID)
	if err != nil {
		return 0, err
	}
	if hasMax && maxV > head {
		if _, err := db.Exec(
			fmt.Sprintf("DELETE FROM %s WHERE filingId = ? AND version > ?", config.TableCfg.FormDataTable),
			filingID, head,
		); err != nil {
			return 0, fmt.Errorf("clearing unfinished version: %w", err)
		}
		maxV, hasMax, err = maxFormVersion(filingID)
		if err != nil {
			return 0, err
		}
	}
	if !hasMax {
		return 1, nil
	}
	return maxV + 1, nil
}

// CommitFilingVersion records a finished extraction as the filing head.
func CommitFilingVersion(filingID string, version int) error {
	res, err := db.Exec(`UPDATE filing_tracker SET currentVersion = ? WHERE id = ?`, version, filingID)
	if err != nil {
		return fmt.Errorf("updating filing currentVersion: %w", err)
	}
	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return fmt.Errorf("filing_tracker row not found for %s", filingID)
	}
	return nil
}

func readCurrentVersion(filingID string) (int, error) {
	var head sql.NullInt64
	err := db.Get(&head, `SELECT currentVersion FROM filing_tracker WHERE id = ?`, filingID)
	if err == sql.ErrNoRows {
		return 0, fmt.Errorf("filing_tracker row not found for %s", filingID)
	}
	if err != nil {
		return 0, fmt.Errorf("reading currentVersion: %w", err)
	}
	if !head.Valid {
		return 0, nil
	}
	return int(head.Int64), nil
}

func maxFormVersion(filingID string) (int, bool, error) {
	var maxV sql.NullInt64
	err := db.Get(&maxV, fmt.Sprintf("SELECT MAX(version) FROM %s WHERE filingId = ?", config.TableCfg.FormDataTable), filingID)
	if err != nil {
		return 0, false, fmt.Errorf("reading max form version: %w", err)
	}
	if !maxV.Valid {
		return 0, false, nil
	}
	return int(maxV.Int64), true, nil
}

func loadVersionCurrents(filingID, fundID string, version int) (map[lineKey]currentSnapshot, error) {
	type row struct {
		FieldNo    string         `db:"fieldNo"`
		FormLineNo string         `db:"formLineNo"`
		Value      string         `db:"value"`
		IsEdit     int            `db:"isEdit"`
		CreatedBy  sql.NullString `db:"createdBy"`
	}
	var rows []row
	err := db.Select(&rows, fmt.Sprintf(`
		SELECT fieldNo, formLineNo, value, isEdit, createdBy
		FROM %s
		WHERE filingId = ? AND fundSeriesId = ? AND version = ? AND isCurrentValue = 1
		ORDER BY createdAt DESC`, config.TableCfg.FormDataTable), filingID, fundID, version)
	if err != nil {
		return nil, fmt.Errorf("loading version %d currents: %w", version, err)
	}
	out := make(map[lineKey]currentSnapshot, len(rows))
	for _, r := range rows {
		k := lineKey{FieldNo: r.FieldNo, FormLineNo: r.FormLineNo}
		if _, seen := out[k]; seen {
			continue
		}
		out[k] = currentSnapshot{Value: r.Value, IsEdit: r.IsEdit != 0, CreatedBy: r.CreatedBy.String}
	}
	return out, nil
}

func loadLatestExtracted(filingID, fundID string, maxVersion int) (map[lineKey]string, error) {
	type row struct {
		FieldNo    string `db:"fieldNo"`
		FormLineNo string `db:"formLineNo"`
		Value      string `db:"value"`
	}
	var rows []row
	err := db.Select(&rows, fmt.Sprintf(`
		SELECT fieldNo, formLineNo, value
		FROM %s
		WHERE filingId = ? AND fundSeriesId = ? AND isEdit = 0 AND version <= ?
		ORDER BY createdAt DESC`, config.TableCfg.FormDataTable), filingID, fundID, maxVersion)
	if err != nil {
		return nil, fmt.Errorf("loading latest extracted values: %w", err)
	}
	out := make(map[lineKey]string, len(rows))
	for _, r := range rows {
		k := lineKey{FieldNo: r.FieldNo, FormLineNo: r.FormLineNo}
		if _, seen := out[k]; seen {
			continue
		}
		out[k] = r.Value
	}
	return out, nil
}
