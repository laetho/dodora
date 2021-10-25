package main

import (
	"github.com/alron/ginlogr"
	"github.com/gin-gonic/gin"
	d "github.com/laetho/doas/pkg/deliveries"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sync"
	"time"

)


func init() {
}

func main() {
	wg := sync.WaitGroup{}

	wg.Add(3)

	r := gin.Default()
	r.Use(ginlogr.Ginlogr(zap.New(), time.RFC3339, true), gin.Recovery())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	go func() {
		defer wg.Done()
		r.Run()
	}()

	go func() {
		defer wg.Done()
		d.Run()
	}()

	wg.Wait()
}
