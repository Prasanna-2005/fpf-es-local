package db

import (
	"fmt"
	"quantrium-tech/form_pf_extractor_service/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func InitialseDB() error {

	cfg := config.Cfg

	connAddr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	_db, err := sqlx.Connect("mysql", connAddr)
	if err != nil {
		return fmt.Errorf("opening remote sql connection: " + err.Error())
	}

	db = _db
	return nil
}
