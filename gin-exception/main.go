package main

import (
	"net/http"

	"gin.exception/exception"
	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()
	e.Use(exception.Handler)

	e.GET("/index", func(c *gin.Context) {
		c.JSON(200, gin.H{"testK": "testV"})
	})
	e.GET("/e1", func(c *gin.Context) {
		panic(&exception.Api{Code: http.StatusInternalServerError, Message: "err.Error()"})
	})
	e.GET("/e2", func(c *gin.Context) {
		panic("error")
	})

	e.Run(":8000")
}
