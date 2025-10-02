package main

import (
	"github.com/wazecode/CRUDTemplate/initializers"
	"github.com/wazecode/CRUDTemplate/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.Connect2DB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
}
