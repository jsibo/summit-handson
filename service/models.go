package service

import "github.com/go-playground/validator/v10"

type thing struct {
	
	// code here 
}

type MovieService struct {
	Validator *validator.Validate
}

// TODO: Update based upon struct/what we want to create

var datastore = defaultData

var defaultData = map[string]movies{
	"Space Jam": {
		Name:    "Space Jam",
		Ratings: 5,
		Actors:  []string{"Michael Jordan", "Bugs Bunny"},
		Watched: true,
	},
	"Scarface": {
		Name:    "Scarface",
		Ratings: 5,
		Actors:  []string{"Al Pacino"},
		Watched: true,
	},
}