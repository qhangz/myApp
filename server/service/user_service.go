package service

import (
	"errors"
	// "log"
	// "fmt"
	// "encoding/json"
	"math/rand"
	"regexp"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/myapp/dao"
	"github.com/myapp/model"
)

func GetUserByUsername(username string) (*model.User, error) {
	return dao.GetUserByUsername(username)
}

var ErrorRegisterInfo = errors.New("invalid registration information")
var ErrorUserExist = errors.New("user already exists")
var ErrorEmailExist = errors.New("email already exists")
var ErrorRegisterFailed = errors.New("registration failed")
var ErrorRegisterSuccess = errors.New("registration success")
var ErrorPasswordWrong = errors.New("password is wrong")
var ErrorEmailWrong = errors.New("email format is wrong")
var ErrorHashPassword = errors.New("hash password failed")

var ErrorLoginInfo = errors.New("invalid login information")
var ErrorLoginFailed = errors.New("login failed")
var ErrorLoginSuccess = errors.New("login success")
var ErrorUserNotExit = errors.New("user not exit")

var ErrorMsg = errors.New("error message")
var ErrorAgeWrong = errors.New("age is wrong")
var ErrorSummaryWrong = errors.New("summary is wrong")

// is email valid
func validateEmail(email string) bool {
	// 正则表达式匹配邮箱格式
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`
	match, _ := regexp.MatchString(regex, email)
	return match
}

// user register
func randomPick(arr []string) string {
	// 设置随机种子
	rand.Seed(time.Now().UnixNano())

	// 生成随机索引
	index := rand.Intn(len(arr))

	// 返回随机选取的元素
	return arr[index]
}
func Register(newUser model.User) error {
	// 数据验证
	if len(newUser.Username) == 0 || len(newUser.Password) < 6 || len(newUser.Email) == 0 {
		return ErrorRegisterInfo
	}

	if !validateEmail(newUser.Email) {
		return ErrorEmailWrong
	}

	// 检查邮箱是否已经存在
	if dao.IsEmailExist(newUser.Email) {
		return ErrorEmailExist
	}

	// 检查用户名是否已经存在
	if dao.IsUsernameExist(newUser.Username) {
		return ErrorUserExist
	}

	// 用户密码加密
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return ErrorHashPassword
	}
	newUser.Password = string(hashPassword)

	// 初始化summary和avatar_image
	newUser.Summary = "This is a new use"
	// newUser.Avatar_image = "https://s2.loli.net/2023/12/07/ObzivQC2fqpTZlx.jpg"
	AvatarList := []string{
		"https://s2.loli.net/2023/12/07/ObzivQC2fqpTZlx.jpg",
		"https://s2.loli.net/2023/12/05/hRQyPjI723LbV9l.png",
		"https://s2.loli.net/2023/12/09/Bg7Wbvs96INl8Sr.jpg",
		"https://s2.loli.net/2023/12/09/9ZPCdiWGLyImfph.png",
		"https://s2.loli.net/2023/12/09/LHT3gx958KJhMBD.jpg",
		"https://s2.loli.net/2023/12/09/RwhL42gUmV1vaEH.jpg",
		"https://s2.loli.net/2023/12/09/8gZxsycdmFjpDC2.jpg",
		"https://s2.loli.net/2023/12/09/3DdwZM5XlrCF16u.jpg",
	}
	newUser.Avatar_image = randomPick(AvatarList)

	return dao.Register(newUser)
}

// user login (return userInfo,token,error)
func Login(user model.User) (*model.User, string, error) {
	// 数据验证(username和password)
	if len(user.Username) == 0 || len(user.Password) < 6 {
		return nil, "", ErrorLoginInfo
	}

	// 检查用户名存在
	thisUser, err := dao.GetUserByUsername(user.Username)
	if err != nil {
		return nil, "", ErrorUserNotExit
	}
	if thisUser == nil {
		return nil, "", ErrorUserNotExit
	}

	// 检查密码是否正确
	err = bcrypt.CompareHashAndPassword([]byte(thisUser.Password), []byte(user.Password))
	if err != nil {
		return nil, "", ErrorPasswordWrong
	}

	// token
	token := strconv.FormatUint(uint64(thisUser.ID), 10)

	return thisUser, token, nil
}

// user list	返回用户信息列表
func GetUserList() ([]model.User, error) {
	var res []model.User
	res, err := dao.GetUserList()
	if err != nil {
		return nil, err
	}

	return res, nil
	// return dao.GetUserList()
}

// update username
func UpdateUsername(user model.User, newUsername string) error {
	// 检查用户名存在
	thisUser, err := dao.GetUserByUsername(user.Username)
	if err != nil {
		return ErrorUserNotExit
	}
	if thisUser == nil {
		return ErrorUserNotExit
	}
	// 检查新用户名是否已经存在
	if dao.IsUsernameExist(newUsername) {
		return ErrorUserExist
	}
	return dao.UpdateUsername(thisUser, newUsername)
}

// update password
func UpdatePassword(user model.User, newPassword string) error {
	// 检查用户名存在
	thisUser, err := dao.GetUserByUsername(user.Username)
	if err != nil {
		return ErrorUserNotExit
	}
	if thisUser == nil {
		return ErrorUserNotExit
	}
	// 验证密码是否正确
	err = bcrypt.CompareHashAndPassword([]byte(thisUser.Password), []byte(user.Password))
	if err != nil {
		return ErrorMsg
	}
	// 验证新密码长度
	if len(newPassword) < 6 {
		return ErrorRegisterInfo
	}
	// 密码加密
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return ErrorHashPassword
	}
	return dao.UpdatePassword(thisUser, string(hashPassword))
}

// update email
func UpdateEmail(user model.User, newEmail string) error {
	// 检查用户名存在
	thisUser, err := dao.GetUserByUsername(user.Username)
	if err != nil {
		return ErrorUserNotExit
	}
	if thisUser == nil {
		return ErrorUserNotExit
	}
	// 验证邮箱格式
	if !validateEmail(newEmail) {
		return ErrorEmailWrong
	}
	// 检查邮箱是否已经存在
	if dao.IsEmailExist(newEmail) {
		return ErrorEmailExist
	}
	return dao.UpdateEmail(thisUser, newEmail)
}

// update age
func UpdateAge(user model.User, newAge int) error {
	// 检查用户名存在
	thisUser, err := dao.GetUserByUsername(user.Username)
	if err != nil {
		return ErrorUserNotExit
	}
	if thisUser == nil {
		return ErrorUserNotExit
	}
	// 验证年龄是否合法
	if newAge < 0 || newAge > 150 {
		return ErrorAgeWrong
	}

	return dao.UpdateAge(thisUser, newAge)
}

// update summary
func UpdateSummary(user model.User, newSummary string) error {
	// 检查用户名存在
	thisUser, err := dao.GetUserByUsername(user.Username)
	if err != nil {
		return ErrorUserNotExit
	}
	if thisUser == nil {
		return ErrorUserNotExit
	}
	// 验证简介是否合法
	if len(newSummary) <= 0 || len(newSummary) > 100 {
		return ErrorSummaryWrong
	}
	return dao.UpdateSummary(thisUser, newSummary)
}

// delete user
func DeleteUser(user model.User) error {
	// 检查用户名存在
	thisUser, err := dao.GetUserByUsername(user.Username)
	if err != nil {
		return ErrorUserNotExit
	}
	if thisUser == nil {
		return ErrorUserNotExit
	}
	return dao.DeleteUser(thisUser)
}