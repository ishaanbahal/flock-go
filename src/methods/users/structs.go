package users

type _getInfo struct {
	Token string `json:"token"`
}

type _getPublicProfile struct {
	Token  string `json:"token"`
	UserID string `json:"userId"`
}

// User struct
type User struct {
	ID           string `json:"id"`
	TeamID       int32  `json:"teamId"`
	Email        string `json:"email"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Role         string `json:"role"`
	Timezone     string `json:"timezone"`
	ProfileImage string `json:"profileImage,omitempty"`
}
