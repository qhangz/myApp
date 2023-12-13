package dao

import (
	"github.com/myapp/db"
	"github.com/myapp/model"
)

// publish comment
func PublishComment(newComment model.Comment) error {
	return db.DB.Create(&newComment).Error
}

