package model

import (
	// "time"
	"gorm.io/gorm"
)

// user model
type User struct {
	gorm.Model
	Username     string `gorm:"type:varchar(20);not null" json:"username"`
	Password     string `gorm:"size:255;not null" json:"password"`
	Email        string `gorm:"type:varchar(20);not null;unique" json:"email"`
	Age          int    `gorm:"type:int;not null" json:"age"`
	Summary      string `gorm:"type:varchar(255);not null" json:"summary"`
	Avatar_image string `gorm:"type:varchar(255);not null" json:"avatar_image"`
	// CreditCard   []CreditCard
}

// discuss model
type Discuss struct {
	gorm.Model
	// UserID   uint   // 外键，用户id
	Author      string `gorm:"type:varchar(20);not null" json:"author"`
	Title       string `gorm:"type:varchar(20);not null" json:"title"`
	Summary     string `gorm:"type:varchar(20);not null" json:"summaty"`
	Content     string `gorm:"type:varchar(255);not null" json:"content"`
	Like_Number int64  `gorm:"type:int" json:"like_num"`
	View_Number int64  `gorm:"type:int" json:"view_num"`
	Comment     []Comment
}

// commment model
type Comment struct {
	gorm.Model
	DiscussID   uint   // 外键，discussid
	Author      string `gorm:"type:varchar(20);not null" json:"author"`
	Content     string `gorm:"type:varchar(255);not null" json:"content"`
	Like_Number int64  `gorm:"type:int" json:"like_num"`
	View_Number int64  `gorm:"type:int" json:"view_num"`
}

