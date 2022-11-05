package model

type UserData struct {
	ID         int    `json:"id"`
	Name       string `json:"name"  validate:"required" example:"Some name"`
	Email      string `json:"email"  validate:"required" example:"Some name"`
	Password   string `json:"password"  validate:"required" example:"Some name"`
	Picture    int    `json:"picture"  validate:"required" example:"Some name"`
	Newsletter bool   `json:"newsletter"  validate:"required" example:"Some name"`
}

type Users struct {
	Users []UserData `json:"users"`
}
