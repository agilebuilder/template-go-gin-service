// Package model 存放数据结构与领域模型。
package model

// HelloResponse 是问候接口的响应体。
type HelloResponse struct {
	Message string `json:"message"`
}
