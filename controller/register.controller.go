package controller

import (
	"member/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type registerInput struct {
	Nama     string `json:"nama"`
	Email    string `json:"Email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func GetAllRegister(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var getreg []models.Registers
	db.Find(&getreg)

	c.JSON(http.StatusOK, gin.H{"data": getreg})
}

func CreateRegister(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var regin registerInput
	if err := c.ShouldBindJSON(&regin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data := models.Registers{Nama: regin.Nama, Email: regin.Email, Password: regin.Password, Role: regin.Role}
	db.Create(&data)

	c.JSON(http.StatusOK, gin.H{"data": data})
}

func GetRegisterById(c *gin.Context) { // Get model if exist
	var getrego []models.Registers

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", c.Param("id")).Find(&getrego).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": getrego})
}

func UpdateRegister(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var upreg models.Registers
	if err := db.Where("id = ?", c.Param("id")).First(&upreg).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate regin
	var upregin registerInput
	if err := c.ShouldBindJSON(&upregin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Registers
	updatedInput.Nama = upregin.Nama
	updatedInput.Email = upregin.Email
	updatedInput.Password = upregin.Password
	updatedInput.Role = upregin.Role

	db.Model(&upreg).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": upreg})
}

func DeleteRegister(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var delreg models.Registers
	if err := db.Where("id = ?", c.Param("id")).First(&delreg).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&delreg)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
