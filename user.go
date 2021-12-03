package lastfm

import (
	"fmt"
	"strings"
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

/*
user (Required) : The last.fm username to fetch the friends of.

recenttracks (Optional) : Whether or not to include information about friends' recent listening in the response.

limit (Optional) : The number of results to fetch per page. Defaults to 50.

page (Optional) : The page number to fetch. Defaults to first page.

api_key (Required) : A Last.fm API key.
*/
func (c *Client) UserGetFriends(user string, opts ...RequestOption) (UserFriends, error) {
	// http://ws.audioscrobbler.com/2.0/?method=user.getfriends&user=rj&api_key=YOUR_API_KEY&format=json
	lastfmURL := fmt.Sprintf("%s&user=%s&method=user.getfriends", c.baseApiURL, user)

	values := processOptions(opts...).urlParams

	if query := values.Encode(); query != "" {
		lastfmURL += "&" + query
	}

	var friends struct {
		Friends UserFriends `json:"friends"`
	}

	err := c.get(lastfmURL, &friends)

	if err != nil {
		return UserFriends{}, err
	}

	return friends.Friends, nil
}

/*
user (Optional) : The user to fetch info for. Defaults to the authenticated user.

api_key (Required) : A Last.fm API key.
*/
func (c *Client) UserGetInfo(user string, opts ...RequestOption) (User, error) {
	// http://ws.audioscrobbler.com/2.0/?method=user.getinfo&user=rj&api_key=YOUR_API_KEY&format=json
	lastfmURL := fmt.Sprintf("%s&user=%s&method=user.getinfo", c.baseApiURL, user)

	values := processOptions(opts...).urlParams

	if query := values.Encode(); query != "" {
		lastfmURL += "&" + query
	}

	var userInfo struct {
		User User `json:"user"`
	}

	err := c.get(lastfmURL, &userInfo)

	if err != nil {
		return User{}, err
	}

	return userInfo.User, nil
}

/*
user (Required) : The user name to fetch the loved tracks for.

limit (Optional) : The number of results to fetch per page. Defaults to 50.

page (Optional) : The page number to fetch. Defaults to first page.

api_key (Required) : A Last.fm API key.
*/
func (c *Client) UserGetLovedTracks(user string, opts ...RequestOption) (UserLovedTracks, error) {
	// http://ws.audioscrobbler.com/2.0/?method=user.getlovedtracks&user=rj&api_key=YOUR_API_KEY&format=json
	lastfmURL := fmt.Sprintf("%s&user=%s&method=user.getlovedtracks", c.baseApiURL, user)

	values := processOptions(opts...).urlParams

	if query := values.Encode(); query != "" {
		lastfmURL += "&" + query
	}

	var lovedTracks struct {
		UserLovedTracks UserLovedTracks `json:"lovedtracks"`
	}

	err := c.get(lastfmURL, &lovedTracks)

	if err != nil {
		return UserLovedTracks{}, err
	}

	return lovedTracks.UserLovedTracks, nil
}

/*
user (Required) : The user who performed the taggings.

tag (Required) : The tag you're interested in.

taggingtype=artist : The type of items which have been tagged

limit (Optional) : The number of results to fetch per page. Defaults to 50.

page (Optional) : The page number to fetch. Defaults to first page.

api_key (Required) : A Last.fm API key.
*/
func (c *Client) UserGetPersonalArtistTags(user, tag string, opts ...RequestOption) (ArtistPersonalTags, error) {
	// http://ws.audioscrobbler.com/2.0/?method=user.getpersonaltags&user=rj&tag=rock&taggingtype=artist&api_key=YOUR_API_KEY&format=json
	lastfmURL := fmt.Sprintf("%s&user=%s&method=user.getpersonaltags&tag=%s", c.baseApiURL, user, tag)

	// Add taggingtype
	opts = append(opts, TaggingTypeOpt(ArtistTag))

	values := processOptions(opts...).urlParams

	if query := values.Encode(); query != "" {
		lastfmURL += "&" + query
	}

	var personalTaggings struct {
		Tags ArtistPersonalTags `json:"taggings"`
	}

	err := c.get(lastfmURL, &personalTaggings)

	if err != nil {
		return ArtistPersonalTags{}, err
	}

	return personalTaggings.Tags, nil
}

/*
user (Required) : The user who performed the taggings.

tag (Required) : The tag you're interested in.

taggingtype=album : The type of items which have been tagged

limit (Optional) : The number of results to fetch per page. Defaults to 50.

page (Optional) : The page number to fetch. Defaults to first page.

api_key (Required) : A Last.fm API key.
*/
func (c *Client) UserGetPersonalAlbumTags(user, tag string, opts ...RequestOption) (AlbumPersonalTags, error) {
	// http://ws.audioscrobbler.com/2.0/?method=user.getpersonaltags&user=rj&tag=rock&taggingtype=artist&api_key=YOUR_API_KEY&format=json

	lastfmURL := fmt.Sprintf("%s&user=%s&method=user.getpersonaltags&tag=%s", c.baseApiURL, user, tag)

	// Add taggingtype
	opts = append(opts, TaggingTypeOpt(AlbumTag))

	values := processOptions(opts...).urlParams

	if query := values.Encode(); query != "" {
		lastfmURL += "&" + query
	}

	var personalTaggings struct {
		Tags AlbumPersonalTags `json:"taggings"`
	}

	err := c.get(lastfmURL, &personalTaggings)

	if err != nil {
		return AlbumPersonalTags{}, err
	}

	return personalTaggings.Tags, nil
}

/*
user (Required) : The user who performed the taggings.

tag (Required) : The tag you're interested in.

taggingtype=track : The type of items which have been tagged

limit (Optional) : The number of results to fetch per page. Defaults to 50.

page (Optional) : The page number to fetch. Defaults to first page.

api_key (Required) : A Last.fm API key.
*/
func (c *Client) UserGetPersonalTrackTags(user, tag string, opts ...RequestOption) (TrackPersonalTags, error) {
	// http://ws.audioscrobbler.com/2.0/?method=user.getpersonaltags&user=rj&tag=rock&taggingtype=artist&api_key=YOUR_API_KEY&format=json
	lastfmURL := fmt.Sprintf("%s&user=%s&method=user.getpersonaltags&tag=%s", c.baseApiURL, user, tag)

	// Add taggingtype
	opts = append(opts, TaggingTypeOpt(TrackTag))

	values := processOptions(opts...).urlParams

	if query := values.Encode(); query != "" {
		lastfmURL += "&" + query
	}

	var personalTaggings struct {
		Tags TrackPersonalTags `json:"taggings"`
	}

	err := c.get(lastfmURL, &personalTaggings)

	if err != nil {
		return TrackPersonalTags{}, err
	}

	return personalTaggings.Tags, nil
}

/*
limit (Optional) : The number of results to fetch per page. Defaults to 50. Maximum is 200.

user (Required) : The last.fm username to fetch the recent tracks of.

page (Optional) : The page number to fetch. Defaults to first page.

from (Optional) : Beginning timestamp of a range - only display scrobbles after this time, in UNIX timestamp format (integer number of seconds since 00:00:00, January 1st 1970 UTC). This must be in the UTC time zone.)

extended (0|1) (Optional) : Includes extended data in each artist, and whether or not the user has loved each track

to (Optional) : End timestamp of a range - only display scrobbles before this time, in UNIX timestamp format (integer number of seconds since 00:00:00, January 1st 1970 UTC). This must be in the UTC time zone.)

api_key (Required) : A Last.fm API key.
*/
func (c *Client) UserGetRecentTracks(user string, opts ...RequestOption) {
	// http://ws.audioscrobbler.com/2.0/?method=user.getrecenttracks&user=rj&api_key=YOUR_API_KEY&format=json
	lastfmURL := fmt.Sprintf("%s&user=%s&method=user.getrecenttracks", c.baseApiURL, user)
	values := processOptions(opts...).urlParams

	if query := values.Encode(); query != "" {
		lastfmURL += "&" + query
	}

	if strings.Contains(lastfmURL, "extended=1") {
		fmt.Println("Contains")
		// Need to set up for extended
	} else {
		// Get non extended info
		fmt.Println("Not contained")
	}

	fmt.Println(lastfmURL)
}
