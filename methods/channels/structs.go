package channels

import "../common"

type _channelInfo struct {
	Token     string `json:"token"`
	ChannelID string `json:"channelId"`
}

type _channelLists struct {
	Token string `json:"token"`
}

type _channelListMembers struct {
	Token             string `json:"token"`
	ChannelID         string `json:"channelId"`
	ShowPublicProfile bool   `json:"showPublidProfile"`
}

// ChannelInfo struct for response from getInfo API
type ChannelInfo struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	MemberCount  int32  `json:"memberCount"`
	ProfileImage string `json:"profileImage"`
}

// ChannelMember struct
type ChannelMember struct {
	UserID        string         `json:"userId"`
	Affiliation   string         `json:"affiliation"`
	PublicProfile common.Profile `json:"publicProfile"`
}
