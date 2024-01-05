package controller

import (
	"member/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type saldoInput struct {
	Pembayaran string    `json:"pembayaran"`
	Nominal    int       `json:"nominal"`
	Total      int       `json:"total"`
	Tanggal    time.Time `json:"tanggal"`
}

func GetAllSaldo(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var getsal []models.Saldos
	db.Find(&getsal)

	c.JSON(http.StatusOK, gin.H{"data": getsal})
}

func CreateSaldo(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var salin saldoInput
	if err := c.ShouldBindJSON(&salin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingData models.Saldos
	if err := db.Where(models.Saldos{Pembayaran: salin.Pembayaran, Tanggal: salin.Tanggal}).First(&existingData).Error; err != nil {
		data := models.Saldos{Pembayaran: salin.Pembayaran, Nominal: salin.Nominal, Total: salin.Nominal, Tanggal: time.Now().UTC()}
		db.Create(&data)
		c.JSON(http.StatusOK, gin.H{"data": data})
	} else {
		existingData.Nominal = salin.Nominal
		existingData.Total += salin.Nominal
		db.Save(&existingData)
		c.JSON(http.StatusOK, gin.H{"data": existingData})
	}
	// data := models.Saldos{Pembayaran: salin.Pembayaran, Nominal: salin.Nominal, Total: salin.Total, Tanggal: time.Now().UTC()}
	// db.Create(&data)

	// c.JSON(http.StatusOK, gin.H{"data": data})
}

func GetSaldoById(c *gin.Context) { // Get model if exist
	var getsalo []models.Saldos

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", c.Param("id")).Find(&getsalo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": getsalo})
}

func UpdateSaldo(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var upsal models.Saldos
	if err := db.Where("id = ?", c.Param("id")).First(&upsal).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate salin
	var upsalin saldoInput
	if err := c.ShouldBindJSON(&upsalin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	upsal.Pembayaran = upsalin.Pembayaran
	upsal.Nominal = upsalin.Nominal
	upsal.Total += upsalin.Nominal
	upsal.Tanggal = time.Now().UTC()

	db.Save(&upsal)

	c.JSON(http.StatusOK, gin.H{"data": upsal})
}

func DeleteSaldo(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var delsal models.Saldos
	if err := db.Where("id = ?", c.Param("id")).First(&delsal).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&delsal)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
