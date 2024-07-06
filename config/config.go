package config

type DatabaseConfig struct {
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
	Port     int    `json:"port"`
}

type DBConfig struct {
	Database DatabaseConfig `json:"database"`
}
