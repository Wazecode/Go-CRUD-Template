package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wazecode/CRUDTemplate/controllers"
	"github.com/wazecode/CRUDTemplate/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.Connect2DB()
}

func main() {
	r := gin.Default()
	// Define a simple GET endpoint
	r.GET("/createUser", controllers.PostCreate)

	r.Run()
}
