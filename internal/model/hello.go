// Package model holds data structures and domain models.
package model

// HelloResponse is the response body of the greeting endpoint.
type HelloResponse struct {
	Message string `json:"message"`
}
