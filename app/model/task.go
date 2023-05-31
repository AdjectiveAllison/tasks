//go:generate easyjson .
package model

//easyjson:json
type Task struct {
	ID          uint64   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Links       []string `json:"links"`
	UpdatedAt   uint64   `json:"updatedAt"`
	Completed   bool     `json:"completed"`
}

//easyjson:json
type ListTasksResponse struct {
	Tasks []Task `json:"tasks"`
}
