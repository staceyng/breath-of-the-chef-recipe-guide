package api

import (
	"net/http"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"breath-of-the-chef-recipe-guide/backend/internal/repo"

	"github.com/gin-gonic/gin"
)

func PingPong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func GetRecipe(c *gin.Context) {
	// GET /recipes?name=<recipe_name>
	queryName := c.Query("name")
	recipeName := strings.ReplaceAll(queryName, "-", " ")

	recipeIngredientList := repo.GetRecipeByName(recipeName)

	var ingredients []string
	var cat string
	caser := cases.Title(language.English)
	for _, v := range recipeIngredientList {
		ingredients = append(ingredients, caser.String(v.IngredientName))
		cat = v.Category
	}

	recipe := Recipe{
		Name:        caser.String(recipeName),
		Category:    caser.String(cat),
		Ingredients: ingredients,
	}
	c.JSON(http.StatusOK, recipe)
}
