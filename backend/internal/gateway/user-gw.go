package gateway

type CreateUser struct {
	Name     string `json:"name"`
	Surname  string
	Password string
	Email    string `json:"email`
}

type LoginInput struct {
	Email    string
	Password string
}
