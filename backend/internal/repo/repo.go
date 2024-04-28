package repo

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	user   = "stacey.ng"
	host   = "localhost"
	dbname = "botc"
	port   = 5432
)

type RecipeIngredient struct {
	RecipeName     string `db:"recipe_name"`
	Category       string `db:"category"`
	IngredientName string `db:"ingredient_name"`
}

func GetRecipeByName(name string) []RecipeIngredient {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password='' dbname=%s sslmode=disable", host, port, user, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// q := "SELECT * FROM categories"
	q := `
	SELECT
		r.name AS recipe_name,
		r.category AS category,
		i.name AS ingredient_name
	FROM
		recipes r
	INNER JOIN
		recipe_ingredients ri ON r.recipe_id = ri.recipe_id
	INNER JOIN
		ingredients i ON ri.ingredient_id = i.ingredient_id
	WHERE
		r.name = $1
	`

	rows, err := db.Query(q, name)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var riList = []RecipeIngredient{}
	for rows.Next() {
		var rn, c, in string
		err := rows.Scan(&rn, &c, &in)
		if err != nil {
			log.Fatal(err)
		}
		ri := RecipeIngredient{
			RecipeName:     rn,
			Category:       c,
			IngredientName: in,
		}
		riList = append(riList, ri)
	}

	log.Printf("[repo] GetRecipeByName: %v", riList)
	return riList
}

func GetRecipesByCategory(category string) []RecipeIngredient {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password='' dbname=%s sslmode=disable", host, port, user, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// q := "SELECT * FROM categories"
	q := `
	SELECT
		r.name AS recipe_name,
		r.category AS category,
		i.name AS ingredient_name
	FROM
		recipes r
	INNER JOIN
		recipe_ingredients ri ON r.recipe_id = ri.recipe_id
	INNER JOIN
		ingredients i ON ri.ingredient_id = i.ingredient_id
	WHERE
		r.category = $1
	`

	rows, err := db.Query(q, category)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var riList = []RecipeIngredient{}
	for rows.Next() {
		var rn, c, in string
		err := rows.Scan(&rn, &c, &in)
		if err != nil {
			log.Fatal(err)
		}
		ri := RecipeIngredient{
			RecipeName:     rn,
			Category:       c,
			IngredientName: in,
		}
		riList = append(riList, ri)
	}

	log.Printf("[repo] GetRecipesByCategory: %v", riList)
	return riList
}
