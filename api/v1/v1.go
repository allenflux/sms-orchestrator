// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package v1

import (
	"context"

	"sms_backend/api/v1/allUser"
	"sms_backend/api/v1/career"
	"sms_backend/api/v1/common"
	"sms_backend/api/v1/log"
	"sms_backend/api/v1/role"
	"sms_backend/api/v1/sms"
	"sms_backend/api/v1/subUser"
	"sms_backend/api/v1/system"
	"sms_backend/api/v1/user"
)

type IV1AllUser interface {
	Login(ctx context.Context, req *allUser.LoginReq) (res *allUser.LoginRes, err error)
	SubLogin(ctx context.Context, req *allUser.SubLoginReq) (res *allUser.SubLoginRes, err error)
	Logout(ctx context.Context, req *allUser.LogoutReq) (res *allUser.LogoutRes, err error)
	ChangePassword(ctx context.Context, req *allUser.ChangePasswordReq) (res *allUser.ChangePasswordRes, err error)
}

type IV1Career interface {
	Register(ctx context.Context, req *career.RegisterReq) (res *career.RegisterRes, err error)
	FetchTask(ctx context.Context, req *career.FetchTaskReq) (res *career.FetchTaskRes, err error)
	ReportTaskResult(ctx context.Context, req *career.ReportTaskResultReq) (res *career.ReportTaskResultRes, err error)
	ReportReceiveContent(ctx context.Context, req *career.ReportReceiveContentReq) (res *career.ReportReceiveContentRes, err error)
}

type IV1Common interface {
	Captcha(ctx context.Context, req *common.CaptchaReq) (res *common.CaptchaRes, err error)
	FileUpload(ctx context.Context, req *common.FileUploadReq) (res *common.FileUploadRes, err error)
}

type IV1Log interface {
	List(ctx context.Context, req *log.ListReq) (res *log.ListRes, err error)
}

type IV1Role interface {
	List(ctx context.Context, req *role.ListReq) (res *role.ListRes, err error)
	GetPermission(ctx context.Context, req *role.GetPermissionReq) (res *role.GetPermissionRes, err error)
	Created(ctx context.Context, req *role.CreatedReq) (res *role.CreatedRes, err error)
	Updated(ctx context.Context, req *role.UpdatedReq) (res *role.UpdatedRes, err error)
	Deleted(ctx context.Context, req *role.DeletedReq) (res *role.DeletedRes, err error)
}

