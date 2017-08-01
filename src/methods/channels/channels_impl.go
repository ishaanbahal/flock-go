package channels

import (
	"encoding/json"

	"../common"
)

// PostImpl struct
type PostImpl struct{}

// GetInfo returns info of channel
func (ch PostImpl) GetInfo(token string, channelID string) (*ChannelInfo, error) {
	reqData := _channelInfo{
		Token:     token,
		ChannelID: channelID,
	}

	body, err := common.HTTPRequestPost(common.ChannelsGetInfo, reqData)
	if err != nil {
		return nil, err
	}
	var channelInfo ChannelInfo
	err = json.Unmarshal(body, &channelInfo)
	if err != nil {
		return nil, err
	}
	return &channelInfo, nil
}

func (ch PostImpl) List(token string) (*[]ChannelInfo, error) {
	reqData := _channelLists{
		Token: token,
	}
	body, err := common.HTTPRequestPost(common.ChannelsList, reqData)
	if err != nil {
		return nil, err
	}
	var channelInfo []ChannelInfo
	err = json.Unmarshal(body, &channelInfo)
	if err != nil {
		return nil, err
	}
	return &channelInfo, nil
}

func (ch PostImpl) ListMembers(token string, channelID string, showPublicProfile bool) (*[]ChannelMember, error) {
	reqData := _channelListMembers{
		Token:             token,
		ChannelID:         channelID,
		ShowPublicProfile: showPublicProfile,
	}
	body, err := common.HTTPRequestPost(common.ChannelsListMembers, reqData)
	if err != nil {
		return nil, err
	}
	var channelMembers []ChannelMember
	err = json.Unmarshal(body, &channelMembers)
	if err != nil {
		return nil, err
	}
	return &channelMembers, nil
}
