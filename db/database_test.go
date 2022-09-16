package db

import (
	"fmt"
	"path/filepath"
	"runtime"
	"testing"
)

var testFolderDir string

func init() {
	_, filename, _, _ := runtime.Caller(0)
	fmt.Printf("Filename: %s", filename)
	if filepath.Base(filename) == "database_test.go" {
		testFolderDir = filepath.Dir(filename)
	} else if filepath.Base(filename) != "db" {
		testFolderDir = filepath.Join(testFolderDir, "db")
	} else {
		testFolderDir = filename
	}
	fmt.Printf("Test folder directory for db package: %s\n", testFolderDir)
}

func TestLoadDatabase(t *testing.T) {
	db, err := LoadDatabase(testFolderDir)
	if err != nil {
		t.Errorf("Failure to parse database at %s: %v", testFolderDir, err)
	}
	if len(db.Media) != 3 {
		t.Errorf("Expected 3 media records in sample database, received %d", len(db.Media))
	}
}
