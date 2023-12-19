package repository

import (
	"backend/services/tripService/models"

	"gorm.io/gorm"
)

type StopRepo interface {
	CreateStop(createStopInput *models.CreateStopInput, userId string) (res *models.TripStop, err error)
	CreateStops(createStopsInput *models.CreateStopsInput, userId string) (stops []models.TripStop, err error)
}

func (r *GormRepository) CreateStop(createStopInput *models.CreateStopInput, userId string) (*models.TripStop, error) {

	// check whether user has the right to create a stop for given trip
	checkResult := r.db.Database.Where("id = ? AND user_id = ?", createStopInput.TripId, userId).First(&models.Trip{})

	if checkResult.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	// build all highlight structs
	highlights := make([]models.StopHighlight, len(createStopInput.Stop.Highlights))
	for i, highlightInput := range createStopInput.Stop.Highlights {
		newHighlight := models.StopHighlight{
			Name:        highlightInput.Name,
			Description: highlightInput.Description,
			Coordinates: highlightInput.Coordinates,
		}
		highlights[i] = newHighlight
	}

	// create new stop
	newStop := models.TripStop{
		Name:        createStopInput.Stop.Name,
		Coordinates: createStopInput.Stop.Coordinates,
		Days:        createStopInput.Stop.Days,
		Highlights:  highlights,
		TripID:      createStopInput.TripId,
	}

	result := r.db.Database.Create(&newStop)

	return &newStop, result.Error
}

func (r *GormRepository) CreateStops(createStopsInput *models.CreateStopsInput, userId string) (stops []models.TripStop, err error) {

	// check whether user has the right to create a stop for given trip
	checkResult := r.db.Database.Where("id = ? AND user_id = ?", createStopsInput.TripId, userId).First(&models.Trip{})

	if checkResult.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	// create all objects and their highlights
	newStops := make([]models.TripStop, len(createStopsInput.Stops))
	var transactionError error

	// start transaction
	r.db.Database.Transaction(func(tx *gorm.DB) error {

		// delete all old existing stops of the given trip
		if err := r.db.Database.Where("trip_id = ?", createStopsInput.TripId).Delete(&models.TripStop{}).Error; err != nil {
			transactionError = err
			return err
		}

		for i, stopInput := range createStopsInput.Stops {

			// Step 1: Create all Highlight objects from a stop
			highlights := make([]models.StopHighlight, len(stopInput.Highlights))
			for j, highlightInput := range stopInput.Highlights {
				newHighlight := models.StopHighlight{
					Name:        highlightInput.Name,
					Description: highlightInput.Description,
					Coordinates: highlightInput.Coordinates,
				}
				highlights[j] = newHighlight
			}

			// Step 2: Create all stops
			newStop := models.TripStop{
				Name:        stopInput.Name,
				Coordinates: stopInput.Coordinates,
				Days:        stopInput.Days,
				Highlights:  highlights,
				TripID:      createStopsInput.TripId,
			}

			newStops[i] = newStop
		}

		if err := r.db.Database.Create(newStops).Error; err != nil {
			transactionError = err
			return err
		}

		return nil

	})

	return newStops, transactionError
}
