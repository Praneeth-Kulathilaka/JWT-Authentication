package main

import (
	"go_gin_JWT_auth/Routes"
	"go_gin_JWT_auth/middlewares"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"JWT/models"
	"JWT/routes"
	"JWT/middlewares"
)
var (
	client *mongo.Client
	jwtSecret = []byte("secret-key")
)
func main() {
	route := gin.Default();

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    client, _ = mongo.Connect(nil, clientOptions)

	route.GET("/ping", func(c *gin.Context) {
        c.String(http.StatusOK, "Pong")
    })

	route.POST("/register", Routes.Register)
	route.POST("/login", Routes.Login)
	route.GET("/protected", middlewares.AuthMiddleware(), Protected)

	route.Run(":8080");

}
func Protected(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "This is a protected route"})
}
