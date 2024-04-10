package main

import (
	"fmt"
	"net/http"

    jwt "2024-T0002-EC09-G01/src/jwt"

	"github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
)

func main() {
    router := gin.Default()

    config := cors.DefaultConfig()
    config.AllowOrigins = []string{"*"}
    router.Use(cors.New(config))

    router.POST("/login", jwt.LoginHandler)

	protected := router.Group("")
    protected.Use(jwt.AuthMiddleware())
    {
        protected.GET("/sensors", getsensors)
        protected.POST("/sensors", postsensor)
        protected.GET("/alerts", getalerts)
        protected.POST("/alerts", postalert)
    }

    port := ":8000"
    fmt.Printf("Server will run on http://localhost%s\n", port)
    router.Run(port)
}

// postsensors adds an sensor from JSON received in the request body.
func postsensor(c *gin.Context) {
    var data Sensor

    if err := c.BindJSON(&data); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	client := ConnectToMongo()

    c.IndentedJSON(http.StatusCreated, InsertIntoMongo(client, data, "sensors"))
}

// getsensors responds with the list of all sensors as JSON.
func getsensors(c *gin.Context) {
	client := ConnectToMongo()
    c.IndentedJSON(http.StatusOK, GetonMongo(client))
}

// postalert adds an alert from JSON received in the request body.
func postalert(c *gin.Context) {
    var data Alert

    if err := c.BindJSON(&data); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    client := ConnectToMongo()

    c.IndentedJSON(http.StatusCreated, InsertIntoMongo(client, data, "alerts"))
}

// getalerts responds with the list of all alerts as JSON.
func getalerts(c *gin.Context) {
    client := ConnectToMongo()
    c.IndentedJSON(http.StatusOK, GetonMongo(client))
}