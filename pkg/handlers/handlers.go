package handlers

import (
	"example.com/enterprise_grade_go_project/pkg/controller"
	"example.com/enterprise_grade_go_project/pkg/model"
	"example.com/enterprise_grade_go_project/pkg/utils"
	"net/http"
)

// GetUsersHandler handles the GET /users request.
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	controller := controller.NewController(service.NewService())
	users, err := controller.Service.UserRepository.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json, err := utils.ToJSON(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, json)
}