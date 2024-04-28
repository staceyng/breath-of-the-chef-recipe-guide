package api

type ErrorResponse struct {
	Code   string `json:"code,omitempty"`
	Title  string `json:"title,omitempty"`
	Detail string `json:"detail,omitempty"`
}

type JsonResponse struct {
	Data   interface{}     `json:"data,omitempty"`
	Errors []ErrorResponse `json:"errors,omitempty"`
	Meta   interface{}     `json:"meta,omitempty"`
}

type Recipe struct {
	Category    string   `json:"category"`
	Name        string   `json:"name"`
	Ingredients []string `json:"ingredients"`
}
