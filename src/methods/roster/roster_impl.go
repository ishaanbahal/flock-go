package roster

import (
	"encoding/json"

	"../common"
)

// PostImpl struct
type PostImpl struct{}

// ListContacts lists the contacts for profile
func (ch PostImpl) ListContacts(token string) (*[]common.Profile, error) {
	reqData := _listContacts{
		Token: token,
	}
	body, err := common.HTTPRequestPost(common.RosterListContacts, reqData)
	if err != nil {
		return nil, err
	}
	var profiles []common.Profile
	err = json.Unmarshal(body, &profiles)
	if err != nil {
		return nil, err
	}
	return &profiles, nil
}
