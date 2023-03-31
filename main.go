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
	{ID: 4, Line: "Fear and surprise are our two chief weapons"},
	{ID: 5, Line: "And ruthless efficiency"},
	{ID: 6, Line: "Amongst our weaponry are such diverse elements as"},
	{ID: 7, Line: "Fear, surprise, ruthless efficiency, and an almost fanatical devotion to the pope."},
	{ID: 8, Line: "Cardinal!"},
}

func getYowLines(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, yowLines)
}

func getYowLine(c *gin.Context) {
	var rand = rand.Intn(len(yowLines))
	c.IndentedJSON(http.StatusOK, yowLines[rand])
}

func main() {
	router := gin.Default()
	router.GET("/all", getYowLines)
	router.GET("/", getYowLine)
	router.Run("0.0.0.0:8080")
}
