package db

import (
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type TrackedMedia map[string]MediaFile

// The given filename for the YML output database
const DatabaseFilename string = "pic-man.db"

// Database is our complete representation of the managed media metadata
type Database struct {
	Media       TrackedMedia `yaml:"media"`
	LastUpdated int64        `yaml:"updated"`
}

func LoadDatabase(managedDirectory string) (*Database, error) {
	dataFile := filepath.Join(managedDirectory, DatabaseFilename)
	file, err := ioutil.ReadFile(dataFile)
	if err != nil {
		return nil, err
	}
	var result Database
	if err := yaml.Unmarshal(file, &result); err != nil {
		return nil, err
	}
	return &result, err
}
