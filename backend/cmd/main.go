package main

import (
	"breath-of-the-chef-recipe-guide/backend/internal/api"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(CORSMiddleware())

	r.GET("/ping", api.PingPong)
	r.GET("/recipes", api.GetRecipe)
	r.GET("/ingredients", api.GetIngredients)

	err := r.Run()
	if err != nil {
		panic(err)
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
