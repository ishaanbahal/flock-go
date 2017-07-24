package tests

import "testing"

import roster ".."

const (
	TOKEN = "03d12122-6009-43fb-b28e-79201fed50d1"
)

func TestListContacts(t *testing.T) {
	t.Log("Testing chat fetch messages API")
	post := roster.PostImpl{}
	res, err := post.ListContacts(TOKEN)
	if err != nil {
		t.Errorf("Error occured:\n%s", err)
		return
	}
	t.Logf("Roster List %v", res)
}
