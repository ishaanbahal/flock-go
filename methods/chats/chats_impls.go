package chats

import "encoding/json"
import "../common"

// PostImpl struct
type PostImpl struct{}

func (ch PostImpl) FetchMessages(token string, chat string, uids []string) (*[]Message, error) {
	reqData := _fetchMessages{
		Token: token,
		Chat:  chat,
		UIDs:  uids,
	}

	body, err := common.HTTPRequestPost(common.ChatFetchMessages, reqData)
	if err != nil {
		return nil, err
	}
	var messages []Message
	err = json.Unmarshal(body, &messages)
	if err != nil {
		return nil, err
	}
	return &messages, nil
}

func (ch PostImpl) SendMessage(sendMessage SendMessage) (SendMessageResponse, error) {
	var response SendMessageResponse
	body, err := common.HTTPRequestPost(common.ChatSendMessage, sendMessage)
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}
