package common

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// ErrorCode 错误码
type ErrorCode int32

const (
	_ int32 = iota + 9999
	// StatusOK : 正常
	StatusOK
	// StatusParamInvalid : 请求参数无效
	StatusParamInvalid
	// StatusServerError : 服务出错
	StatusServerError
	// StatusRegisterFailed : 注册失败
	StatusRegisterFailed
	// StatusLoginFailed : 登录失败
	StatusLoginFailed
	// StatusTokenInvalid : 10005 Token无效
	StatusTokenInvalid
	// FileAlreadExists : 10006 文件已存在
	FileAlreadExists
	// StatusUserNotExists : 10007 用户不存在
	StatusUserNotExists
	//TODO session age
	SessionAge = 36000
	//TODO 口令
	LoginPassWord = "yourpassword"
	LoginPwdSalt  = "yoursalt"
)

//TODO 鉴权存储
var Token = "-3.6295141"
var Db *gorm.DB

func init() {
	var err error
	Db, err = gorm.Open(sqlite.Open("./db/test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}
