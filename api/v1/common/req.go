/*
* @desc:公共接口相关
* @company:云南奇讯科技有限公司
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2022/3/30 9:28
 */

package common

import "sms_backend/internal/model"

// PageReq 公共请求参数
type PageReq struct {
	model.PageReq
	Condition map[string]interface{} `json:"condition" description:"查询条件"`
}

type Author struct {
	Authorization string `p:"Authorization" in:"header" dc:"Bearer {{token}}"`
}
