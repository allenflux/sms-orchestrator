package model

type UserPermission struct {
	Id       int    `json:"id"         orm:"id"         description:""`        //
	Function string `json:"function"   orm:"function"   description:"功能名称"`    // 功能名称
	Redirect string `json:"redirect"   orm:"redirect"   description:"路由重定向地址"` // 路由重定向地址
}
