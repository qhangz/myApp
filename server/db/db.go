package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/myapp/config"
	"github.com/myapp/model"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	dbConfig := config.GetDatabaseConfig()
	// driveName := dbConfig.DriveName
	host := dbConfig.Host
	port := dbConfig.Port
	database := dbConfig.DBName
	username := dbConfig.Username
	password := dbConfig.Password
	charset := dbConfig.Charset

	args := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", username, password, host, port, database, charset)
	//连接MYSQL
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		panic("failed to connect database,err:" + err.Error())
	}

	//延时关闭数据库连接
	// defer db.Close()
	db.AutoMigrate(&model.User{},&model.Discuss{},&model.Comment{})

	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
