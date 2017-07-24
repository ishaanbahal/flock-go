package tests

import "testing"

import roster ".."

const (
	TOKEN = ""
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
