package dao

import (
	"github.com/myapp/db"
	"github.com/myapp/model"
)

// publish discuss
func PublishDiscuss(newDiscuss model.Discuss) error {
	return db.DB.Create(&newDiscuss).Error
}

// get comment of this discuss
// func GetComment(discussID uint) (*model.Comment, error) {
// 	discuss := model.Discuss{}
// 	db.DB.Where("id = ?", discussID).First(&discuss)
// 	fmt.Println(discuss)

// 	db.DB.Model(&discuss).Association("Comment").Find(&discuss.Comment)

// 	return &discuss.Comment[discussID], nil
// }

// get discuss info by discuss id from request
func GetDiscussInfo(discussID uint) (*model.Discuss, error) {
	discuss := model.Discuss{}
	db.DB.Where("id = ?", discussID).First(&discuss)

	db.DB.Model(&discuss).Association("Comment").Find(&discuss.Comment)

	return &discuss, nil
}

