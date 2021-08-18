package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jsibo/summit-handson/service"
	"github.com/go-playground/validator/v10"
)

func main() {
	svc := service.StudentDirectory{
		Validator: validator.New(),
	}

	router := gin.Default()

	router.GET("/student/:id", svc.GetStudentById)
	router.Run()
}