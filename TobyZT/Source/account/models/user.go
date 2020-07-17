package models

type User struct {
	UserID int64 `json:"userid"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}