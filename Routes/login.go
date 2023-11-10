package Routes

import (
	"JWT/models"
	"context"
	"go/token"
	"go_gin_JWT_auth/models"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context){
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}
	collection := (*mongo.Client).Database("testJWT").Collection("users")
	result := collection.FindOne(context.Background(), bson.M{"username":user.Username})
	storedUser := models.User{}
	err := result.Decode(&storedUser)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error":"Invalid Credentials"})
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password),[]byte(user.Password))
	if err != nil{
		c.JSON(http.StatusUnauthorized, gin.H{"error":"Invalid Password"})
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"username":storedUser.Username
		"exp": time.Now().Add(time.Hour*24).Unix(),
	})
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"Error generated token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token":tokenString})
}