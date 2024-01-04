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
	r.POST("/api/register", controller.CreateRegister)
	r.POST("/login", controller.Login)

	//non member
	r.GET("/api/produk", controller.GetAllProduct)
	r.GET("/api/produk/:id", controller.GetProductById)
	r.GET("/api/pesan", controller.GetAllPesan)
	r.GET("/api/feedback", controller.GetAllFeedback)

	//admin
	AdminMiddlewareRoute := r.Group("/")
	AdminMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	AdminMiddlewareRoute.Use(middlewares.IsAdmin())
	AdminMiddlewareRoute.POST("/api/produk/", controller.CreateProduct)
	AdminMiddlewareRoute.PATCH("/api/produk/:id", controller.UpdateProduct)
	AdminMiddlewareRoute.DELETE("/api/produk/:id", controller.DeleteProduct)

	AdminMiddlewareRoute.GET("/api/register", controller.GetAllRegister)
	AdminMiddlewareRoute.GET("/api/register/:id", controller.GetRegisterById)

	AdminMiddlewareRoute.GET("/api/profil", controller.GetAllProfil)
	AdminMiddlewareRoute.GET("/api/profil/:id", controller.GetProfilById)

	//member
	MemberMiddlewareRoute := r.Group("/")
	MemberMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	MemberMiddlewareRoute.GET("/", controller.GetAllSaldo)
	MemberMiddlewareRoute.POST("/", controller.CreateSaldo)
	MemberMiddlewareRoute.GET("/:id", controller.GetSaldoById)
	MemberMiddlewareRoute.PATCH("/:id", controller.UpdateSaldo)
	MemberMiddlewareRoute.DELETE("/:id", controller.DeleteSaldo)

	MemberMiddlewareRoute.PATCH("/api/register/:id", controller.UpdateRegister)
	MemberMiddlewareRoute.DELETE("api/register/:id", controller.DeleteRegister)

	MemberMiddlewareRoute.POST("/api/profil", controller.CreateProfil)
	MemberMiddlewareRoute.PATCH("/api/profil/:id", controller.UpdateProfil)
	MemberMiddlewareRoute.DELETE("api/profil/:id", controller.DeleteProfil)

	//member
	MemberMiddlewareRoute.GET("/api/transaksi", controller.GetAllTransaksi)
	MemberMiddlewareRoute.POST("/api/transaksi", controller.CreateTransaksi)
	MemberMiddlewareRoute.GET("/api/transaksi/:id", controller.GetTransaksiById)
	MemberMiddlewareRoute.PATCH("/api/transaksi/:id", controller.UpdateTransaksi)
	MemberMiddlewareRoute.DELETE("api/transaksi/:id", controller.DeleteTransaksi)

	MemberMiddlewareRoute.POST("/api/pesan", controller.CreatePesan)
	//r.GET("/api/pesan/:id", controller.GetPesanById)
	MemberMiddlewareRoute.PATCH("/api/pesan/:id", controller.UpdatePesan)
	MemberMiddlewareRoute.DELETE("api/pesan/:id", controller.DeletePesan)

	MemberMiddlewareRoute.POST("/api/feedback", controller.CreateFeedback)
	//r.GET("/api/feedback/:id", controller.GetFeedbackById)
	MemberMiddlewareRoute.PATCH("/api/feedback/:id", controller.UpdateFeedback)
	MemberMiddlewareRoute.DELETE("api/feedback/:id", controller.DeleteFeedback)

	return r
}
