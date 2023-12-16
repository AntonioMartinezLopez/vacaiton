package repository

import (
	"backend/services/tripService/models"
	"errors"
)

type TripRepo interface {
	CreateTrip(createTripInput *models.CreateTripQueryInput, userId string) (res *models.Trip, err error)
	GetTrip(id string, userId string) (res *models.Trip, err error)
	GetTrips(userId string) (res []models.Trip, err error)
	// UpdateTrip(updateTripInput *models.UpdateTripQueryInput) (res *models.Trip, err error)
}

func (r *GormRepository) CreateTrip(createTripInput *models.CreateTripQueryInput, userId string) (trip *models.Trip, err error) {

	// create new trip
	newTrip := models.Trip{
		Query: models.TripQuery{
			Country:         createTripInput.Country,
			Duration:        createTripInput.Duration,
			Secrets:         createTripInput.Secrets,
			MaximumDistance: createTripInput.MaximumDistance,
			Focus:           createTripInput.Focus,
		},
		Stops:  []models.TripStop{},
		UserId: userId,
	}

	result := r.db.Database.Create(&newTrip)

	return &newTrip, result.Error
}

func (r *GormRepository) GetTrip(tripId string, userId string) (*models.Trip, error) {

	// get trip instance
	trip := models.Trip{}
	result := r.db.Database.Preload("Stops").Where("id = ? AND user_id = ?", tripId, userId).First(&trip)

	if result.RowsAffected != 1 {
		return &trip, errors.New("Unknown trip id for given user")
	}

	return &trip, result.Error
}

func (r *GormRepository) GetTrips(userId string) ([]models.Trip, error) {

	// get trip instances for one user
	trip := []models.Trip{}
	result := r.db.Database.Preload("Stops").Where("user_id = ?", userId).Find(&trip)

	return trip, result.Error
}
