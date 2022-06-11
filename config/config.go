package config

import "github.com/stakkato95/service-engineering-go-lib/config"

type Config struct {
	GrpcPort   string `mapstructure:"GRPC_PORT"`
	ServerPort string `mapstructure:"SERVER_PORT"`
	JwtSecret  string `mapstructure:"JWT_SECRET"`
	DbDriver   string `mapstructure:"DB_DRIVER"`
	DbSource   string `mapstructure:"DB_SOURCE"`
	DbName     string `mapstructure:"DB_NAME"`
}

var AppConfig Config

func init() {
	config.Init(&AppConfig, Config{})
}

func GetConnectionString() string {
	return AppConfig.DbSource + "/" + AppConfig.DbName
}