type IV1Sms interface {
	ProjectList(ctx context.Context, req *sms.ProjectListReq) (res *sms.ProjectListRes, err error)
	ProjectListForFront(ctx context.Context, req *sms.ProjectListForFrontReq) (res *sms.ProjectListForFrontRes, err error)
	ProjectCreate(ctx context.Context, req *sms.ProjectCreateReq) (res *sms.ProjectCreateRes, err error)
	ProjectUpdate(ctx context.Context, req *sms.ProjectUpdateReq) (res *sms.ProjectUpdateRes, err error)
	ProjectDelete(ctx context.Context, req *sms.ProjectDeleteReq) (res *sms.ProjectDeleteRes, err error)
	DeviceList(ctx context.Context, req *sms.DeviceListReq) (res *sms.DeviceListRes, err error)
	AllocateDevice2Project(ctx context.Context, req *sms.AllocateDevice2ProjectReq) (res *sms.AllocateDevice2ProjectRes, err error)
	AllocateAccount2Project(ctx context.Context, req *sms.AllocateAccount2ProjectReq) (res *sms.AllocateAccount2ProjectRes, err error)
	TaskList(ctx context.Context, req *sms.TaskListReq) (res *sms.TaskListRes, err error)
	TaskRecord(ctx context.Context, req *sms.TaskRecordReq) (res *sms.TaskRecordRes, err error)
	ConversationList(ctx context.Context, req *sms.ConversationListReq) (res *sms.ConversationListRes, err error)
	ConversationRecord(ctx context.Context, req *sms.ConversationRecordReq) (res *sms.ConversationRecordRes, err error)
	PostConversationRecord(ctx context.Context, req *sms.PostConversationRecordReq) (res *sms.PostConversationRecordRes, err error)
	TaskDevices(ctx context.Context, req *sms.TaskDevicesReq) (res *sms.TaskDevicesRes, err error)
	SubTaskDevices(ctx context.Context, req *sms.SubTaskDevicesReq) (res *sms.SubTaskDevicesRes, err error)
	PendingTask(ctx context.Context, req *sms.PendingTaskReq) (res *sms.PendingTaskRes, err error)
	SubPendingTask(ctx context.Context, req *sms.SubPendingTaskReq) (res *sms.SubPendingTaskRes, err error)
	SubDeviceList(ctx context.Context, req *sms.SubDeviceListReq) (res *sms.SubDeviceListRes, err error)
	SubGroupList(ctx context.Context, req *sms.SubGroupListReq) (res *sms.SubGroupListRes, err error)
	SubCreateGroup(ctx context.Context, req *sms.SubCreateGroupReq) (res *sms.SubCreateGroupRes, err error)
	SubUpdateGroup(ctx context.Context, req *sms.SubUpdateGroupReq) (res *sms.SubUpdateGroupRes, err error)
	SubDeleteGroup(ctx context.Context, req *sms.SubDeleteGroupReq) (res *sms.SubDeleteGroupRes, err error)
	AllocateDevice2Group(ctx context.Context, req *sms.AllocateDevice2GroupReq) (res *sms.AllocateDevice2GroupRes, err error)
	SubProjectListForFront(ctx context.Context, req *sms.SubProjectListForFrontReq) (res *sms.SubProjectListForFrontRes, err error)
	SubTaskList(ctx context.Context, req *sms.SubTaskListReq) (res *sms.SubTaskListRes, err error)
	SubTaskCreate(ctx context.Context, req *sms.SubTaskCreateReq) (res *sms.SubTaskCreateRes, err error)
	TaskFileDownload(ctx context.Context, req *sms.TaskFileDownloadReq) (res *sms.TaskFileDownloadRes, err error)
	SubTaskDelete(ctx context.Context, req *sms.SubTaskDeleteReq) (res *sms.SubTaskDeleteRes, err error)
	SubTaskRecord(ctx context.Context, req *sms.SubTaskRecordReq) (res *sms.SubTaskRecordRes, err error)
	SubGetConversationRecord(ctx context.Context, req *sms.SubGetConversationRecordReq) (res *sms.SubGetConversationRecordRes, err error)
	SubGetConversationRecordList(ctx context.Context, req *sms.SubGetConversationRecordListReq) (res *sms.SubGetConversationRecordListRes, err error)
	SubPostConversationRecord(ctx context.Context, req *sms.SubPostConversationRecordReq) (res *sms.SubPostConversationRecordRes, err error)
}

type IV1SubUser interface {
	SubRegister(ctx context.Context, req *subUser.SubRegisterReq) (res *subUser.SubRegisterRes, err error)
	SubGetList(ctx context.Context, req *subUser.SubGetListReq) (res *subUser.SubGetListRes, err error)
	SubUpdate(ctx context.Context, req *subUser.SubUpdateReq) (res *subUser.SubUpdateRes, err error)
	SubChangeStatus(ctx context.Context, req *subUser.SubChangeStatusReq) (res *subUser.SubChangeStatusRes, err error)
	SubDelete(ctx context.Context, req *subUser.SubDeleteReq) (res *subUser.SubDeleteRes, err error)
}

