// Package controllers is used for controlleing the database
package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/wazecode/CRUDTemplate/initializers"
	"github.com/wazecode/CRUDTemplate/models"
)

func PostCreate(c *gin.Context) {
	newUser := models.User{Name: "Shuwais", Email: "shuwais2003@gmail.com"}
	result := initializers.DB.Create(&newUser)

	if result.Error != nil {
		c.Status(400)
		log.Fatal("ERROR: Controllers : Failed to add user to the Database")
	}
	// Return JSON response
	c.JSON(200, gin.H{
		"New User": newUser,
	})
}
