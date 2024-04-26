package main

import (
	"breath-of-the-chef-recipe-guide/backend/internal/api"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", api.PingPong)
	r.Run()
}
