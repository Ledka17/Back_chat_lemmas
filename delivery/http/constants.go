package http

import "time"

const (
	ApiPV1 = "api/v1"

	ApiV1Chat = ApiPV1 + "/chat"

	ApiV1User    = ApiV1Chat + "/user"
	ApiV1Support = ApiV1Chat + "/support"

	ApiV1UserGetMessages = ApiV1User + "/messages"
	APiV1UserSendMessage = ApiV1User + "/send"
	ApiV1SupportGetChats = ApiV1Support + "/chats"
	ApiV1SupportGetChat  = ApiV1Support + "/chats/{id}"
)

const (
	SessionIDCookieName   = "user_id"
	SessionIDCookieExpire = 10 * time.Hour
)
