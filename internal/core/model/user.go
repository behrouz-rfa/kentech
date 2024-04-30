package model

import (
	"time"
)

// User represents a user in the system
type User struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Username  string    `json:"username" `
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Password  string    `json:"-"`
	Films     []Film    `json:"films"`
	JwtToken  string    `json:"jwtToken,omitempty"`
}

type UserInput struct {
	Username  string `json:"username" `
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Password  string `json:"password"`
}

type UserUpdateInput struct {
	Username  *string `json:"username"`
	Firstname *string `json:"firstname"`
	Lastname  *string `json:"lastname"`
}

type UserLoginInput struct {
	Username string `json:"username" `
	Password string `json:"password"`
}
