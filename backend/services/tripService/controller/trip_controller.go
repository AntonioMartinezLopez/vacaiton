package controller

import (
	"backend/pkg/jsonHelper"
	"backend/pkg/middlewares"
	"backend/services/tripService/models"
	"backend/services/tripService/repository"
	"net/http"
	"strconv"

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

//	@Summary		Create a new trip
//	@Tags			Trip
//	@Description	This endpoint can be used to create a trip. Requirements: authenticated
//	@ID				create-trip
//	@Accept			json
//	@Produce		json
//	@Param			createTripInput	body		models.CreateTripQueryInput	true	"User Input creating a new trip"
//	@Success		200				{object}	models.Trip
//	@Failure		400				{object}	jsonHelper.HTTPError	"In case of invalid createTrip DTO"
//	@Failure		401				{object}	jsonHelper.HTTPError	"In case of unauthenticated request"
//	@Failure		500				{object}	jsonHelper.HTTPError	"In case of persistence error"
//	@Router			/trip [post]
//	@Security		ApiKeyAuth
func (h *TripHandler) CreateTrip(w http.ResponseWriter, request *http.Request) {

	// Validate Input
	createTripInfo := new(models.CreateTripQueryInput)
	err := jsonHelper.DecodeJSONAndValidate(request.Body, createTripInfo)
	if err != nil {
		jsonHelper.HttpErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	// user data
	userClaims := middlewares.Claims{}
	middlewares.ReadUserClaimsFromContext(w, request, &userClaims)

	// Create Trip
	trip, err := h.repo.CreateTrip(createTripInfo, userClaims.UserId)
	if err != nil {
		jsonHelper.HttpErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	// Return result
	jsonHelper.HttpResponse(&trip, w)
}

//	@Summary		Get trip
//	@Tags			Trip
//	@Description	This endpoint is used to query a trip. Requirements: authenticated and requested trip is assigned to user
//	@ID				get-trip
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Trip ID"
//	@Success		200	{object}	models.Trip
//	@Failure		400	{object}	jsonHelper.HTTPError	"In case of missing path param"
//	@Failure		401	{object}	jsonHelper.HTTPError	"In case of unauthenticated request"
//	@Failure		404	{object}	jsonHelper.HTTPError	"In case non existing trip for given user"
//	@Failure		500	{object}	jsonHelper.HTTPError	"In case of	persistence error"
//	@Router			/trip/{id} [get]
//	@Security		ApiKeyAuth
func (h *TripHandler) GetTrip(w http.ResponseWriter, request *http.Request) {

	// Route param
	tripID := chi.URLParam(request, "id")

	// user data
	userClaims := middlewares.Claims{}
	middlewares.ReadUserClaimsFromContext(w, request, &userClaims)

	trip, err := h.repo.GetTrip(tripID, userClaims.UserId)

	if err != nil {
		jsonHelper.HttpErrorResponse(w, http.StatusNotFound, err)
		return
	}

	jsonHelper.HttpResponse(trip, w)
}

//	@Summary		Get all trips
//	@Tags			Trip
//	@Description	This endpoint is used to query all trips of a given user. Requirements: authenticated and requested trip is assigned to user
//	@ID				get-trips
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		models.Trip
//	@Failure		401	{object}	jsonHelper.HTTPError	"In	case of	unauthenticated	request"
//	@Failure		404	{object}	jsonHelper.HTTPError	"In	case non existing trip for given user"
//	@Failure		500	{object}	jsonHelper.HTTPError	"In	case of	persistence error"
//	@Router			/trips [get]
//	@Security		ApiKeyAuth
func (h *TripHandler) GetTrips(w http.ResponseWriter, request *http.Request) {

	// user data
	userClaims := middlewares.Claims{}
	middlewares.ReadUserClaimsFromContext(w, request, &userClaims)

	trips, err := h.repo.GetTrips(userClaims.UserId)

	if err != nil {
		jsonHelper.HttpErrorResponse(w, http.StatusNotFound, err)
		return
	}

	jsonHelper.HttpResponse(trips, w)
}

//	@Summary		Update a trip
//	@Tags			Trip
//	@Description	This endpoint can be used update a trip. This action deletes existing stops and initiates a new calculation. Requirements: authenticated
//	@ID				update-trip
//	@Accept			json
//	@Produce		json
//	@Param			updateTripInput	body		models.UpdateTripQueryInput	true	"User Input for updating a trip"
//	@Param			id				path		int							true	"Trip ID"
//	@Success		200				{object}	models.Trip
//	@Failure		400				{object}	jsonHelper.HTTPError	"In	case of invalid updateTrip DTO or missing path param"
//	@Failure		401				{object}	jsonHelper.HTTPError	"In	case of unauthenticated	request"
//	@Failure		500				{object}	jsonHelper.HTTPError	"In	case of persistence	error"
//	@Router			/trip/{id} [put]
//	@Security		ApiKeyAuth
func (h *TripHandler) UpdateTrip(w http.ResponseWriter, request *http.Request) {

	// Route param
	tripID := chi.URLParam(request, "id")

	// Validate Input
	updateTripInput := new(models.UpdateTripQueryInput)
	err := jsonHelper.DecodeJSONAndValidate(request.Body, updateTripInput)
	if err != nil {
		jsonHelper.HttpErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	// Overwrite if id is set in param
	if tripID != "" {
		id, _ := strconv.Atoi(tripID)
		updateTripInput.Id = uint(id)
	}

	// user data
	userClaims := middlewares.Claims{}
	middlewares.ReadUserClaimsFromContext(w, request, &userClaims)

	trip, err := h.repo.UpdateTrip(updateTripInput, userClaims.UserId)

	if err != nil {
		jsonHelper.HttpErrorResponse(w, http.StatusNotFound, err)
		return
	}

	jsonHelper.HttpResponse(trip, w)
}

//	@Summary		Delete a trip
//	@Tags			Trip
//	@Description	This endpoint can be used delete a trip including its stops. Requirements: authenticated
//	@ID				delete-trip
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Trip ID"
//	@Success		200	{object}	models.Trip
//	@Failure		400	{object}	jsonHelper.HTTPError	"In case of missing path param"
//	@Failure		401	{object}	jsonHelper.HTTPError	"In case of unauthenticated	request"
//	@Failure		500	{object}	jsonHelper.HTTPError	"In case of persistence	error"
//	@Router			/trip/{id} [delete]
//	@Security		ApiKeyAuth
func (h *TripHandler) DeleteTrip(w http.ResponseWriter, request *http.Request) {

	// Route param
	tripID := chi.URLParam(request, "id")

	// user data
	userClaims := middlewares.Claims{}
	middlewares.ReadUserClaimsFromContext(w, request, &userClaims)

	trip, err := h.repo.DeleteTrip(tripID, userClaims.UserId)

	if err != nil {
		jsonHelper.HttpErrorResponse(w, http.StatusNotFound, err)
		return
	}

	jsonHelper.HttpResponse(trip, w)
}
