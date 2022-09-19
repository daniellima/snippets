package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// getAlbums responds with the list of all albums as JSON.
func getHello(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, map[string]string{
		"one":   "This",
		"two":   "is",
		"three": "a value",
	})
}

func GinFrameworkMain() {
	router := gin.Default()
	router.GET("/hello", getHello)

	router.Run("localhost:8080")
}
