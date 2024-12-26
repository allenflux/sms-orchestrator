package consts

type EnumsUserStatus int

const (
	EnumUserStatusEnable  EnumsUserStatus = 2
	EnumUserStatusDisable EnumsUserStatus = 1
)

func (e EnumsUserStatus) IsValid() bool {
	return e >= EnumUserStatusDisable && e <= EnumUserStatusEnable
}
