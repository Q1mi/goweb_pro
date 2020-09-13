package controller

import "bluebell/models"

// 专门用来放接口文档用到的model
// 因为我们的接口文档返回的数据格式是一致的，但是具体的data类型不一致

// _ResponsePostList 帖子列表接口响应数据
type _ResponsePostList struct {
	Code    ResCode                 `json:"code"`    // 业务响应状态码
	Message string                  `json:"message"` // 提示信息
	Data    []*models.ApiPostDetail `json:"data"`    // 数据
}
