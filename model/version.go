package model

import "os"

//Version struct
type Version struct {
	Version string `json:"version"`
	Host    string `json:"host"`
}

// VersionService struct
type VersionService struct {
	Version *Version
}

// GetVersion - retrieve current version
func (vs *VersionService) GetVersion() (*Version, error) {
	version := Version{}

	//TODO: temp, implement it as global variable
	version.Version = "V1"
	version.Host, _ = os.Hostname()

	return &version, nil
}
