package config

import (
	"github.com/spf13/viper"
)

// func GetDatabaseConfig(){
// 	viper.SetConfigName("config")
// 	viper.SetConfigType("yaml")
// 	viper.AddConfigPath("./config")
// 	err:=viper.ReadInConfig()
// 	if err!=nil{
// 		panic(err)
// 	}
// }

// func InitDB() {
//     var err error
//     dbConfig := config.GetDatabaseConfig()
//     DB, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
//         dbConfig.Username,
//         dbConfig.Password,
//         dbConfig.Host,
//         dbConfig.Port,
//         dbConfig.DBName))
//     if err != nil {
//         panic(err)
//     }
// }

type DatabaseConfig struct {
	Username string
	Password string
	Host     string
	Port     int
	DBName   string
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
		Username: viper.GetString("database.username"),
		Password: viper.GetString("database.password"),
		Host:     viper.GetString("database.host"),
		Port:     viper.GetInt("database.port"),
		DBName:   viper.GetString("database.dbname"),
	}
}
