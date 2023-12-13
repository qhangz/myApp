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

// publish discuss
func PublishDiscuss(newDiscuss model.Discuss) error {
	return dao.PublishDiscuss(newDiscuss)
}

// get discuss info by discuss id from request
func GetDiscussInfo(discussID uint) (*model.Discuss, error) {
	return dao.GetDiscussInfo(discussID)
}
