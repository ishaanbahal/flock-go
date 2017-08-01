package common

// Profile struct
type Profile struct {
	ID           string `json:"id"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	ProfileImage string `json:"profileImage,omitempty"`
}

type ErrorResponse struct {
	ErrorName   string `json:"error,omitempty"`
	Description string `json:"description,omitempty"`
	Parameter   string `json:"parameter,omitempty"`
}
