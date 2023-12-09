package dao

import (
	"github.com/myApp/db"
	"github.com/myApp/model"
)

// get user by username
func GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	if err := db.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// get user by email
func GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	if err := db.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// get user by id
func GetUserById(id int) (*model.User, error) {
	var user model.User
	if err := db.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// is email exist
func IsEmailExist(email string) bool {
	var user model.User
	db.DB.Where("email = ?", email).First(&user)
	return user.ID != 0
}

// is username exist
func IsUsernameExist(username string) bool {
	var user model.User
	db.DB.Where("username = ?", username).First(&user)
	return user.ID != 0
}

// user register(create new user)
func Register(user model.User) error {
	return db.DB.Create(&user).Error
}

// user list
func GetUserList() ([]model.User, error) {
	var users []model.User
	// 返回所有用户的信息
	if err := db.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
