package http

import "time"

const (
	ApiPV1 = "api/v1"

	ApiV1Chat = ApiPV1 + "/chat"

	ApiV1User = ApiV1Chat + "/user"
	ApiV1Supporter = ApiV1Chat + "/supporter"

	ApiV1UserGetMessages = ApiV1User + "/messages"
	APiV1UserSendMessage = ApiV1User + "/send"
	ApiV1SupporterGetChats = ApiV1Supporter + "/chats"
	ApiV1SupporterGetChat = ApiV1Supporter + "/chats/{id}"
)

const (
	SessionIDCookieName   = "user_id"
	SessionIDCookieExpire = 10 * time.Hour
)