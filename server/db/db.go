package db

import (
    "fmt"
    "github.com/jinzhu/gorm"
    // _ "github.com/jinzhu/gorm/dialects/mysql"
    _ "github.com/go-sql-driver/mysql"

    "github.com/myApp/config"
    "github.com/myApp/model"
)

var DB *gorm.DB

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
// 		panic("failed to connect database,err:" + err.Error())
//     }
    
// }

func InitDB() *gorm.DB {
    dbConfig := config.GetDatabaseConfig()
    driveName := "mysql"
    // driveName := dbConfig.DriveName
    host := dbConfig.Host
    port := dbConfig.Port
    database := dbConfig.DBName
    username := dbConfig.Username
    password := dbConfig.Password
    charset := "utf8"


	args := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", username, password, host, port, database, charset)
	db, err := gorm.Open(driveName, args)
	if err != nil {
		panic("failed to connect database,err:" + err.Error())
	}
	//defer db.Close()
	db.AutoMigrate(&model.User{})
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
