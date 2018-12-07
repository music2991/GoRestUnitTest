package model

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	Address  string `json:"address"`
}
