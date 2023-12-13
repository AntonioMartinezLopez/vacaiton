package models

import (
	"database/sql/driver"
	"encoding/json"
)

type TripFocus int

const (
	Cities TripFocus = iota
	Nature
	Mixed
)

// Stirng handling and marshalling
func (t TripFocus) String() string {
	return [...]string{"Cities", "Nature", "Mixed"}[t]
}

func (t *TripFocus) FromString(focus string) TripFocus {
	return map[string]TripFocus{
		"Cities": Cities,
		"Nature": Nature,
		"Mixed":  Mixed,
	}[focus]
}

func (t TripFocus) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *TripFocus) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*t = t.FromString(s)
	return nil
}

// database handling
func (st *TripFocus) Scan(value interface{}) error {
	b, ok := value.(int)
	if !ok {
		*st = TripFocus(b)
	}
	return nil
}

func (st TripFocus) Value() (driver.Value, error) {
	return int(st), nil
}
