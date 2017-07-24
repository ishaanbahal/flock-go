package test

import "testing"
import channels ".."

const (
	TOKEN      = ""
	CHANNEL_ID = ""

	TOKEN_GARBAGE      = "random string"
	CHANNEL_ID_GARBAGE = "random string"
)

func TestGetInfoNegative(t *testing.T) {
	t.Log("Testing wrong channel getInfo")
	post := channels.PostImpl{}
	_, err := post.GetInfo(TOKEN_GARBAGE, CHANNEL_ID_GARBAGE)
	if err != nil {
		return
	}
}

func TestGetInfoPositive(t *testing.T) {
	t.Log("Testing correct channel getInfo")
	post := channels.PostImpl{}
	ret, err := post.GetInfo(TOKEN, CHANNEL_ID)
	if err != nil {
		t.Errorf("Error occured:\n%s", err)
		return
	}
	t.Logf("Channel info -> %v", ret)
}

func TestListNegative(t *testing.T) {
	t.Log("Testing wrong channel listing")
	post := channels.PostImpl{}
	_, err := post.List(TOKEN_GARBAGE)
	if err != nil {
		return
	}
}

func TestListPositive(t *testing.T) {
	t.Log("Testing correct channel listing")
	post := channels.PostImpl{}
	ret, err := post.List(TOKEN)
	if err != nil {
		t.Errorf("Error occured:\n%s", err)
		return
	}
	t.Logf("ID of first channel %v", (*ret)[0].ID)
}

func TestListMembersNegative(t *testing.T) {
	t.Log("Testing wrong channel listing")
	post := channels.PostImpl{}
	_, err := post.ListMembers(TOKEN_GARBAGE, CHANNEL_ID_GARBAGE, true)
	if err != nil {
		return
	}
}

func TestListMembersPositive(t *testing.T) {
	t.Log("Testing correct channel list members")
	post := channels.PostImpl{}
	ret, err := post.ListMembers(TOKEN, CHANNEL_ID, true)
	if err != nil {
		t.Errorf("Error occured:\n%s", err)
		return
	}
	t.Logf("User ID of first user %v", (*ret)[0].UserID)
}
