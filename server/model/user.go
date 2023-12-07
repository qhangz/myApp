package model

import (
	// "time"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"size:255;not null" json:"password"`
	Email    string `gorm:"type:varchar(20);not null;unique" json:"email"`
	Age      int    `gorm:"type:int;not null" json:"age"`
	Summary  string `gorm:"type:varchar(255);not null" json:"summary"`
	Avatar_image string `gorm:"type:varchar(255);not null" json:"avatar_image"`
}

// type User struct {
// 	ID           uint      `gorm:"primary_key" json:"uid"`
// 	Username     string    `gorm:"unique_index" json:"username"`
// 	Password     string    `json:"password"`
// 	Email        string    `json:"email"`
// 	Age          int       `json:"age"`
// 	Summary      string    `json:"summary"`
// 	CreatedAt    time.Time `json:"created_at"`
// 	Avatar_imaeg string    `json:"avatar_image"`
// }
