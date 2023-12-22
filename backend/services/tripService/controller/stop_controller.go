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

//	@Summary		Create a new stop
//	@Tags			Stop
//	@Description	This endpoint can be used to add a stop to an existing trip. Requirements: authenticated
//	@ID				create-stop
//	@Accept			json
//	@Produce		json
//	@Param			CreateStopInput	body		models.CreateStopInput	true	"User Input for creating a new stop"
//	@Success		200				{object}	models.TripStop
//	@Failure		400				{object}	jsonHelper.HTTPError	"In case of invalid CreateStop DTO"
//	@Failure		401				{object}	jsonHelper.HTTPError	"In case of unauthenticated request"
//	@Failure		404				{object}	jsonHelper.HTTPError	"In case of unknown trip id"
//	@Failure		500				{object}	jsonHelper.HTTPError	"In case of persistence error"
//	@Router			/stop [post]
//	@Security		ApiKeyAuth
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

//	@Summary		Create multiple stops
//	@Tags			Stop
//	@Description	This endpoint can be used to multiple stops to an existing trip. Requirements: authenticated
//	@ID				create-stops
//	@Accept			json
//	@Produce		json
//	@Param			CreateStopsInput	body		models.CreateStopsInput	true	"User Input for creating new stops"
//	@Success		200					{array}		models.TripStop
//	@Failure		400					{object}	jsonHelper.HTTPError	"In case of invalid CreateStops DTO"
//	@Failure		401					{object}	jsonHelper.HTTPError	"In case of unauthenticated request"
//	@Failure		404					{object}	jsonHelper.HTTPError	"In case of unknown trip id"
//	@Failure		500					{object}	jsonHelper.HTTPError	"In case of persistence error"
//	@Router			/stops [post]
//	@Security		ApiKeyAuth
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