type IV1System interface {
	CacheRemove(ctx context.Context, req *system.CacheRemoveReq) (res *system.CacheRemoveRes, err error)
	PersonalInfo(ctx context.Context, req *system.PersonalInfoReq) (res *system.PersonalInfoRes, err error)
	PersonalEdit(ctx context.Context, req *system.PersonalEditReq) (res *system.PersonalEditRes, err error)
	PersonalResetPwd(ctx context.Context, req *system.PersonalResetPwdReq) (res *system.PersonalResetPwdRes, err error)
	RuleSearch(ctx context.Context, req *system.RuleSearchReq) (res *system.RuleSearchRes, err error)
	RuleAdd(ctx context.Context, req *system.RuleAddReq) (res *system.RuleAddRes, err error)
	RuleGetParams(ctx context.Context, req *system.RuleGetParamsReq) (res *system.RuleGetParamsRes, err error)
	RuleInfo(ctx context.Context, req *system.RuleInfoReq) (res *system.RuleInfoRes, err error)
	RuleUpdate(ctx context.Context, req *system.RuleUpdateReq) (res *system.RuleUpdateRes, err error)
	RuleDelete(ctx context.Context, req *system.RuleDeleteReq) (res *system.RuleDeleteRes, err error)
	ConfigSearch(ctx context.Context, req *system.ConfigSearchReq) (res *system.ConfigSearchRes, err error)
	ConfigAdd(ctx context.Context, req *system.ConfigAddReq) (res *system.ConfigAddRes, err error)
	ConfigGet(ctx context.Context, req *system.ConfigGetReq) (res *system.ConfigGetRes, err error)
	ConfigEdit(ctx context.Context, req *system.ConfigEditReq) (res *system.ConfigEditRes, err error)
	ConfigDelete(ctx context.Context, req *system.ConfigDeleteReq) (res *system.ConfigDeleteRes, err error)
	ConfigGetOneByKey(ctx context.Context, req *system.ConfigGetOneByKeyReq) (res *system.ConfigGetOneByKeyRes, err error)
	DeptSearch(ctx context.Context, req *system.DeptSearchReq) (res *system.DeptSearchRes, err error)
	DeptAdd(ctx context.Context, req *system.DeptAddReq) (res *system.DeptAddRes, err error)
	DeptEdit(ctx context.Context, req *system.DeptEditReq) (res *system.DeptEditRes, err error)
	DeptDelete(ctx context.Context, req *system.DeptDeleteReq) (res *system.DeptDeleteRes, err error)
	DeptTreeSelect(ctx context.Context, req *system.DeptTreeSelectReq) (res *system.DeptTreeSelectRes, err error)
	GetDict(ctx context.Context, req *system.GetDictReq) (res *system.GetDictRes, err error)
	DictDataSearch(ctx context.Context, req *system.DictDataSearchReq) (res *system.DictDataSearchRes, err error)
	DictDataAdd(ctx context.Context, req *system.DictDataAddReq) (res *system.DictDataAddRes, err error)
	DictDataGet(ctx context.Context, req *system.DictDataGetReq) (res *system.DictDataGetRes, err error)
	DictDataEdit(ctx context.Context, req *system.DictDataEditReq) (res *system.DictDataEditRes, err error)
	DictDataDelete(ctx context.Context, req *system.DictDataDeleteReq) (res *system.DictDataDeleteRes, err error)
	DictTypeSearch(ctx context.Context, req *system.DictTypeSearchReq) (res *system.DictTypeSearchRes, err error)
	DictTypeAdd(ctx context.Context, req *system.DictTypeAddReq) (res *system.DictTypeAddRes, err error)
	DictTypeGet(ctx context.Context, req *system.DictTypeGetReq) (res *system.DictTypeGetRes, err error)
	DictTypeEdit(ctx context.Context, req *system.DictTypeEditReq) (res *system.DictTypeEditRes, err error)
	DictTypeDelete(ctx context.Context, req *system.DictTypeDeleteReq) (res *system.DictTypeDeleteRes, err error)
	DictTypeAll(ctx context.Context, req *system.DictTypeAllReq) (res *system.DictTypeAllRes, err error)
	DbInitIsInit(ctx context.Context, req *system.DbInitIsInitReq) (res *system.DbInitIsInitRes, err error)
	DbInitGetEnvInfo(ctx context.Context, req *system.DbInitGetEnvInfoReq) (res *system.DbInitGetEnvInfoRes, err error)
	DbInitCreateDb(ctx context.Context, req *system.DbInitCreateDbReq) (res *system.DbInitCreateDbRes, err error)
	UserLogin(ctx context.Context, req *system.UserLoginReq) (res *system.UserLoginRes, err error)
	UserLoginOut(ctx context.Context, req *system.UserLoginOutReq) (res *system.UserLoginOutRes, err error)
	LoginLogSearch(ctx context.Context, req *system.LoginLogSearchReq) (res *system.LoginLogSearchRes, err error)
	LoginLogDel(ctx context.Context, req *system.LoginLogDelReq) (res *system.LoginLogDelRes, err error)
	LoginLogClear(ctx context.Context, req *system.LoginLogClearReq) (res *system.LoginLogClearRes, err error)
	MonitorSearch(ctx context.Context, req *system.MonitorSearchReq) (res *system.MonitorSearchRes, err error)
	SysOperLogSearch(ctx context.Context, req *system.SysOperLogSearchReq) (res *system.SysOperLogSearchRes, err error)
	SysOperLogGet(ctx context.Context, req *system.SysOperLogGetReq) (res *system.SysOperLogGetRes, err error)
	SysOperLogDelete(ctx context.Context, req *system.SysOperLogDeleteReq) (res *system.SysOperLogDeleteRes, err error)
	SysOperLogClear(ctx context.Context, req *system.SysOperLogClearReq) (res *system.SysOperLogClearRes, err error)
	PostSearch(ctx context.Context, req *system.PostSearchReq) (res *system.PostSearchRes, err error)
	PostAdd(ctx context.Context, req *system.PostAddReq) (res *system.PostAddRes, err error)
	PostEdit(ctx context.Context, req *system.PostEditReq) (res *system.PostEditRes, err error)
	PostDelete(ctx context.Context, req *system.PostDeleteReq) (res *system.PostDeleteRes, err error)
	RoleList(ctx context.Context, req *system.RoleListReq) (res *system.RoleListRes, err error)
	RoleGetParams(ctx context.Context, req *system.RoleGetParamsReq) (res *system.RoleGetParamsRes, err error)
	RoleAdd(ctx context.Context, req *system.RoleAddReq) (res *system.RoleAddRes, err error)
	RoleGet(ctx context.Context, req *system.RoleGetReq) (res *system.RoleGetRes, err error)
	RoleEdit(ctx context.Context, req *system.RoleEditReq) (res *system.RoleEditRes, err error)
	RoleDelete(ctx context.Context, req *system.RoleDeleteReq) (res *system.RoleDeleteRes, err error)
	UserMenus(ctx context.Context, req *system.UserMenusReq) (res *system.UserMenusRes, err error)
	UserSearch(ctx context.Context, req *system.UserSearchReq) (res *system.UserSearchRes, err error)
	UserGetParams(ctx context.Context, req *system.UserGetParamsReq) (res *system.UserGetParamsRes, err error)
	UserAdd(ctx context.Context, req *system.UserAddReq) (res *system.UserAddRes, err error)
	UserEdit(ctx context.Context, req *system.UserEditReq) (res *system.UserEditRes, err error)
	UserGetEdit(ctx context.Context, req *system.UserGetEditReq) (res *system.UserGetEditRes, err error)
	UserResetPwd(ctx context.Context, req *system.UserResetPwdReq) (res *system.UserResetPwdRes, err error)
	UserStatus(ctx context.Context, req *system.UserStatusReq) (res *system.UserStatusRes, err error)
	UserDelete(ctx context.Context, req *system.UserDeleteReq) (res *system.UserDeleteRes, err error)
	UserGetByIds(ctx context.Context, req *system.UserGetByIdsReq) (res *system.UserGetByIdsRes, err error)
	SysUserOnlineSearch(ctx context.Context, req *system.SysUserOnlineSearchReq) (res *system.SysUserOnlineSearchRes, err error)
	SysUserOnlineForceLogout(ctx context.Context, req *system.SysUserOnlineForceLogoutReq) (res *system.SysUserOnlineForceLogoutRes, err error)
}

type IV1User interface {
	Register(ctx context.Context, req *user.RegisterReq) (res *user.RegisterRes, err error)
	GetList(ctx context.Context, req *user.GetListReq) (res *user.GetListRes, err error)
	Update(ctx context.Context, req *user.UpdateReq) (res *user.UpdateRes, err error)
	ChangeStatus(ctx context.Context, req *user.ChangeStatusReq) (res *user.ChangeStatusRes, err error)
	Delete(ctx context.Context, req *user.DeleteReq) (res *user.DeleteRes, err error)
}
