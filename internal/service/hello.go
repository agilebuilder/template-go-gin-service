// Package service 存放业务逻辑（服务层）。
package service

import "<%= modulePath %>/internal/model"

// HelloService 是示例业务服务。
type HelloService struct{}

// NewHelloService 创建 HelloService。
func NewHelloService() *HelloService {
	return &HelloService{}
}

// Greet 根据名字生成问候语，名字为空时使用 world。
func (s *HelloService) Greet(name string) model.HelloResponse {
	if name == "" {
		name = "world"
	}
	return model.HelloResponse{Message: "hello, " + name}
}
