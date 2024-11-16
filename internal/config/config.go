package config

type AppConfig struct{
	Port string
}

type Config struct{
	App *AppConfig
}

var config Config

func App() *AppConfig {
	return config.App
}

func LoadConfig() {
	setDefaultConfig()
}

func setDefaultConfig() {
	config.App = &AppConfig{
		Port: "8080",
	}
}