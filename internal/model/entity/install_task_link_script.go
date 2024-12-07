// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// InstallTaskLinkScript is the golang structure for table install_task_link_script.
type InstallTaskLinkScript struct {
	Id            uint `json:"id"              orm:"id"              description:""`       //
	InstallTaskId int  `json:"install_task_id" orm:"install_task_id" description:"安装任务id"` // 安装任务id
	ScriptId      int  `json:"script_id"       orm:"script_id"       description:"脚本id"`   // 脚本id
}
