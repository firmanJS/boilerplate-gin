package route

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func InitRoute(route *gin.Engine) {
	/*
	   @description of index route
	*/
	route.GET("/", func(c *gin.Context) {
		name, err := os.Hostname()
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusNotFound, gin.H{"host": name})
	})

	route.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	route.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "resource not found", "data": nil})
	})

}
