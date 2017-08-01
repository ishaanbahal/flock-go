package methods

import "testing"
import channels "../../src/methods/channels"
import config ".."

func TestGetInfoNegative(t *testing.T) {
	t.Log("Testing wrong channel getInfo")
	post := channels.PostImpl{}
	_, err := post.GetInfo(config.TOKEN_GARBAGE, config.CHANNEL_ID_GARBAGE)
	if err != nil {
		return
	}
}

func TestGetInfoPositive(t *testing.T) {
	t.Log("Testing correct channel getInfo")
	post := channels.PostImpl{}
	ret, err := post.GetInfo(config.TOKEN, config.CHANNEL_ID)
	if err != nil {
		t.Errorf("Error occured:\n%s", err)
		return
	}
	t.Logf("Channel info -> %v", ret)
}

func TestListNegative(t *testing.T) {
	t.Log("Testing wrong channel listing")
	post := channels.PostImpl{}
	_, err := post.List(config.TOKEN_GARBAGE)
	if err != nil {
		return
	}
}

func TestListPositive(t *testing.T) {
	t.Log("Testing correct channel listing")
	post := channels.PostImpl{}
	ret, err := post.List(config.TOKEN)
	if err != nil {
		t.Errorf("Error occured:\n%s", err)
		return
	}
	t.Logf("ID of first channel %v", (*ret)[0].ID)
}

func TestListMembersNegative(t *testing.T) {
	t.Log("Testing wrong channel listing")
	post := channels.PostImpl{}
	_, err := post.ListMembers(config.TOKEN_GARBAGE, config.CHANNEL_ID_GARBAGE, true)
	if err != nil {
		return
	}
}

func TestListMembersPositive(t *testing.T) {
	t.Log("Testing correct channel list members")
	post := channels.PostImpl{}
	ret, err := post.ListMembers(config.TOKEN, config.CHANNEL_ID, true)
	if err != nil {
		t.Errorf("Error occured:\n%s", err)
		return
	}
	t.Logf("User ID of first user %v", (*ret)[0].UserID)
}
