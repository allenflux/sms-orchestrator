package consts

type User struct {
	Data interface{}
	Type EnumUserType
}

type EnumUserType int

const (
	EnumAdminUser EnumUserType = 1
	EnumEntUser   EnumUserType = 2
)

func GetUser(userData interface{}, userType EnumUserType) User {
	return User{
		Data: userData,
		Type: userType,
	}
}

var SuperUserPermission = []string{"*/*/*"}
