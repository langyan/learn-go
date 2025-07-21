package model

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
type UpdateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
