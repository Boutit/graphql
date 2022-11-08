package config

type AppConfig struct {
}

type Config struct {
	AppConfig 			AppConfig
}

func GetConfig(env string) Config {
	return Config{
		AppConfig: AppConfig{
		},
	}
}