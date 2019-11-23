package model

type User struct {
	ID        int    `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	Email     string `json:"email" db:"email"`
	IsSupport bool   `json:"is_support" db:"is_support"`
}
