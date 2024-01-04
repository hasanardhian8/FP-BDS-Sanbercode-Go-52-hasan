package controller

import (
	"member/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type profilInput struct {
	IdRegister int `json:"idRegister"`
	IdSaldo    int `json:"idSaldo"`
}

func GetAllProfil(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var getPro []models.Profils
	db.Find(&getPro)

	c.JSON(http.StatusOK, gin.H{"data": getPro})
}

func CreateProfil(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var inPro profilInput
	if err := c.ShouldBindJSON(&inPro); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var inReg models.Registers
	if err := db.Where("id = ?", inPro.IdRegister).First(&inReg).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id register not found!"})
		return
	}
	var inSal models.Saldos
	if err := db.Where("id = ?", inPro.IdRegister).First(&inSal).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id saldo not found!"})
		return
	}
	data := models.Profils{IdRegister: inPro.IdRegister, IdSaldo: inPro.IdSaldo}
	db.Create(&data)

	c.JSON(http.StatusOK, gin.H{"data": data})
}

func GetProfilById(c *gin.Context) { // Get model if exist
	var getProId []models.Profils

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", c.Param("id")).Find(&getProId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": getProId})
}

func UpdateProfil(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var upPro models.Profils
	if err := db.Where("id = ?", c.Param("id")).First(&upPro).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate inPro
	var upinPro profilInput
	if err := c.ShouldBindJSON(&upinPro); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var upReg models.Registers
	if err := db.Where("id = ?", upinPro.IdRegister).First(&upReg).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	var upSal models.Saldos
	if err := db.Where("id = ?", upinPro.IdRegister).First(&upSal).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var updatedInput models.Profils
	updatedInput.IdRegister = upinPro.IdRegister
	updatedInput.IdSaldo = upinPro.IdSaldo

	db.Model(&upPro).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": upPro})
}

func DeleteProfil(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var delPro models.Profils
	if err := db.Where("id = ?", c.Param("id")).First(&delPro).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&delPro)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
