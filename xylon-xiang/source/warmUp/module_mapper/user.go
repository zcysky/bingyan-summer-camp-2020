package module_mapper

type User struct {
	ID       string
	Password string
	Name     string
	Phone    string
	Email    string
	Admin    int
}

type RegisterUser struct {
	RegisterCode string `json:"register_code"`
	User
}
