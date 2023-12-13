package models

import "time"

type StopHighlight struct {
	Id          uint       `json:"id" example:"1"`
	CreatedAt   *time.Time `json:"created_at,string,omitempty" example:"2023-12-01T12:37:59.008583Z"`
	UpdatedAt   *time.Time `json:"updated_at_at,string,omitempty" example:"2023-12-01T12:37:59.008583Z"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Coordinates string     `json:"coordinates"`
	TripStopID  uint
}
