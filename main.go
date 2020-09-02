package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func statusOKHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func versionHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"version": "v1.0"})
}

func main() {
	router := gin.New()
	router.Use(gin.Recovery())
	router.GET("/ping", statusOKHandler)
	router.GET("/version", versionHandler)
	router.Run(":8080")
}
