package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"member/controller"
	"member/middlewares"
)

func Router(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	//register & login
	r.POST("api/register", controller.CreateRegister)
	r.POST("api/login", controller.Login)

	//non member
	r.GET("api/produk", controller.GetAllProduct)
	r.GET("api/produk/:id", controller.GetProductById)

	r.POST("/pesan", controller.CreatePesan)
	r.GET("/pesan/:id", controller.GetPesanById)
	r.GET("/transaksi/:id", controller.GetTransaksiById)

	r.GET("api/feedback", controller.GetAllFeedback)

	//admin
	AdminMiddlewareRoute := r.Group("/api")
	AdminMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	AdminMiddlewareRoute.Use(middlewares.IsAdmin())
	AdminMiddlewareRoute.POST("/produk", controller.CreateProduct)
	AdminMiddlewareRoute.PATCH("/produk/:id", controller.UpdateProduct)
	AdminMiddlewareRoute.DELETE("/produk/:id", controller.DeleteProduct)

	AdminMiddlewareRoute.GET("/pesan", controller.GetAllPesan)
	AdminMiddlewareRoute.PATCH("/pesan/:id", controller.UpdatePesan)
	AdminMiddlewareRoute.DELETE("api/pesan/:id", controller.DeletePesan)

	AdminMiddlewareRoute.GET("/register", controller.GetAllRegister)

	AdminMiddlewareRoute.GET("/profil", controller.GetAllProfil)
	AdminMiddlewareRoute.GET("/profil/:id", controller.GetProfilById)

	AdminMiddlewareRoute.PATCH("/saldo/:id", controller.UpdateSaldo)
	AdminMiddlewareRoute.DELETE("/saldo/:id", controller.DeleteSaldo)

	AdminMiddlewareRoute.GET("/transaksi", controller.GetAllTransaksi)
	AdminMiddlewareRoute.PATCH("/transaksi/:id", controller.UpdateTransaksi)
	AdminMiddlewareRoute.DELETE("api/transaksi/:id", controller.DeleteTransaksi)

	//member
	MemberMiddlewareRoute := r.Group("/api")
	MemberMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	MemberMiddlewareRoute.GET("/saldo", controller.GetAllSaldo)
	MemberMiddlewareRoute.POST("/saldo", controller.CreateSaldo)
	MemberMiddlewareRoute.GET("/saldo/:id", controller.GetSaldoById)

	MemberMiddlewareRoute.GET("/register/:id", controller.GetRegisterById)
	MemberMiddlewareRoute.PATCH("/register/:id", controller.UpdateRegister)
	MemberMiddlewareRoute.DELETE("api/register/:id", controller.DeleteRegister)

	MemberMiddlewareRoute.POST("/profil", controller.CreateProfil)
	MemberMiddlewareRoute.PATCH("/profil/:id", controller.UpdateProfil)
	MemberMiddlewareRoute.DELETE("api/profil/:id", controller.DeleteProfil)

	MemberMiddlewareRoute.POST("/transaksi", controller.CreateTransaksi)

	MemberMiddlewareRoute.POST("/feedback", controller.CreateFeedback)
	MemberMiddlewareRoute.GET("/feedback/:id", controller.GetFeedbackById)
	MemberMiddlewareRoute.PATCH("/feedback/:id", controller.UpdateFeedback)
	MemberMiddlewareRoute.DELETE("api/feedback/:id", controller.DeleteFeedback)

	return r
}
