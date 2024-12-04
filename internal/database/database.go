package database

import (
	"fmt"
	"log"

	"github.com/TorchofFire/uRelay-guild/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func InitDbConnectionPool() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		config.DBUser,
		config.DBPassword,
		config.DBHost,
		"uRelay",
	)

	var err error
	DB, err = sqlx.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	fmt.Println("Database Connection Pool initialized")
}
