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
	// GET /recipes?category=<category_name>

	recipeQueryName := c.Query("name")
	categoryQueryName := c.Query("category")

	if len(recipeQueryName) > 0 {
		getRecipeByRecipeName(c, recipeQueryName)
	}

	if len(categoryQueryName) > 0 {
		getRecipesByCategory(c, categoryQueryName)
	}

}

func getRecipeByRecipeName(c *gin.Context, s string) {
	recipeName := strings.ReplaceAll(s, "-", " ")
	recipeIngredientList := repo.GetRecipeByName(recipeName)

	var ingredients []string
	var cat string
	var recipe Recipe
	caser := cases.Title(language.English)

	if len(recipeIngredientList) == 0 {
		e := ErrorResponse{Code: "recipe_not_found", Title: "Recipe Not Found", Detail: fmt.Sprintf("Unable to find recipe with name '%s'", s)}
		errors := []ErrorResponse{e}
		c.JSON(http.StatusNotFound, JsonResponse{Data: nil, Errors: errors})
	}

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
}

func getRecipesByCategory(c *gin.Context, s string) {
	cat := strings.ReplaceAll(s, "-", " ")
	recipeIngredientList := repo.GetRecipesByCategory(cat)

	if len(recipeIngredientList) == 0 {
		e := ErrorResponse{Code: "category_not_found", Title: "Category Not Found", Detail: fmt.Sprintf("Unable to find recipes under category '%s'", cat)}
		errors := []ErrorResponse{e}
		c.JSON(http.StatusNotFound, JsonResponse{Data: nil, Errors: errors})
	}

	recipesMap := make(map[string]*Recipe)

	for _, row := range recipeIngredientList {
		name := row.RecipeName
		ingredient := row.IngredientName

		recipe, ok := recipesMap[name]
		if !ok {
			recipe = &Recipe{
				Name:        name,
				Category:    cat,
				Ingredients: []string{ingredient},
			}
			recipesMap[name] = recipe
		}
		recipe.Ingredients = append(recipe.Ingredients, ingredient)
	}

	var recipes []Recipe
	for _, recipe := range recipesMap {
		recipes = append(recipes, *recipe)
	}

	c.JSON(http.StatusOK, JsonResponse{Data: recipes})
}
