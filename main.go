package main

import(
	"net/http"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "github.com/dgrijalva/jwt-go"
	"./models/user.go"
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

	route.POST("/register", Register)
	route.POST("/login", Login)
	route.GET("/protected", AuthMiddleware(), Protected)

	route.Run(":8080");

}