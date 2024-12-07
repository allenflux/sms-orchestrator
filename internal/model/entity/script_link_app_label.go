// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// ScriptLinkAppLabel is the golang structure for table script_link_app_label.
type ScriptLinkAppLabel struct {
	Id         uint `json:"id"           orm:"id"           description:""`        //
	ScriptId   int  `json:"script_id"    orm:"script_id"    description:"脚本id"`    // 脚本id
	AppLabelId int  `json:"app_label_id" orm:"app_label_id" description:"app标签id"` // app标签id
}
