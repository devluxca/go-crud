package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-crud/models"
)

type CreateUserInput struct {
	Name string `json:"name" validate:"required"`
    Age int `json:"age" validate:"required"`
}

type UpdateUserInput struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
}

// @router /api/v1/users [get]
func FindUsers(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// @router /api/v1/users/:id [get]
func FindUser(c *gin.Context) {
	var user models.User

	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// @router /api/v1/users/:id [delete]
func DeleteUser(c *gin.Context) {
    var user models.User

    if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error})
		return
	}

	models.DB.Delete(&user)
	
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// @router /api/v1/user/adults [get]
func FindAdults(c *gin.Context) {
    var users []models.User

	if err := models.DB.Where("age > ?", 18).Find(&users).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
	}

    c.JSON(http.StatusOK, gin.H{"data": users})
}

// @router /api/v1/users [post]
func CreateUser(c *gin.Context) {
    var input CreateUserInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{Name: input.Name, Age: input.Age}
	models.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// @router /api/v1/users/:id [patch]
func UpdateUser(c *gin.Context) {
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
	}

	var input UpdateUserInput
	if err := c.ShouldBindJSON(&input);err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}

	models.DB.Model(&user).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": user})
}