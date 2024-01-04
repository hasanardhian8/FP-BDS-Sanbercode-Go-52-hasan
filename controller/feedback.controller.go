package controller

import (
	"member/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type feedInput struct {
	IdTransaksi int       `json:"idTransaksi"`
	IdProfil    int       `json:"idProfil"`
	Komen       string    `json:"komen"`
	Rating      int       `json:"rating"`
	Tanggal     time.Time `json:"tanggal"`
}

func GetAllFeedback(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var getFee []models.Feedbacks
	db.Find(&getFee)

	c.JSON(http.StatusOK, gin.H{"data": getFee})
}

func CreateFeedback(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var inFee feedInput
	if err := c.ShouldBindJSON(&inFee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var inPro models.Profils
	if err := db.Where("id = ?", inFee.IdProfil).First(&inPro).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id profil not found!"})
		return
	}
	var inTrans models.Transaksis
	if err := db.Where("id = ?", inFee.IdTransaksi).First(&inTrans).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id transaksi not found!"})
		return
	}
	data := models.Feedbacks{IdProfil: inFee.IdProfil, IdTransaksi: inFee.IdTransaksi, Komen: inFee.Komen, Rating: inFee.Rating, Tanggal: time.Now().UTC()}
	db.Create(&data)

	c.JSON(http.StatusOK, gin.H{"data": data})
}

func GetFeedbackById(c *gin.Context) { // Get model if exist
	var getFeeId []models.Feedbacks

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", c.Param("id")).Find(&getFeeId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": getFeeId})
}

func UpdateFeedback(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var upFee models.Feedbacks
	if err := db.Where("id = ?", c.Param("id")).First(&upFee).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var upInFee feedInput
	if err := c.ShouldBindJSON(&upInFee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var upPro models.Profils
	if err := db.Where("id = ?", upInFee.IdProfil).First(&upPro).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	var upTrans models.Transaksis
	if err := db.Where("id = ?", upInFee.IdProfil).First(&upTrans).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var updatedInput models.Feedbacks
	updatedInput.IdProfil = upInFee.IdProfil
	updatedInput.IdTransaksi = upInFee.IdTransaksi
	updatedInput.Tanggal = time.Now().UTC()
	updatedInput.Komen = upInFee.Komen
	updatedInput.Rating = upInFee.Rating

	db.Model(&upFee).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": upFee})
}

func DeleteFeedback(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var delFee models.Feedbacks
	if err := db.Where("id = ?", c.Param("id")).First(&delFee).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&delFee)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
