package api

type Recipe struct {
	Category    string   `json:"category"`
	Name        string   `json:"name"`
	Ingredients []string `json:"ingredients"`
}
