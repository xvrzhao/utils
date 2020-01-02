package gorm

import "testing"

func TestOpen(t *testing.T) {
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
