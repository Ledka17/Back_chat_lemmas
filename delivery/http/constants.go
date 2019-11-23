package http

import "time"

const (
	ApiV1UserRegister    = "api/v1/chat/user/register"
	ApiV1UserUpdate      = "api/v1/chat/user/update"
	ApiV1UserGetMessages = "api/v1/chat/user/messages"

	ApiV1SupportGetChats = "api/v1/chat/support/chats"
	ApiV1SupportGetChat  = "api/v1/chat/support/chats/:userId"

	ApiV1ChatStream = "api/v1/chat/stream/:anotherUserID"
)

const (
	SessionIDCookieName   = "user_id"
	SessionIDCookieExpire = 10 * time.Hour
)
