package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// the yow struct is a collection of one-line phrases
type yow struct {
	ID     string `json:"id"`
	Line   string `json:"line"`
}

// let's populate some yow lines
var yowLines = []yow {
	{ID: "1", Line: "One of the crossbeams is out of skew on the treddle"},
	{ID: "2", Line: "Nobody expects the Spanish Inquisition"},
	{ID: "3", Line: "Our chief weapon is surprise"},
}

func getYowLines(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, yowLines)
}

func main() {
	router := gin.Default()
	router.GET("/", getYowLines)

	router.Run("localhost:8080")
}