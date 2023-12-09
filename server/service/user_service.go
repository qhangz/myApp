package service

import (
	"errors"
	// "fmt"
	// "encoding/json"
	"math/rand"
	"regexp"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/myApp/dao"
	"github.com/myApp/model"
)

func GetUserByUsername(username string) (*model.User, error) {
	return dao.GetUserByUsername(username)
}

var ErrorRegisterInfo = errors.New("invalid registration information")
var ErrorRegisterUserExist = errors.New("user already exists")
var ErrorRegisterEmailExist = errors.New("email already exists")
var ErrorRegisterFailed = errors.New("registration failed")
var ErrorRegisterSuccess = errors.New("registration success")
var ErrorPasswordWrong = errors.New("password is wrong")
var ErrorEmailWrong = errors.New("email format is wrong")
var ErrorHashPassword = errors.New("hash password failed")

var ErrorLoginInfo = errors.New("invalid login information")
var ErrorLoginFailed = errors.New("login failed")
var ErrorLoginSuccess = errors.New("login success")
var ErrorUserNotExit = errors.New("user not exit")

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
		return ErrorRegisterEmailExist
	}

	// 检查用户名是否已经存在
	if dao.IsUsernameExist(newUser.Username) {
		return ErrorRegisterUserExist
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

// user login
func Login(user model.User) (*model.User, string, error) {
	// 数据验证(username和password)
	if len(user.Username) == 0 || len(user.Password) < 6 {
		return nil, "", ErrorLoginInfo
	}

	// 检查用户名存在
	thisUser, err := dao.GetUserByUsername(user.Username)
	if err != nil {
		return nil, "", ErrorLoginFailed
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
	// 密码解密后返回
	
	return res, nil
	// return dao.GetUserList()
}
