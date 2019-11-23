package chat

import "time"

type Repository interface {
	CreateMessage(userFromID, userToID int, text string, time time.Time) error
}
