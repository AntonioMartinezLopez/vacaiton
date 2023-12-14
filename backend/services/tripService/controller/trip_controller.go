package controller

import (
	"backend/pkg/jsonHelper"
	"backend/services/tripService/models"
	"backend/services/tripService/repository"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type TripHandler struct {
	repo      repository.TripRepo
	validator *validator.Validate
}

func NewTripHandler(repo repository.TripRepo, validator *validator.Validate) *TripHandler {
	return &TripHandler{
		repo:      repo,
		validator: validator,
	}
}

func (h *TripHandler) CreateTrip(w http.ResponseWriter, request *http.Request) {

	// Validate Input
	createTripInfo := new(models.CreateTripQueryInput)
	err := jsonHelper.DecodeJSONAndValidate(request.Body, createTripInfo)
	if err != nil {
		jsonHelper.HttpErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	// Create Trip
	trip, err := h.repo.CreateTrip(createTripInfo)
	if err != nil {
		jsonHelper.HttpErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	// Return result
	jsonHelper.HttpResponse(&trip, w)
}
