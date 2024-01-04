package controller

import (
	"member/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type produkInput struct {
	Nama  string `json:"nama"`
	Harga int    `json:"harga"`
}

func GetAllProduct(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var produk []models.Produks
	db.Find(&produk)

	c.JSON(http.StatusOK, gin.H{"data": produk})
}

func CreateProduct(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var input produkInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data := models.Produks{Nama: input.Nama, Harga: input.Harga}
	db.Create(&data)

	c.JSON(http.StatusOK, gin.H{"data": data})
}

func GetProductById(c *gin.Context) { // Get model if exist
	var produk []models.Produks

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", c.Param("id")).Find(&produk).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": produk})
}

func UpdateProduct(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var produks models.Produks
	if err := db.Where("id = ?", c.Param("id")).First(&produks).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input produkInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Produks
	updatedInput.Nama = input.Nama
	updatedInput.Harga = input.Harga

	db.Model(&produks).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": produks})
}

func DeleteProduct(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var pro models.Produks
	if err := db.Where("id = ?", c.Param("id")).First(&pro).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&pro)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
