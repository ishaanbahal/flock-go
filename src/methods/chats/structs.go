package chats

type _fetchMessages struct {
	Token string   `json:"token"`
	Chat  string   `json:"chant"`
	UIDs  []string `json:"uids"`
}

// Message struct
type Message struct {
	From        string       `json:"from"`
	To          string       `json:"to"`
	ID          string       `json:"id"`
	UID         string       `json:"uid"`
	Text        string       `json:"text"`
	Timestamp   int32        `json:"timestamp"`
	Attachments []Attachment `json:"attachments,omitempty"`
}

// Attachment struct
type Attachment struct {
	ID          string     `json:"id,omitempty"`
	AppID       string     `json:"appId,omitempty"`
	Title       string     `json:"title,omitempty"`
	Description string     `json:"description,omitempty"`
	Color       string     `json:"color,omitempty"`
	Views       Views      `json:"views,omitempty"`
	URL         string     `json:"url,omitempty"`
	Forward     bool       `json:"forward,omitempty"`
	Downloads   []Download `json:"downloads,omitmpty"`
	Buttons     []Button   `json:"buttons,omitempty"`
}

// Views struct
type Views struct {
	Widget  Widget      `json:"widget,omitempty"`
	HTML    HTML        `json:"html,omitempty"`
	FlockML string      `json:"flockml,omitempty"`
	Image   ImageHolder `json:"image,omitempty"`
}

// HTML struct
type HTML struct {
	Inline string `json:"string"`
	Width  string `json:"width"`
	Height string `json:"height"`
}

// Widget struct
type Widget struct {
	Src    string `json:"src"`
	Width  int32  `json:"width"`
	Height int32  `json:"height"`
}

// Image struct
type Image struct {
	Src    string `json:"src"`
	Width  int32  `json:"width"`
	Height int32  `json:"height"`
}

// ImageHolder struct
type ImageHolder struct {
	Original  Image  `json:"original"`
	Thumbnail Image  `json:"thumbnail,omitempty"`
	Filename  string `json:"filename,omitempty"`
}

// Download struct
type Download struct {
	Src      string `json:"src"`
	Mime     string `json:"mime"`
	Filename string `json:"filename"`
	Size     string `json:"size"`
}

// Button struct
type Button struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
	// See: https://docs.flock.com/display/flockos/Attachment#Attachment-ActionConfig
	Action map[string]interface{} `json:"action"`
}

type SendAs struct {
	Name         string `json:"name"`
	ProfileImage string `json:"profileImage"`
}

type SendMessage struct {
	Token        string       `json:"token"`
	To           string       `json:"to"`
	Text         string       `json:"text"`
	FlockML      string       `json:"flockml,omitempty"`
	Notification string       `json:"notification,omitempty"`
	Mentions     []string     `json:"mentions,omitempty"`
	SendAs       SendAs       `json:"sendAs"`
	Attachments  []Attachment `json:"attachments,omitempty"`
}

type SendMessageResponse struct {
	UID string `json:"uid"`
}
