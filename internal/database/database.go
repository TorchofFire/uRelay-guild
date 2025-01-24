package database

import (
	"fmt"

	"github.com/TorchofFire/uRelay-guild/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func NewDbConnectionPool() (*sqlx.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		config.DBUser,
		config.DBPassword,
		config.DBHost,
		"uRelay",
	)

	var db *sqlx.DB
	var err error
	db, err = sqlx.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to the database: %v", err)
	}

	fmt.Println("Database Connection Pool initialized")
	return db, nil
}
