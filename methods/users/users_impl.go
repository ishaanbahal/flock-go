package users

import "encoding/json"
import "../common"

// PostImpl strut
type PostImpl struct{}

// GetInfo returns info on current user based on Token ID
func (ch PostImpl) GetInfo(token string) (User, error) {
	var user User
	reqData := _getInfo{
		Token: token,
	}
	body, err := common.HTTPRequestPost(common.UsersGetInfo, reqData)
	if err != nil {
		return user, err
	}
	err = json.Unmarshal(body, &user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (ch PostImpl) GetPublicProfile(token, userID string) (common.Profile, error) {
	var userProfile common.Profile
	reqData := _getPublicProfile{
		Token:  token,
		UserID: userID,
	}
	body, err := common.HTTPRequestPost(common.UsersGetPublicProfile, reqData)
	if err != nil {
		return userProfile, err
	}
	err = json.Unmarshal(body, &userProfile)
	if err != nil {
		return userProfile, err
	}
	return userProfile, nil
}
