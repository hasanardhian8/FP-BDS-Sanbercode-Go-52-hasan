package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"member/controller"
)

func Router(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.GET("/api/produk", controller.GetAllProduct)
	r.POST("/api/produk", controller.CreateProduct)
	r.GET("/api/produk/:id", controller.GetProductById)
	r.PATCH("/api/produk/:id", controller.UpdateProduct)
	r.DELETE("api/produk/:id", controller.DeleteProduct)

	r.GET("/api/saldo", controller.GetAllSaldo)
	r.POST("/api/saldo", controller.CreateSaldo)
	r.GET("/api/saldo/:id", controller.GetSaldoById)
	r.PATCH("/api/saldo/:id", controller.UpdateSaldo)
	r.DELETE("api/saldo/:id", controller.DeleteSaldo)

	r.GET("/api/register", controller.GetAllRegister)
	r.POST("/api/register", controller.CreateRegister)
	r.GET("/api/register/:id", controller.GetRegisterById)
	r.PATCH("/api/register/:id", controller.UpdateRegister)
	r.DELETE("api/register/:id", controller.DeleteRegister)

	r.GET("/api/pesan", controller.GetAllPesan)
	r.POST("/api/pesan", controller.CreatePesan)
	r.GET("/api/pesan/:id", controller.GetPesanById)
	r.PATCH("/api/pesan/:id", controller.UpdatePesan)
	r.DELETE("api/pesan/:id", controller.DeletePesan)

	r.GET("/api/profil", controller.GetAllProfil)
	r.POST("/api/profil", controller.CreateProfil)
	r.GET("/api/profil/:id", controller.GetProfilById)
	r.PATCH("/api/profil/:id", controller.UpdateProfil)
	r.DELETE("api/profil/:id", controller.DeleteProfil)

	r.GET("/api/transaksi", controller.GetAllTransaksi)
	r.POST("/api/transaksi", controller.CreateTransaksi)
	r.GET("/api/transaksi/:id", controller.GetTransaksiById)
	r.PATCH("/api/transaksi/:id", controller.UpdateTransaksi)
	r.DELETE("api/transaksi/:id", controller.DeleteTransaksi)

	r.GET("/api/feedback", controller.GetAllFeedback)
	r.POST("/api/feedback", controller.CreateFeedback)
	r.GET("/api/feedback/:id", controller.GetFeedbackById)
	r.PATCH("/api/feedback/:id", controller.UpdateFeedback)
	r.DELETE("api/feedback/:id", controller.DeleteFeedback)
	return r
}
