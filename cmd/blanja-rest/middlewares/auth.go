package middlewares

import (
	"strings"
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/maulidihsan/interop-commerce/config"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		config := config.GetConfig()
		reqJWT := c.Request.Header.Get("Authorization");
		if len(strings.TrimSpace(reqJWT)) == 0 {
			c.AbortWithStatus(401)
			return
        }
		reqJWT = strings.TrimPrefix(reqJWT, "Bearer ")

		token, err := jwt.Parse(reqJWT, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(config.GetString("jwt.secret")), nil
		})
		if err != nil {
			c.AbortWithStatus(401)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("user", claims["data"])
			c.Next()
		} else {
			c.AbortWithStatus(401)
			return
		}
	}
}
