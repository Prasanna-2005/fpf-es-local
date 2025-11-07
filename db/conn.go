package db

import (
	"fmt"
	"quantrium-tech/form_pf_extractor_service/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

// InitialseDB opens up the connection and pings the connection via sqlx.Connect.
// Stores the connection in the unexported variable
func InitialseDB() error {

	cfg := config.Cfg
	fmt.Printf("cfg: %v\n", cfg)

	connAddr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	fmt.Println(connAddr)
	dbConn, err := sqlx.Connect("mysql", connAddr)
	if err != nil {
		return fmt.Errorf("opening remote sql connection: " + err.Error())
	}

	db = dbConn
	return nil
}
