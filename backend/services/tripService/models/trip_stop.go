package models

import (
	"time"
)

type TripStop struct {
	Id          uint            `json:"id" example:"1"`
	CreatedAt   *time.Time      `json:"created_at,string,omitempty" example:"2023-12-01T12:37:59.008583Z"`
	UpdatedAt   *time.Time      `json:"updated_at_at,string,omitempty" example:"2023-12-01T12:37:59.008583Z"`
	Name        string          `json:"stopName"`
	Coordinates string          `json:"coordinates"`
	Days        int             `json:"days"`
	Highlights  []StopHighlight `json:"highlights" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TripID      uint            `json:"trip_id"`
}

type TripStopInput struct {
	Name        string               `json:"name"`
	Coordinates string               `json:"coordinates"`
	Days        int                  `json:"days"`
	Highlights  []StopHighlightInput `json:"highlights" validate:"required,dive,required"`
}

type CreateStopInput struct {
	Stop   TripStopInput `json:"stop" validate:"required"`
	TripId uint          `json:"trip_id"`
}

type CreateStopsInput struct {
	Stops  []TripStopInput `json:"stop" validate:"required,dive,required"`
	TripId uint            `json:"trip_id"`
}
