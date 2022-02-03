package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/red-life/simple-authentication-with-jwt/pkg/jwt"
	"net/http"
)

func isAuthorized(jwtService jwt.IJWTPackage) gin.HandlerFunc{
	return func(c *gin.Context) {
		token, err := c.Cookie("token")
		if err != nil{
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		_, err = jwtService.CreateToken(token)
		if err != nil {
			//context.JSON(http.StatusBadRequest, gin.H{"error": "Missing token", "status": false})
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.Next()
	}
}