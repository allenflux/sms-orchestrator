package consts

import (
	"fmt"
	"github.com/spf13/cast"
)

type CodeType = int //手机验证码类型

const (
	CodeRegister       = 1 //注册
	CodeForgetPassword = 2 //忘记密码
)

// 获取类型名称
func GetCodeTypeName(codeType CodeType) string {
	switch codeType {
	case CodeRegister:
		return "Register"
	case CodeForgetPassword:
		return "ForgetPassword"
	}
	return cast.ToString(codeType)
}

// 获取类型超时时间
func GetCodeTypeTimeout(codeType CodeType) int {
	switch codeType {
	case CodeRegister:
		return 300
	case CodeForgetPassword:
		return 300
	}
	return 300
}

func GetCodeRedisKey(codeType CodeType, mobile string) string {
	return fmt.Sprintf("%v-%v", GetCodeTypeName(codeType), mobile)
}
