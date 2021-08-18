package service

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

// Create / Post
func (s *StudentDirectory) AddStudent(c *gin.Context) {

	// Code for function

	// Error handling
	}
}

// Read / - Get all Students
func (s *StudentDirectory) GetAllStudents(c *gin.Context) {
	// Code for function

	// Error handling
}

// Read / - Get a Student by ID 
func (s *StudentDirectory) GetStudentById(c *gin.Context) {
	if val, ok := datastore[c.Param("student_id")]; ok {

		c.JSON(http.StatusOK, datastore[c.Param("student_id")])
		return
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": fmt.Sprintf("Unable to find a student with id: %v.", c.Param("student_id")),
	})
}

// Update / Put - Update a Student by ID
func (s *StudentDirectory) UpdateStudentById(c *gin.Context) {
	// Code for function

	// Error handling
}

// Update / Put - Add Course to Student by ID
func (s *StudentDirectory) AddCourseToStudentById(c *gin.Context) {
	// Code for function

	// Error handling
}

// Update / Put - Update Graduation Status by ID
func (s *StudentDirectory) UpdateGraduatedById(c *gin.Context) {
	// Code for function

	// Error handling
}

// Delete / Delete - Delete a Student By ID
func (s *StudentDirectory) DeleteStudentById(c *gin.Context) {
	// Code for function

	// Error handling
}

// Reset data 
func (s *StudentDirectory) ResetData(c *gin.Context) {
	datastore = map[string]student{
		"Michael Jordan": {
			Name:    "Michael Jordan",
			StudentId: "mj",
			Program:  "App Development",
			Courses: []string{"Agile 101", "Java 101"},
			Graduated: false,
			EnrollmentDate: "7/21/21"
		},
		"Tyrone Plunkett": {
			Name:    "Tyrone Plunkett",
			StudentId: "tp",
			Program:  "App Development",
			Courses: []string{"Agile 101", "Java 101", "Public Speaking", "Capstone"},
			Graduated: true,
			EnrollmentDate: "6/1/21"
		},
	}
	c.JSON(http.StatusAccepted, gin.H{
		"message": "Data reset.",
	})
}