package dto

type ResponseHealthInfo struct {
	ServiceName    string   `json:"service_name"`
	ServiceVersion string   `json:"service_version"`
	Compiler       string   `json:"compiler"`
	Database       Database `json:"database,omitempty"`
}

type Database struct {
	Mysql bool `json:"mysql"`
}
