package repository

import "backend/services/tripService/models"

type TripRepo interface {
	CreateTrip(createTripInput *models.CreateTripQueryInput) (res *models.Trip, err error)
	// GetTrip(id string) (res *models.Trip, err error)
	// UpdateTrip(updateTripInput *models.UpdateTripQueryInput) (res *models.Trip, err error)
}

func (r *GormRepository) CreateTrip(createTripInput *models.CreateTripQueryInput) (trip *models.Trip, err error) {

	// create new trip
	newTrip := models.Trip{
		Query: models.TripQuery{
			Country:         createTripInput.Country,
			Duration:        createTripInput.Duration,
			Secrets:         createTripInput.Secrets,
			MaximumDistance: createTripInput.MaximumDistance,
			Focus:           createTripInput.Focus,
		},
		Stops: []models.TripStop{},
	}

	result := r.db.Database.Create(&newTrip)

	return &newTrip, result.Error

}
