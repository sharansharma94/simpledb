package test

import (
	"os"
	"testing"

	"github.com/sharansharma94/simpledb/internal/db"
)

func TestDatabse(t *testing.T) {
	dataDir := "data"
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		t.Errorf("error creating data directory: %v", err)
	}
	testFilePath := "data/test_db.log"
	defer os.Remove(testFilePath)
	defer os.RemoveAll(dataDir)

	db, err := db.NewDatabase(testFilePath)
	if err != nil {
		t.Errorf("error creating database: %v", err)
	}

	err = db.Write("testKey", "testValue", false)
	if err != nil {
		t.Errorf("Write error: %v", err)
	}

	value, err := db.Read("testKey")
	if err != nil || value != "testValue" {
		t.Errorf("Read error: %v", err)
	}

	err = db.Delete("testKey")
	if err != nil {
		t.Errorf("Delete error: %v", err)
	}

	_, err = db.Read("testKey")
	if err == nil {
		t.Errorf("key should not exist")
	}

}
