package main

import "github.cm/gin-gonic/gin"

func start(c *gin.Context) {
	c.JSON(200, gin.H{
		"start": "heap",
	})
}

func main() {
	router := gin.Default()

	r.GET("/start", start)

	router.Run()
}
