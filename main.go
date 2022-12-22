package main

import (
	"net/http"
	"math/rand"
	"github.com/gin-gonic/gin"
)

// the yow struct is a collection of one-line phrases
type yow struct {
	ID     int `json:"id"`
	Line   string `json:"line"`
}

// let's populate some yow lines
var yowLines = []yow {
	{ID: 1, Line: "One of the crossbeams is out of skew on the treddle"},
	{ID: 2, Line: "Nobody expects the Spanish Inquisition"},
	{ID: 3, Line: "Our chief weapon is surprise"},
}

func getYowLines(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, yowLines)
}

func getYowLine(c *gin.Context) {
	var rand = rand.Intn(len(yowLines));

	c.IndentedJSON(http.StatusOK, yowLines[rand];
}

func main() {
	router := gin.Default()
	router.GET("/all", getYowLines)
	router.GET("/", getYowLine)
	router.Run("localhost:8080")
}
