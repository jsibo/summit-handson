package service

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

// Create / Post
func (s *MovieService) AddThing(c *gin.Context) {

	// Code for function

	// Error handling
	}
}

// Read / Get all movies
func (s *MovieService) GetThings(c *gin.Context) {
	// Code for function

	// Error handling
}

// Read / Get an individual movie
func (s *MovieService) GetOneThing(c *gin.Context) {

	// Code for function
}

// Update / Put
func (s *MovieService) UpdateThing(c *gin.Context) {
	// Code for function

	// Error handling
}

// Delete / Delete
func (s *MovieService) DeleteThing(c *gin.Context) {
	// Code for function

	// Error handling
}

// Reset datastore 
// TODO: Update based upon service we create
func (s *MovieService) ResetData(c *gin.Context) {
	datastore = map[string]movies{
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
	c.JSON(http.StatusAccepted, gin.H{
		"message": "Data reset.",
	})
}