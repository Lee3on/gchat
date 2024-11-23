package model

type Sender struct {
	UserId    int64  // Sender ID
	DeviceId  int64  // Sender's device ID
	Nickname  string // Nickname
	AvatarUrl string // Avatar URL
	Extra     string // Additional fields
}
