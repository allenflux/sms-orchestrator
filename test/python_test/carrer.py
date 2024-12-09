ProjectUrl = "127.0.0.1"
Port = "8000"
Version = "api/v1"
BaseUrl = "http://" + ProjectUrl + ":" + Port + "/" + Version

UrlRouterMap = [
    {"path":"/career/device", "method":"POST", "dc":"注册设备到平台"},
    {"path":"/career/info", "method":"POST", "dc":"根据设备号获取设备信息"},
    {"path":"/career/task", "method":"POST", "dc":"根据设备号获取设备任务"},
    {"path":"/career/task/result", "method":"POST", "dc":"根据设备号上报任务执行结果"},
    {"path":"/career/task/content", "method":"POST", "dc":"根据设备号上报接收的内容"},
]