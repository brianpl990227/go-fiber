package dto

type CreateTaskDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool `json:"done"`
	UserID      string `json:"userId"`
}