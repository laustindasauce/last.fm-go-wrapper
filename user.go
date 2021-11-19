package lastfm

import (
	"errors"
)

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
	URL        string         `json:"url"`
	Gender     string         `json:"gender"`
	Name       string         `json:"name"`
	Type       string         `json:"type"`
}

type UserFriend struct {
	Country    string            `json:"country"`
	Age        string            `json:"age"`
	Playcount  string            `json:"playcount"`
	Subscriber string            `json:"subscriber"`
	Realname   string            `json:"realname"`
	Playlists  string            `json:"playlists"`
	Bootstrap  string            `json:"bootstrap"`
	Image      []Image           `json:"image"`
	Registered UserRegisteredStr `json:"registered"`
	URL        string            `json:"url"`
	Gender     string            `json:"gender"`
	Name       string            `json:"name"`
	Type       string            `json:"type"`
}

type ArtistPersonalTags struct {
	ArtistTags ArtistPersonalTag `json:"artists"`
	Attr       PersonalTagAttr   `json:"@attr"`
}

type AlbumPersonalTags struct {
	AlbumTags AlbumPersonalTag `json:"albums"`
	Attr      PersonalTagAttr  `json:"@attr"`
}

type TrackPersonalTags struct {
	TrackTags TrackPersonalTag `json:"tracks"`
	Attr      PersonalTagAttr  `json:"@attr"`
}

type ArtistPersonalTag struct {
	Artists []Artist `json:"artist"`
}

type AlbumPersonalTag struct {
	Albums []Album `json:"album"`
}

type TrackPersonalTag struct {
	Tracks []PersonalTrack `json:"track"`
}

type PersonalTrack struct {
	Name       string          `json:"name"`
	MBID       string          `json:"mbid"`
	URL        string          `json:"url"`
	Duration   string          `json:"duration"`
	Streamable StreamableTrack `json:"streamable"`
	Artist     AlbumArtist     `json:"artist"`
	Image      []Image         `json:"image"`
}

type PersonalTagAttr struct {
	User       string `json:"user"`
	Tag        string `json:"tag"`
	Page       string `json:"page"`
	PerPage    string `json:"perPage"`
	TotalPages string `json:"totalPages"`
	Total      string `json:"total"`
}

type UserLovedTracks struct {
	Tracks []LovedTrack `json:"track"`
	Attr   UserAttr     `json:"@attr"`
}

type LovedTrack struct {
	Artist     AlbumArtist     `json:"artist"`
	Date       TrackDate       `json:"date"`
	MBID       string          `json:"mbid"`
	URL        string          `json:"url"`
	Name       string          `json:"name"`
	Image      []Image         `json:"image"`
	Streamable StreamableTrack `json:"streamable"`
}

type UserFriends struct {
	Attr  UserAttr     `json:"@attr"`
	Users []UserFriend `json:"user"`
}

type UserRegistered struct {
	Unixtime string `json:"unixtime"`
	Text     int    `json:"#text"`
}

type UserRegisteredStr struct {
	Unixtime string `json:"unixtime"`
}

type UserAttr struct {
	User       string `json:"user"`
	Page       string `json:"page"`
	PerPage    string `json:"perPage"`
	TotalPages string `json:"totalPages"`
	Total      string `json:"total"`
}

func (c *Client) UserGetFriends(limit, page string) (UserFriends, error) {
	// http://ws.audioscrobbler.com/2.0/?method=user.getfriends&user=rj&api_key=YOUR_API_KEY&format=json
	var lastfmURL string
	var allOpts []string

	allOpts = append(allOpts, "method.user.getfriends")

	if c.User == "" {
		return UserFriends{}, errors.New("empty user... please run set user method first")
	} else {
		allOpts = append(allOpts, "user."+c.User)
	}

	if limit != "" {
		allOpts = append(allOpts, "limit."+limit)
	}

	if page != "" {
		allOpts = append(allOpts, "page."+page)
	}

	lastfmURL = c.getNoAuthURL(allOpts...)

	var friends struct {
		Friends UserFriends `json:"friends"`
	}

	err := c.get(lastfmURL, &friends)

	if err != nil {
		return UserFriends{}, err
	}

	return friends.Friends, nil
}

func (c *Client) UserGetInfo() (User, error) {
	// http://ws.audioscrobbler.com/2.0/?method=user.getinfo&user=rj&api_key=YOUR_API_KEY&format=json
	var lastfmURL string
	var allOpts []string

	allOpts = append(allOpts, "method.user.getinfo")

	if c.User == "" {
		return User{}, errors.New("empty user... please run set user method first")
	} else {
		allOpts = append(allOpts, "user."+c.User)
	}

	lastfmURL = c.getNoAuthURL(allOpts...)

	var userInfo struct {
		User User `json:"user"`
	}

	err := c.get(lastfmURL, &userInfo)

	if err != nil {
		return User{}, err
	}

	return userInfo.User, nil
}

