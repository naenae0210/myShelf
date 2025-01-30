package db

import (
	"database/sql"
	"fmt"
	"log"
	"myshelf/config"

	_ "github.com/go-sql-driver/mysql"
)

func MySQLConnect(cfg config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.DBPort,
		cfg.DBName,
	)

	dbConn, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
		return nil, err
	}

	if err = dbConn.Ping(); err != nil {
		log.Fatal("Database unreachable:", err)
		return nil, err
	}

	return dbConn, nil
}
