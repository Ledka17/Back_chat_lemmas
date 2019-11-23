package model

import "time"

type User struct {
	ID        int    `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	Email     string `json:"email" db:"email"`
	IsSupport bool   `json:"is_support" db:"is_support"`
}

type Message struct {
	ID         int       `json:"id" db:"id"`
	UserFromID int       `json:"user_from_id" db:"user_from_id"`
	UserToID   int       `json:"user_to_id" db:"user_to_id"`
	Text       string    `json:"text" db:"text"`
	Time       time.Time `json:"time" db:"time"`
}
