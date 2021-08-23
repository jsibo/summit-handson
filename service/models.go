package service

import "github.com/go-playground/validator/v10"

type student struct {
	Name    string   `json:"name,omitempty" validate:"required"`
	StudentId    string   `json:"student_id,omitempty" validate:"required"`
	Program  string `json:"program,omitempty" validate:"required"`
	Courses []string `json:"courses,omitempty"`
	Graduated bool     `json:"graduated,omitempty"`
	EnrollmentDate string     `json:"enrollment_date,omitempty" validate:"required"`
}

type StudentDirectory struct {
	Validator *validator.Validate
}

var datastore = defaultData

var defaultData = map[string]student{
	"1": {
		Name:    "Michael Jordan",
		StudentId: "1",
		Program:  "App Development",
		Courses: []string{"Agile 101", "Java 101"},
		Graduated: false,
		EnrollmentDate: "7/21/21",
	},
	"2": {
		Name:    "Tyrone Plunkett",
		StudentId: "2",
		Program:  "App Development",
		Courses: []string{"Agile 101", "Java 101", "Public Speaking", "Capstone"},
		Graduated: true,
		EnrollmentDate: "6/1/21",
	},
}