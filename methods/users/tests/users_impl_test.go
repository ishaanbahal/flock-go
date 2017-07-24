package tests

import "testing"
import user ".."

const (
	TOKEN   = ""
	USER_ID = ""
)

func TestGetInfo(t *testing.T) {
	t.Log("Testing User getInfo API")
	post := user.PostImpl{}
	res, err := post.GetInfo(TOKEN)
	if err != nil {
		t.Errorf("Error occured:\n%s", err)
		return
	}
	t.Logf("User %v", res)
}

func TestGetPublicProfile(t *testing.T) {
	t.Log("Testing User getPublicProfile API")
	post := user.PostImpl{}
	res, err := post.GetPublicProfile(TOKEN, USER_ID)
	if err != nil {
		t.Errorf("Error occured:\n%s", err)
		return
	}
	t.Logf("Public Profile %v", res)
}
