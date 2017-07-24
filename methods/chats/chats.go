package chats

type Chats interface {
	FetchMessages(token string, chat string, uids []string) (*[]Message, error)

	SendMessage(sendMessage SendMessage) (SendMessageResponse, error)
}
