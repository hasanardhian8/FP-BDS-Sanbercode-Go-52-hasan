package middlewares

import (
	"fmt"
	"net/http"

	"member/models"
	"member/utils/token"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}
		c.Next()
	}
}
func IsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := c.MustGet("db").(*gorm.DB)

		userID, err := token.ExtractTokenID(c)
		fmt.Println("user:", userID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized", "message": "Invalid or missing token"})
			c.Abort()
			return
		}

		var user models.Registers
		if err := db.First(&user, userID).Error; err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden", "message": "User not found or invalid role"})
			c.Abort()
			return
		}
		fmt.Println("user:", user)
		isAdmin := user.Role == "admin"

		if isAdmin {
			c.Next()
		} else {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden", "message": "You are not an admin"})
			c.Abort()
		}
	}

}
