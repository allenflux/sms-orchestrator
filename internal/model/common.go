/*
* @desc:公用model
* @company:云南奇讯科技有限公司
* @Author: yixiaohu<yxh669@qq.com>
* @Date:   2023/5/11 22:43
 */

package model

// PageReq 公共请求参数
type PageReq struct {
	DateRange []string `json:"date_range" p:"dateRange" description:"日期范围"`         //日期范围
	PageNum   int      `json:"page_num" p:"pageNum" description:"当前页码" default:"1"` //当前页码
	PageSize  int      `json:"page_size" p:"pageSize" description:"每页数量"`           //每页数
	OrderBy   string   `json:"order_by"`                                            //排序方式
}

// ListRes 列表公共返回
type ListRes struct {
	CurrentPage int         `json:"currentPage" description:"当前页"`
	Total       interface{} `json:"total" description:"总数"`
}

func (p *PageReq) PageReqInit() {
	//if p.PageNum == 0 {
	//	p.PageNum = 1
	//}
	//if p.PageSize == 0 {
	//	p.PageSize = 10
	//}
}

func (p *PageReq) GetStartAndEnd(length int) (int, int) {
	start := (p.PageNum - 1) * p.PageSize
	end := start + p.PageSize
	if start > length {
		return 0, 0
	}
	if end > length {
		return start, length
	}
	return start, end
}
