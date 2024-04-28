package api

import (
	"fmt"
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
	var recipe Recipe
	caser := cases.Title(language.English)

	if len(recipeIngredientList) > 0 {
		for _, v := range recipeIngredientList {
			ingredients = append(ingredients, caser.String(v.IngredientName))
			cat = v.Category
		}
		recipe = Recipe{
			Name:        caser.String(recipeName),
			Category:    caser.String(cat),
			Ingredients: ingredients,
		}
		c.JSON(http.StatusOK, JsonResponse{Data: recipe})
	} else {
		e := ErrorResponse{Code: "recipe_not_found", Title: "Recipe Not Found", Detail: fmt.Sprintf("Unable to find recipe with name '%s'", recipeName)}
		errors := []ErrorResponse{e}
		c.JSON(http.StatusNotFound, JsonResponse{Data: nil, Errors: errors})
	}
}
