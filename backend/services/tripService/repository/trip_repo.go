package repository

import (
	"backend/pkg/logger"
	"backend/services/tripService/models"
	"errors"

	"gorm.io/gorm"
)

type TripRepo interface {
	CreateTrip(createTripInput *models.CreateTripQueryInput, userId string) (res *models.Trip, err error)
	GetTrip(id string, userId string) (res *models.Trip, err error)
	GetTrips(userId string) (res []models.Trip, err error)
	UpdateTrip(updateTripInput *models.UpdateTripQueryInput, userId string) (res *models.Trip, err error)
	DeleteTrip(id string, userId string) (err error)
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
	result := r.db.Database.Preload("Stops").Preload("Stops.Highlights").Preload("Query").Where("id = ? AND user_id = ?", tripId, userId).First(&trip)

	if result.RowsAffected != 1 {
		return &trip, errors.New("Unknown trip id for given user")
	}

	logger.Log("Country" + trip.Query.Country)

	return &trip, result.Error
}

func (r *GormRepository) GetTrips(userId string) ([]models.Trip, error) {

	// get trip instances for one user
	trip := []models.Trip{}
	result := r.db.Database.Preload("Stops").Preload("Stops.Highlights").Preload("Query").Where("user_id = ?", userId).Find(&trip)

	return trip, result.Error
}

func (r *GormRepository) UpdateTrip(updateTripInput *models.UpdateTripQueryInput, userId string) (*models.Trip, error) {

	// get trip instance
	trip := models.Trip{}
	var transactionError error

	r.db.Database.Transaction(func(tx *gorm.DB) error {
		result := r.db.Database.Preload("Stops").Preload("Query").Where("id = ? AND user_id = ?", updateTripInput.Id, userId).First(&trip)

		if result.RowsAffected != 1 {
			transactionError = errors.New("Unknown trip id for given user")
			return transactionError
		}

		// remove old query
		if err := r.db.Database.Where("trip_id = ?", trip.Id).Delete(&models.TripQuery{}).Error; err != nil {
			transactionError = err
			return err
		}

		// reset stops - they are being calculated new
		if err := r.db.Database.Where("trip_id = ?", trip.Id).Delete(&models.TripStop{}).Error; err != nil {
			transactionError = err
			return err
		}

		// overwrite query object and reset all stops
		trip.Query = models.TripQuery{
			Country:         updateTripInput.Country,
			Duration:        updateTripInput.Duration,
			Secrets:         updateTripInput.Secrets,
			MaximumDistance: updateTripInput.MaximumDistance,
			Focus:           updateTripInput.Focus,
		}
		trip.Stops = []models.TripStop{}

		if err := r.db.Database.Save(&trip).Error; err != nil {
			transactionError = err
			return err
		}

		return nil
	})

	return &trip, transactionError
}

func (r *GormRepository) DeleteTrip(tripId string, userId string) error {

	result := r.db.Database.Where("user_id = ? AND id = ?", userId, tripId).Delete(&models.Trip{})

	if result.RowsAffected == 0 {
		return errors.New("User does not have a trip with id: " + tripId)
	}

	return result.Error
}
