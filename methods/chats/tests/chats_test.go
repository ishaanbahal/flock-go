package tests

import "testing"
import chats ".."

const (
	TOKEN      = ""
	CHAT       = ""
	CHANNEL_ID = ""
)

func TestFetchMessags(t *testing.T) {
	t.Log("Testing chat fetch messages API")
	post := chats.PostImpl{}
	res, err := post.FetchMessages(TOKEN, CHAT, []string{"1500894415927-f2G-m203"})
	if err != nil {
		t.Errorf("Error occured:\n%s", err)
		return
	}
	t.Logf("Data found: %v", res)
}

func TestSendMessage(t *testing.T) {
	t.Log("Testing chat fetch messages API")
	post := chats.PostImpl{}
	testData := chats.SendMessage{
		Token:    TOKEN,
		To:       CHANNEL_ID,
		Text:     "test message",
		Mentions: make([]string, 0),
	}
	res, err := post.SendMessage(testData)
	if err != nil {
		t.Errorf("Error occured:\n%s", err)
		return
	}
	t.Logf("Data found: %v", res)
}
