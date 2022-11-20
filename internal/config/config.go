package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type GraphQLConfig struct {
	GraphQLHost string
	GraphQLHTTPPort string
}

type Config struct {
	GraphQLConfig GraphQLConfig
}

const (
	graphQLHost = "graphql.host"
	graphQLHTTPPort = "graphql.port"
)

func GetConfig(env string) Config {
	n := fmt.Sprintf("config.%s", env)

	viper.SetConfigName(n)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	
	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println(err)
	}
	
	return Config{
		GraphQLConfig: GraphQLConfig{
			GraphQLHost: viper.GetString(graphQLHost),
			GraphQLHTTPPort: viper.GetString(graphQLHTTPPort),
		},
	}
}