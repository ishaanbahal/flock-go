package channels

// Channels interface for channels implementations
type Channels interface {
	GetInfo(token string, channelID string) (*ChannelInfo, error)

	List(token string) (*[]ChannelInfo, error)

	ListMembers(token string, channelID string, showPublicProfile bool) (*[]ChannelMember, error)
}
