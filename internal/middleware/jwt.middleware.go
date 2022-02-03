package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/red-life/simple-authentication-with-jwt/pkg/jwt"
	"net/http"
	"strings"
)

func IsAuthorized(jwtService jwt.IJWTPackage) gin.HandlerFunc{
	return func(c *gin.Context) {
		token := strings.Split(c.GetHeader("Authorization")," ")[1]
		if token == ""{
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Missing token"})
			return
		}
		email, err := jwtService.ValidateToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.Set("Email", email)
		c.Next()
	}
}