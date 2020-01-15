package gorm

import (
	"os"
	"testing"
)

func TestOpen(t *testing.T) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_DB_PASSWORD")

	if host == "" || port == "" || dbname == "" || username == "" || password == "" {
		t.Log("missing environment variables, bypassing this unit")
		return
	}

	db, err := Open("localhost", "3306", "test", "root", "root")
	if err != nil {
		t.Logf("failed to open: %v", err)
		return
	}
	err = db.DB().Ping()
	if err != nil {
		t.Logf("failed to ping: %v", err)
	} else {
		t.Log("pong")
	}
}
