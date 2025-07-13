package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

// TokenAuthMiddleware checks for a valid token in the Authorization header
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if len(authHeader) < 7 || authHeader[:7] != "Bearer " {
			c.AbortWithStatusJSON(401, gin.H{"error": "Authorization header missing or invalid"})
			return
		}
		token := authHeader[7:]
		if !isValidToken(token) {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid token"})
			return
		}
		c.Next()
	}
}

// isValidToken validates a JWT token using a secret key
func isValidToken(tokenString string) bool {
	const secretKey = "sykell" // Replace with your actual secret key

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		return false
	}
	return true
}

type CrawlerService struct {
	DbConnection *gorm.DB `json:"dbConnection"`
}

func (v *CrawlerService) InitRouter(routerEngine *gin.Engine) {
	maintenanceGeneral := routerEngine.Group("/")
	// maintenanceGeneral.Use(TokenAuthMiddleware()) // Apply middleware to all routes in this group
	maintenanceGeneral.POST("/login", v.Login)
	maintenanceGeneral.POST("/refresh-token", v.RefreshToken)
}
