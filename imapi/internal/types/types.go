// Code generated by goctl. DO NOT EDIT.
package types

type SendMsgRequest struct {
	ToUserId int64  `json:"toUserId"`
	Content  string `json:"content"`
}

type SendMsgResponse struct {
	Id         int64  `json:"id"`
	CreateTime int64  `json:"createTime"`
	Content    string `json:"content"`
}