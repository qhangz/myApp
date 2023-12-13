package service

import (
	// "errors"
	// "fmt"
	// "encoding/json"
	// "math/rand"
	// "regexp"
	// "strconv"
	// "time"

	// "golang.org/x/crypto/bcrypt"

	"github.com/myapp/dao"
	"github.com/myapp/model"
)

// publish comment
func PublishComment(newComment model.Comment) error {
	return dao.PublishComment(newComment)
}
