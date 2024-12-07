package consts

type WsErrType int

const (
	WsConnErr  WsErrType = 1 //ws连接错误
	WsWriteErr WsErrType = 2 //ws推送给客户端错误
	WsReadErr  WsErrType = 3 //ws读取客户端消息报错
)

const (
	EventFacebookLogin         = "facebook_login"          //登录facebook账号
	EventFacebookLoginOut      = "facebook_loginout"       //登出facebook账号
	EventFacebookAccountModify = "facebook_account_modify" //修改facebook账号信息
	EventHasInstallApp         = "has_install_app"         //根据包名判断是否安装app
)

// 客户端响应数据
type ReadMessage struct {
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data"`
}

type WriteMessage struct {
	Event string      `json:"event"`
	Data  interface{} `json:"data"`
}

func GetWriteMessage(event string, data interface{}) *WriteMessage {
	return &WriteMessage{
		Event: event,
		Data:  data,
	}
}

type LoginOutParams struct {
	TaskId      string `json:"task_id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	LoginType   int    `json:"login_type"`
	AccountType int    `json:"account_type"`
	VerifyUrl   string `json:"verify_url"`
}

// 获取-退出登录参数
func WriteLoginOutParams(params *LoginOutParams) *WriteMessage {
	return GetWriteMessage(EventFacebookLoginOut, params)
}

type LoginParams struct {
	TaskId      string `json:"task_id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	LoginType   int    `json:"login_type"`
	AccountType int    `json:"account_type"`
	VerifyUrl   string `json:"verify_url"`
	Code        string `json:"code"`
}

// 获取-登录参数
func WriteLoginParams(params *LoginParams) *WriteMessage {
	return GetWriteMessage(EventFacebookLogin, params)
}

type AccountModifyParams struct {
	TaskId   string `json:"task_id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Cover    string `json:"cover"`
	Sex      int    `json:"sex"`
}

// 获取-修改facebook用户信息参数
func WriteAccountModifyParams(params *AccountModifyParams) *WriteMessage {
	return GetWriteMessage(EventFacebookAccountModify, params)
}

type HasInstallAppParams struct {
	TaskId  string `json:"task_id"`
	AppName string `json:"app_name"`
}

// 获取-是否安装app参数
func WriteHasInstallAppParams(params *HasInstallAppParams) *WriteMessage {
	return GetWriteMessage(EventHasInstallApp, params)
}
