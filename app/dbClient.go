package app

import (
	"fmt"
	"os"
	"time"

	"github.com/KuSeMi/banking/logger"
	"github.com/jmoiron/sqlx"
)

const (
	maxRetries      = 10
	retryDelay      = 2 * time.Second
	connMaxLifetime = 3 * time.Minute
	maxOpenConns    = 10
	maxIdleConns    = 10
)

func buildDSN() (string, string) {
	host := os.Getenv("MYSQL_HOST")
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	dbname := os.Getenv("MYSQL_DATABASE")
	fullDSN := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", user, password, host, dbname)
	maskedDSN := fmt.Sprintf("%s:***@tcp(%s:3306)/%s", user, host, dbname)

	return fullDSN, maskedDSN
}

func connectWithRetries(driverName, dsn string) (*sqlx.DB, error) {
	var client *sqlx.DB
	var err error

	for attempt := 1; attempt <= maxRetries; attempt++ {
		client, err = sqlx.Open(driverName, dsn)
		if err != nil {
			logger.Error(fmt.Sprintf("Attempt %d: Error opening connection: %v", attempt, err))
			time.Sleep(retryDelay)
			continue
		}

		err = client.Ping()
		if err == nil {
			logger.Info(fmt.Sprintf("Successfully connected on attempt %d", attempt))
			return client, nil
		}

		logger.Error(fmt.Sprintf("Attempt %d: Failed to ping: %v", attempt, err))
		time.Sleep(retryDelay)
	}

	return nil, fmt.Errorf("failed to connect to MySQL after %d attempts: %w", maxRetries, err)
}

func getDbClient() *sqlx.DB {
	fullDSN, maskedDSN := buildDSN()
	logger.Info("Attempting to connect to MySQL with DSN: " + maskedDSN)
	client, err := connectWithRetries("mysql", fullDSN)
	if err != nil {
		logger.Error("Failed to get database client: " + err.Error())
		os.Exit(1)
	}

	client.SetConnMaxLifetime(connMaxLifetime)
	client.SetMaxOpenConns(maxOpenConns)
	client.SetMaxIdleConns(maxIdleConns)

	return client
}