func (c *Client) UserGetLovedTracks(limit, page string) (UserLovedTracks, error) {
	// http://ws.audioscrobbler.com/2.0/?method=user.getlovedtracks&user=rj&api_key=YOUR_API_KEY&format=json
	var lastfmURL string
	var allOpts []string

	allOpts = append(allOpts, "method.user.getlovedtracks")

	if c.User == "" {
		return UserLovedTracks{}, errors.New("empty user... please run set user method first")
	} else {
		allOpts = append(allOpts, "user."+c.User)
	}

	if limit != "" {
		allOpts = append(allOpts, "limit."+limit)
	}

	if page != "" {
		allOpts = append(allOpts, "page."+page)
	}

	lastfmURL = c.getNoAuthURL(allOpts...)

	var lovedTracks struct {
		UserLovedTracks UserLovedTracks `json:"lovedtracks"`
	}

	err := c.get(lastfmURL, &lovedTracks)

	if err != nil {
		return UserLovedTracks{}, err
	}

	return lovedTracks.UserLovedTracks, nil
}

func (c *Client) UserGetPersonalArtistTags(tag, limit, page string) (ArtistPersonalTags, error) {
	// http://ws.audioscrobbler.com/2.0/?method=user.getpersonaltags&user=rj&tag=rock&taggingtype=artist&api_key=YOUR_API_KEY&format=json
	var lastfmURL string
	var allOpts []string

	allOpts = append(allOpts, "method.user.getpersonaltags", "taggingtype.artist", "tag."+tag)

	if c.User == "" {
		return ArtistPersonalTags{}, errors.New("empty user... please run set user method first")
	} else {
		allOpts = append(allOpts, "user."+c.User)
	}

	if limit != "" {
		allOpts = append(allOpts, "limit."+limit)
	}

	if page != "" {
		allOpts = append(allOpts, "page."+page)
	}

	lastfmURL = c.getNoAuthURL(allOpts...)

	var personalTaggings struct {
		Tags ArtistPersonalTags `json:"taggings"`
	}

	err := c.get(lastfmURL, &personalTaggings)

	if err != nil {
		return ArtistPersonalTags{}, err
	}

	return personalTaggings.Tags, nil
}

func (c *Client) UserGetPersonalAlbumTags(tag, limit, page string) (AlbumPersonalTags, error) {
	// http://ws.audioscrobbler.com/2.0/?method=user.getpersonaltags&user=rj&tag=rock&taggingtype=artist&api_key=YOUR_API_KEY&format=json
	var lastfmURL string
	var allOpts []string

	allOpts = append(allOpts, "method.user.getpersonaltags", "taggingtype.album", "tag."+tag)

	if c.User == "" {
		return AlbumPersonalTags{}, errors.New("empty user... please run set user method first")
	} else {
		allOpts = append(allOpts, "user."+c.User)
	}

	if limit != "" {
		allOpts = append(allOpts, "limit."+limit)
	}

	if page != "" {
		allOpts = append(allOpts, "page."+page)
	}

	lastfmURL = c.getNoAuthURL(allOpts...)

	var personalTaggings struct {
		Tags AlbumPersonalTags `json:"taggings"`
	}

	err := c.get(lastfmURL, &personalTaggings)

	if err != nil {
		return AlbumPersonalTags{}, err
	}

	return personalTaggings.Tags, nil
}

func (c *Client) UserGetPersonalTrackTags(tag, limit, page string) (TrackPersonalTags, error) {
	// http://ws.audioscrobbler.com/2.0/?method=user.getpersonaltags&user=rj&tag=rock&taggingtype=artist&api_key=YOUR_API_KEY&format=json
	var lastfmURL string
	var allOpts []string

	allOpts = append(allOpts, "method.user.getpersonaltags", "taggingtype.track", "tag."+tag)

	if c.User == "" {
		return TrackPersonalTags{}, errors.New("empty user... please run set user method first")
	} else {
		allOpts = append(allOpts, "user."+c.User)
	}

	if limit != "" {
		allOpts = append(allOpts, "limit."+limit)
	}

	if page != "" {
		allOpts = append(allOpts, "page."+page)
	}

	lastfmURL = c.getNoAuthURL(allOpts...)

	var personalTaggings struct {
		Tags TrackPersonalTags `json:"taggings"`
	}

	err := c.get(lastfmURL, &personalTaggings)

	if err != nil {
		return TrackPersonalTags{}, err
	}

	return personalTaggings.Tags, nil
}
