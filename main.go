package main

import (
	//"net/http"

	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "berhasil akses home",
		})
	})

	r.Run()
}
