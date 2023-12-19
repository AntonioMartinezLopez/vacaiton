package controller

import (
	"backend/pkg/jsonHelper"
	"backend/pkg/middlewares"
	"backend/services/tripService/models"
	"backend/services/tripService/repository"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type StopHandler struct {
	repo      repository.StopRepo
	validator *validator.Validate
}

func NewStopHandler(repo repository.StopRepo, validator *validator.Validate) *StopHandler {
	return &StopHandler{
		repo:      repo,
		validator: validator,
	}
}

func (h *StopHandler) CreateStop(w http.ResponseWriter, request *http.Request) {

	// validate input
	createStopInput := new(models.CreateStopInput)
	err := jsonHelper.DecodeJSONAndValidate(request.Body, createStopInput)

	if err != nil {
		jsonHelper.HttpErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	// user data
	userClaims := middlewares.Claims{}
	middlewares.ReadUserClaimsFromContext(w, request, &userClaims)

	// Create stop
	stop, err := h.repo.CreateStop(createStopInput, userClaims.UserId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		jsonHelper.HttpErrorResponse(w, http.StatusNotFound, err)
		return
	}
	if err != nil {
		jsonHelper.HttpErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	// Return result
	jsonHelper.HttpResponse(&stop, w)
}

func (h *StopHandler) CreateStops(w http.ResponseWriter, request *http.Request) {

	// validate input
	createStopsInput := new(models.CreateStopsInput)
	err := jsonHelper.DecodeJSONAndValidate(request.Body, createStopsInput)

	if err != nil {
		jsonHelper.HttpErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	// user data
	userClaims := middlewares.Claims{}
	middlewares.ReadUserClaimsFromContext(w, request, &userClaims)

	// Create stops
	stops, err := h.repo.CreateStops(createStopsInput, userClaims.UserId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		jsonHelper.HttpErrorResponse(w, http.StatusNotFound, err)
		return
	}
	if err != nil {
		jsonHelper.HttpErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	// Return created results
	jsonHelper.HttpResponse(stops, w)

}
