package config

import (
	"github.com/spf13/viper"
)


type DatabaseConfig struct {
	DriveName string
	Username  string
	Password  string
	Host      string
	Port      int
	DBName    string
	Charset   string
}

func GetDatabaseConfig() *DatabaseConfig {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	return &DatabaseConfig{
		DriveName: viper.GetString("database.driveName"),
		Username:  viper.GetString("database.username"),
		Password:  viper.GetString("database.password"),
		Host:      viper.GetString("database.host"),
		Port:      viper.GetInt("database.port"),
		DBName:    viper.GetString("database.dbname"),
		Charset:   viper.GetString("database.charset"),
	}
}
