package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Recipe struct {
	Name        string   `json:"name"`
	Ingredients []string `json:"ingredients"`
	Category    string   `json:"category"`
}

type RecipesJSONData struct {
	Recipes []Recipe `json:"recipes"`
}

func GetAllRecipes() []Recipe {
	recipes, err := readFromJson()
	if err != nil {
		fmt.Println("Error reading recipe data:", err)
		return []Recipe{}
	}

	fmt.Println("Number of recipes:", len(recipes))

	return recipes
}

func readFromJson() ([]Recipe, error) {
	jsonFile := "data/recipes.json"
	file, err := os.Open(jsonFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data RecipesJSONData
	byteValue, _ := ioutil.ReadAll(file)
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		return nil, err
	}

	return data.Recipes, nil
}
