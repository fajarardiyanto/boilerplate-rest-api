package dto

type ResponseHealthInfo struct {
	ServiceName    string   `json:"service_name"`
	ServiceVersion string   `json:"service_version"`
	Database       Database `json:"database"`
}

type Database struct {
	Mysql bool `json:"mysql"`
}
