package main

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
)

type Sensor struct {

    Latitude float64 `json:"latitude"`
    Longitude float64 `json:"longitude"`
    Sensor string `json:"sensor"`
    Code int `json:"code"`
    Manufacturer string `json:"manufacturer"`
    Author string `json:"author"`
    Date string `json:"date"`
}

// postsensors adds an sensor from JSON received in the request body.
func postsensor(c *gin.Context) {
    var data Sensor

    if err := c.BindJSON(&data); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.IndentedJSON(http.StatusCreated, InsertIntoSensors(data))
}

func getsensor(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, GetInSensors())
}

type Gas struct {
    SensorId int `json:"sensorId"`
    SensorName string `json:"sensorName"`
    Unit string `json:"unit"`
    Time string `json:"time"`
    NH3 float64 `json:"NH3"`
    CO float64 `json:"CO"`
    C2H5OH float64 `json:"C2H5OH"`
    H2 float64 `json:"H2"`
    IC4H10 float64 `json:"IC4H10"`
    CH4 float64 `json:"CH4"`
    NO2 float64 `json:"NO2"`
    C3H8 float64 `json:"C3H8"`
}

// postgas adds an gas from JSON received in the request body.
func postgas(c *gin.Context) {
    var data Gas

    if err := c.BindJSON(&data); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.IndentedJSON(http.StatusCreated, insertIntoGases(data))
}

func getgas(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, GetInGases())
}

type Radiation struct {
    SensorId int `json:"sensorId"`
    SensorName string `json:"sensorName"`
    Unit string `json:"unit"`
    Time string `json:"time"`
    Radiation float64 `json:"radiation"`
}

// postradiation adds an radiation from JSON received in the request body.
func postradiation(c *gin.Context) {
    var data Radiation

    if err := c.BindJSON(&data); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.IndentedJSON(http.StatusCreated, insertIntoRadiation(data))
}

func getradiation(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, GetInRadiation())
}

func main() {
    CreateTable()
    router := gin.Default()
    router.POST("/sensors", postsensor)
    router.GET("/sensors", getsensor)
    router.POST("/gases", postgas)
    router.GET("/gases", getgas)
    router.POST("/radiations", postradiation)
    router.GET("/radiations", getradiation)

    fmt.Println("Server will run on http://localhost:8000")
    router.Run(":8000")
}