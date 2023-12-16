package controller

import (
	"backend/pkg/jsonHelper"
	"backend/pkg/middlewares"
	"backend/services/tripService/models"
	"backend/services/tripService/repository"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
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

	// user data
	userClaims, assertionCorrect := request.Context().Value("user-claims").(middlewares.Claims)
	if !assertionCorrect {
		jsonHelper.HttpErrorResponse(w, http.StatusInternalServerError, errors.New("Error in user claim type assertion"))
		return
	}

	// Create Trip
	trip, err := h.repo.CreateTrip(createTripInfo, userClaims.UserId)
	if err != nil {
		jsonHelper.HttpErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	// Return result
	jsonHelper.HttpResponse(&trip, w)
}

func (h *TripHandler) GetTrip(w http.ResponseWriter, request *http.Request) {

	// Route param
	tripID := chi.URLParam(request, "id")

	// user data
	userClaims, assertionCorrect := request.Context().Value("user-claims").(middlewares.Claims)
	if !assertionCorrect {
		jsonHelper.HttpErrorResponse(w, http.StatusInternalServerError, errors.New("Error in user claim type assertion"))
		return
	}

	trip, err := h.repo.GetTrip(tripID, userClaims.UserId)

	if err != nil {
		jsonHelper.HttpErrorResponse(w, http.StatusNotFound, err)
		return
	}

	jsonHelper.HttpResponse(trip, w)
}

func (h *TripHandler) GetTrips(w http.ResponseWriter, request *http.Request) {

	// user data
	userClaims, assertionCorrect := request.Context().Value("user-claims").(middlewares.Claims)
	if !assertionCorrect {
		jsonHelper.HttpErrorResponse(w, http.StatusInternalServerError, errors.New("Error in user claim type assertion"))
		return
	}

	trips, err := h.repo.GetTrips(userClaims.UserId)

	if err != nil {
		jsonHelper.HttpErrorResponse(w, http.StatusNotFound, err)
		return
	}

	jsonHelper.HttpResponse(trips, w)
}
