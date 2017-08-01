package users

import "../common"

type Users interface {
	GetInfo(token string) (User, error)

	GetPublicProfile(token, userID string) (common.Profile, error)
}
