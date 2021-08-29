package config

// Config is app config
type Config struct {
	Port string
	MongoURI string
	DBName string
	BaseURL string
}

// GetConfig returns the app config
func GetConfig() Config {
	var config Config
	config.Port = ":8080"
	config.MongoURI = "mongodb://127.0.0.1:27017"
	config.DBName = "LinkShortener"
	config.BaseURL = "https://www.example.com"
	return config
}
