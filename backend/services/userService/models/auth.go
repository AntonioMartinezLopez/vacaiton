package models

import (
	"encoding/json"
	"fmt"
)

type LoginStatus int

const (
	LoggedIn  LoginStatus = 1
	LoggedOut LoginStatus = 2
)

func (l LoginStatus) String() string {
	switch l {
	case LoggedIn:
		return "Logged in."
	case LoggedOut:
		return "Logged out."
	default:
		return fmt.Sprintf("%d", int(l))
	}
}

func (l LoginStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal((l.String()))
}

type AuthResponse struct {
	Status LoginStatus `json:"status" binding:"required"`
}
