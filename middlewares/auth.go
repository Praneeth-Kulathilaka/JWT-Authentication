package middlewares

import (
    "github.com/gin-gonic/gin"
    "github.com/dgrijalva/jwt-go"
    "net/http"
    "time"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, "Missing token")
            c.Abort()
            return
        }

        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return []byte("secret-key"), nil
        })

        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, "Invalid token")
            c.Abort()
            return
        }

        c.Next()
    }
}
