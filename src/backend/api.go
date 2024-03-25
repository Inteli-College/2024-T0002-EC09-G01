package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    router.GET("/sensors", getsensors)
    router.POST("/sensors", postsensor)

    fmt.Println("Server will run on http://localhost:8000")
    router.Run(":8000")
}

type Sensor struct {
    Sensor string `json:"sensor"`
    Tipo string `json:"tipo"`
	Latitude string `json:"latitude"`
	Longitude string `json:"longitude"`
}

// postsensors adds an sensor from JSON received in the request body.
func postsensor(c *gin.Context) {
    var data Sensor

    if err := c.BindJSON(&data); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	c.IndentedJSON(http.StatusCreated, data)

	// client := ConnectToMongo()

    // c.IndentedJSON(http.StatusCreated, InsertIntoMongo(client, data))
}

// getsensors responds with the list of all sensors as JSON.
func getsensors(c *gin.Context) {
	client := ConnectToMongo()
    c.IndentedJSON(http.StatusOK, GetonMongo(client))
}