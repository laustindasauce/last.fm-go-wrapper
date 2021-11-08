package lastfm

type User struct {
	Country    string         `json:"country"`
	Age        string         `json:"age"`
	Playcount  string         `json:"playcount"`
	Subscriber string         `json:"subscriber"`
	Realname   string         `json:"realname"`
	Playlists  string         `json:"playlists"`
	Bootstrap  string         `json:"bootstrap"`
	Image      []Image        `json:"image"`
	Registered UserRegistered `json:"registered"`
	Url        string         `json:"url"`
	Gender     string         `json:"gender"`
	Name       string         `json:"name"`
	Type       string         `json:"type"`
}

type UserRegistered struct {
	Unixtime string `json:"unixtime"`
	Text     int    `json:"#text"`
}

type UserAttr struct {
	User       string `json:"user"`
	Page       string `json:"page"`
	PerPage    string `json:"perPage"`
	TotalPages string `json:"totalPages"`
	Total      string `json:"total"`
}
