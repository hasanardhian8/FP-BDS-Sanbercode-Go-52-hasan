package controller

import (
	"member/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type pesanInput struct {
	IdProduk     int `json:"idProduk"`
	JumlahBarang int `json:"jumlahBarang"`
	Total        int `json:"total"`
}

func GetAllPesan(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var getpes []models.Pemesanans
	db.Find(&getpes)

	c.JSON(http.StatusOK, gin.H{"data": getpes})
}

func CreatePesan(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var pesin pesanInput
	if err := c.ShouldBindJSON(&pesin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var inpro models.Produks
	if err := db.Where("id = ?", pesin.IdProduk).First(&inpro).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id product not found!"})
		return
	}

	data := models.Pemesanans{
		IdProduk:     pesin.IdProduk,
		JumlahBarang: pesin.JumlahBarang,
		Total:        pesin.JumlahBarang * inpro.Harga,
	}
	db.Create(&data)

	db.Preload("Produks").First(&data, data.Id)

	c.JSON(http.StatusOK, gin.H{"data": data})
}

func GetPesanById(c *gin.Context) { // Get model if exist
	var getpeso []models.Pemesanans

	db := c.MustGet("db").(*gorm.DB)
	//pengin menampilkan berdasarkan id login

	if err := db.Where("id = ?", c.Param("id")).Find(&getpeso).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": getpeso})
}

func UpdatePesan(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var uppes models.Pemesanans
	if err := db.Where("id = ?", c.Param("id")).First(&uppes).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate pesin
	var uppesin pesanInput
	if err := c.ShouldBindJSON(&uppesin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var uppro models.Produks
	if err := db.Where("id = ?", uppesin.IdProduk).First(&uppro).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var updatedInput models.Pemesanans
	updatedInput.IdProduk = uppesin.IdProduk
	updatedInput.JumlahBarang = uppesin.JumlahBarang
	updatedInput.Total = uppesin.JumlahBarang * uppro.Harga

	db.Model(&uppes).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": uppes})
}

func DeletePesan(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var delpes models.Pemesanans
	if err := db.Where("id = ?", c.Param("id")).First(&delpes).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&delpes)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
