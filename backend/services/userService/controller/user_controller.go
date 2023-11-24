package controller

import (
	"backend/pkg/jsonHelper"
	"backend/services/userService/models"
	"backend/services/userService/repository"
	"encoding/json"
	"net/http"
)

type UserHandler struct {
	repo repository.UserRepo
}

func NewUserHandler(repo repository.UserRepo) *UserHandler {
	return &UserHandler{
		repo: repo,
	}
}

//   func (h *ExampleHandler) GetData(w http.ResponseWriter, request *http.Request) {
// 	q := request.URL.Query()
// 	limit, _ := strconv.Atoi(q.Get("limit"))
// 	offset, _ := strconv.Atoi(q.Get("offset"))

// 	data, err := h.repo.GetExamples(int64(limit), int64(offset))
// 	if err != nil {
// 	  http.Error(w, err.Error(), http.StatusBadRequest)
// 	  return
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(&data)
//   }

func (h *UserHandler) CreateUser(w http.ResponseWriter, request *http.Request) {
	userInput := new(models.RegisterUserInput)
	err := json.NewDecoder(request.Body).Decode(&userInput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.repo.RegisterUser(userInput)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jsonHelper.HttpResponse(struct {
		UserId int
	}{UserId: id}, w)
}
