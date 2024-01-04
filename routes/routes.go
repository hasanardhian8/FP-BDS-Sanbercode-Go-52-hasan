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
	r.POST("/register", controller.CreateRegister)
	r.POST("/login", controller.Login)

	//non member
	r.GET("/produk", controller.GetAllProduct)
	r.GET("/produk/:id", controller.GetProductById)
	r.GET("/pesan", controller.GetAllPesan)
	r.GET("/feedback", controller.GetAllFeedback)

	//admin
	AdminMiddlewareRoute := r.Group("/")
	AdminMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	AdminMiddlewareRoute.Use(middlewares.IsAdmin())
	AdminMiddlewareRoute.POST("/produk", controller.CreateProduct)
	AdminMiddlewareRoute.PATCH("/produk/:id", controller.UpdateProduct)
	AdminMiddlewareRoute.DELETE("/produk/:id", controller.DeleteProduct)

	AdminMiddlewareRoute.GET("/register", controller.GetAllRegister)
	AdminMiddlewareRoute.GET("/register/:id", controller.GetRegisterById)

	AdminMiddlewareRoute.GET("/profil", controller.GetAllProfil)
	AdminMiddlewareRoute.GET("/profil/:id", controller.GetProfilById)

	//member
	MemberMiddlewareRoute := r.Group("/")
	MemberMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	MemberMiddlewareRoute.GET("/saldo", controller.GetAllSaldo)
	MemberMiddlewareRoute.POST("/saldo", controller.CreateSaldo)
	MemberMiddlewareRoute.GET("/saldo/:id", controller.GetSaldoById)
	MemberMiddlewareRoute.PATCH("/saldo/:id", controller.UpdateSaldo)
	MemberMiddlewareRoute.DELETE("/saldo/:id", controller.DeleteSaldo)

	MemberMiddlewareRoute.PATCH("/register/:id", controller.UpdateRegister)
	MemberMiddlewareRoute.DELETE("api/register/:id", controller.DeleteRegister)

	MemberMiddlewareRoute.POST("/profil", controller.CreateProfil)
	MemberMiddlewareRoute.PATCH("/profil/:id", controller.UpdateProfil)
	MemberMiddlewareRoute.DELETE("api/profil/:id", controller.DeleteProfil)

	MemberMiddlewareRoute.GET("/transaksi", controller.GetAllTransaksi)
	MemberMiddlewareRoute.POST("/transaksi", controller.CreateTransaksi)
	MemberMiddlewareRoute.GET("/transaksi/:id", controller.GetTransaksiById)
	MemberMiddlewareRoute.PATCH("/transaksi/:id", controller.UpdateTransaksi)
	MemberMiddlewareRoute.DELETE("api/transaksi/:id", controller.DeleteTransaksi)

	MemberMiddlewareRoute.POST("/pesan", controller.CreatePesan)
	//r.GET("/pesan/:id", controller.GetPesanById)
	MemberMiddlewareRoute.PATCH("/pesan/:id", controller.UpdatePesan)
	MemberMiddlewareRoute.DELETE("api/pesan/:id", controller.DeletePesan)

	MemberMiddlewareRoute.POST("/feedback", controller.CreateFeedback)
	//r.GET("/feedback/:id", controller.GetFeedbackById)
	MemberMiddlewareRoute.PATCH("/feedback/:id", controller.UpdateFeedback)
	MemberMiddlewareRoute.DELETE("api/feedback/:id", controller.DeleteFeedback)

	return r
}
