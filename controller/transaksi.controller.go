package controller

import (
	"member/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type transInput struct {
	IdProfil   int       `json:"idProfil"`
	IdPesan    int       `json:"idPesan"`
	Tanggal    time.Time `json:"tanggal"`
	Pembayaran string    `json:"pembayaran"`
	Status     bool      `json:"status"`
}

func GetAllTransaksi(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var getTra []models.Transaksis
	db.Find(&getTra)

	c.JSON(http.StatusOK, gin.H{"data": getTra})
}

func CreateTransaksi(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var inTrans transInput
	if err := c.ShouldBindJSON(&inTrans); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var inPro models.Profils
	if err := db.Where("id = ?", inTrans.IdProfil).First(&inPro).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id profil not found!"})
		return
	}
	var inPes models.Pemesanans
	if err := db.Where("id = ?", inTrans.IdPesan).First(&inPes).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id pemesanan not found!"})
		return
	}
	data := models.Transaksis{IdProfil: inTrans.IdProfil, IdPesan: inTrans.IdPesan, Tanggal: time.Now().UTC(), Pembayaran: inTrans.Pembayaran}
	db.Create(&data)

	c.JSON(http.StatusOK, gin.H{"data": data})
}

func GetTransaksiById(c *gin.Context) { // Get model if exist
	var getTraId []models.Transaksis

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", c.Param("id")).Find(&getTraId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": getTraId})
}

func UpdateTransaksi(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var upTrans models.Transaksis
	if err := db.Where("id = ?", c.Param("id")).First(&upTrans).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var upInTrans transInput
	if err := c.ShouldBindJSON(&upInTrans); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var upPro models.Profils
	if err := db.Where("id = ?", upInTrans.IdProfil).First(&upPro).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	var upPes models.Pemesanans
	if err := db.Where("id = ?", upInTrans.IdProfil).First(&upPes).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var updatedInput models.Transaksis
	updatedInput.IdProfil = upInTrans.IdProfil
	updatedInput.IdPesan = upInTrans.IdPesan
	updatedInput.Tanggal = time.Now().UTC()
	updatedInput.Pembayaran = upInTrans.Pembayaran
	updatedInput.Status = upInTrans.Status

	db.Model(&upTrans).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": upTrans})
}

func DeleteTransaksi(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var delTrans models.Transaksis
	if err := db.Where("id = ?", c.Param("id")).First(&delTrans).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&delTrans)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
