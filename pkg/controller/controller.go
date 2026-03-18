package controller

import (
	"example.com/enterprise_grade_go_project/pkg/service"
	"net/http"
)

// Controller represents the application controller.
type Controller struct {
	Service *service.Service
}

// NewController returns a new instance of the Controller.
func NewController(s *service.Service) *Controller {
	return &Controller{Service: s}
}

// Start starts the server.
func (c *Controller) Start() {
	r := http.NewServeMux()
	r.HandleFunc("/users", c.GetUsers)
	http.ListenAndServe(":8080", r)
}