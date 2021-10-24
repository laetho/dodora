package main

import (
	"github.com/gin-gonic/gin"
	d "github.com/laetho/doas/pkg/deliveries"
	"sync"
)

var (
	wg sync.WaitGroup
)

func main() {

	wg.Add(2)
	defer wg.Done()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	go r.Run() // listen and serve on 0.0.0.0:8080
	go d.Run() // Run controller

	wg.Wait()
}
