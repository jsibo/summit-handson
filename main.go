package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jsibo/summit-handson/service"
)

func main() {

	router := gin.Default()

	// Our endpoints here

	router.Run()
	}
}