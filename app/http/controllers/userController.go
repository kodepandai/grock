package controllers

type UserController struct{}

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	IsActive bool   `json:"is_active"`
}

func (c *UserController) Init() {
	// TODO: register controller middleweares here
}

func (c *UserController) ListUser() []User {
	return []User{{ID: 1, Email: "grock@kodepandai.com", IsActive: true}}
}
