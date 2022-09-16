package db

type TrackedMedia map[string]MediaFile

type Database struct {
	Media       TrackedMedia `yaml:"media"`
	LastUpdated int64        `yaml:"updated"`
}
