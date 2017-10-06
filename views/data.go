package views

import (
	"errors"
	"strconv"
)

// Athlete is a simple struct holding athlete data
type Athlete struct {
	ID   int
	Name string
}

// NewFromID retrieves the athlete from the main slice
func NewFromID(id int) (Athlete, error) {
	for _, a := range athletes {
		if a.ID == id {
			return a, nil
		}
	}
	return Athlete{}, errors.New("Not found")
}

// NewFromString retrieves the athlete from the main slice
func NewFromString(id string) (Athlete, error) {
	i, err := strconv.Atoi(id)
	if err != nil {
		return Athlete{}, err
	}
	return NewFromID(i)
}

var athletes = []Athlete{
	{14583, "Robin ‘TGC Boy’ Leboeuf"},
	{14580, "Kevin ‘Priestt’ Prettre"},
	{14582, "Fabrice ‘Papa de Robin’ Leboeuf"},
}
