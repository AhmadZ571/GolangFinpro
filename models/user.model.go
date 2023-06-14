package models

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"Name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserInput struct {
	Name     string `json:"Name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
