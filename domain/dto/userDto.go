package dto

type CreateUserDTO struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email string `json:"email"`
}
