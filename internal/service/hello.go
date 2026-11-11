// Package service holds business logic (service layer).
package service

import "<%= modulePath %>/internal/model"

// HelloService is the example business service.
type HelloService struct{}

// NewHelloService creates a HelloService.
func NewHelloService() *HelloService {
	return &HelloService{}
}

// Greet builds a greeting for the given name, defaulting to world when empty.
func (s *HelloService) Greet(name string) model.HelloResponse {
	if name == "" {
		name = "world"
	}
	return model.HelloResponse{Message: "hello, " + name}
}
