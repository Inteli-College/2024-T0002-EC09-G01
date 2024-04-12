package jwt

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	mongo "2024-T0002-EC09-G01/src/internal/mongo"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var secretKey = []byte(os.Getenv("JWT_SECRET"))

func PublishData(c *gin.Context) {
	c.String(http.StatusOK, "Implement Data Publish Here!")
}

func CreateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func LoginHandler(c *gin.Context) {
	var u User
	if err := c.BindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := mongo.ConnectToMongo("./../config/.env")

	coll := client.Database("testdb").Collection("users")

	query := bson.D{primitive.E{Key: "username", Value: u.Username}, primitive.E{Key: "password", Value: u.Password}}

	var result User

	err := coll.FindOne(context.TODO(), query).Decode(&result)

	if err == nil {
		tokenString, err := CreateToken(result.Username)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": tokenString})
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
}

func AuthMiddleware() gin.HandlerFunc {
	
	return func(c *gin.Context) {
		// tokenString := c.GetHeader("Authorization")
		// if tokenString == "" {
		// 	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization header"})
		// 	return
		// }

		// tokenString = tokenString[len("Bearer "):]

		// token, err := VerifyToken(tokenString)
		// if err != nil || !token.Valid {
		// 	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		// 	return
		// }

		c.Next()
	}
}