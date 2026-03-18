package main

import (
	"example.com/enterprise_grade_go_project/pkg/config"
	"example.com/enterprise_grade_go_project/pkg/controller"
	"example.com/enterprise_grade_go_project/pkg/service"
	"fmt"
)

func main() {
	config.InitConfig()
	controller := controller.NewController(service.NewService())
	controller.Start()
	fmt.Println("Server started successfully")
}