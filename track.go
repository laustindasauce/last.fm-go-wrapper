package lastfm

import (
	"fmt"
)

type Tracks struct {
	Track []Track    `json:"track"`
	Attr  ArtistAttr `json:"@attr"`
}

type Track struct {
	Name       string      `json:"name"`
	Playcount  string      `json:"playcount"`
	Listeners  string      `json:"listeners"`
	MBID       string      `json:"mbid"`
	URL        string      `json:"url"`
	Streamable string      `json:"streamable"`
	Artist     AlbumArtist `json:"artist"`
	Image      []Image     `json:"image"`
	Attr       TrackRank   `json:"@attr"`
}

type TrackInfo struct {
	Name       string          `json:"name"`
	MBID       string          `json:"mbid"`
	URL        string          `json:"url"`
	Duration   string          `json:"duration"`
	Streamable StreamableTrack `json:"streamable"`
	Listeners  string          `json:"listeners"`
	Playcount  string          `json:"playcount"`
	Artist     AlbumArtist     `json:"artist"`
	Album      TrackInfoAlbum  `json:"album"`
	TopTags    Tags            `json:"toptags"`
	Wiki       Wiki            `json:"wiki"`
}

type SimilarTracks struct {
	Tracks []SimilarTrack   `json:"track"`
	Attr   SimpleArtistAttr `json:"@attr"`
}

type TrackTags struct {
	Tags []Tag           `json:"tag,omitempty"`
	Text string          `json:"#text,omitempty"`
	Attr ArtistTrackAttr `json:"@attr"`
}

type TrackTopTags struct {
	Tags []TagWithCount  `json:"tag"`
	Attr ArtistTrackAttr `json:"@attr"`
}

type TrackSearchRes struct {
	Query             OpenSearchQuery `json:"opensearch:Query"`
	QueryTotalResults string          `json:"opensearch:totalResults"`
	QueryStartIndex   string          `json:"opensearch:startIndex"`
	QueryItemsPerPage string          `json:"opensearch:itemsPerPage"`
	TrackMatches      TrackMatches    `json:"trackmatches"`
	Attr              SearchAttr      `json:"@attr"`
}

type TrackMatches struct {
	Tracks []TrackMatch `json:"track"`
}

type TrackMatch struct {
	Name       string  `json:"name"`
	Artist     string  `json:"artist"`
	URL        string  `json:"url"`
	Streamable string  `json:"streamable"`
	Listeners  string  `json:"listeners"`
	Image      []Image `json:"image"`
	MBID       string  `json:"mbid"`
}

type SimilarTrack struct {
	Name       string          `json:"name"`
	Playcount  int             `json:"playcount"`
	MBID       string          `json:"mbid"`
	Match      float32         `json:"match"`
	URL        string          `json:"url"`
	Streamable StreamableTrack `json:"streamable"`
	Duration   int             `json:"duration"`
	Artist     AlbumArtist     `json:"artist"`
	Image      []Image         `json:"image"`
}

type TrackCorrection struct {
	Track CorrectedTrack `json:"track"`
	Attr  CorrectedAttr  `json:"@attr"`
}

type CorrectedTrack struct {
	Name   string      `json:"name"`
	URL    string      `json:"url"`
	MBID   string      `json:"mbid"`
	Artist AlbumArtist `json:"artist"`
}

type TrackDate struct {
	UTS  string `json:"uts"`
	Text string `json:"#text"`
}

type CorrectedAttr struct {
	Index           string `json:"index"`
	ArtistCorrected string `json:"artistcorrected"`
	TrackCorrected  string `json:"trackcorrected"`
}

type StreamableTrack struct {
	FullTrack string `json:"fulltrack"`
	Text      string `json:"#text"`
}

type TrackAlbum struct {
	Streamable StreamableTrack `json:"streamable"`
	Duration   int             `json:"duration"`
	URL        string          `json:"url"`
	Name       string          `json:"name"`
	Attr       TrackRankInt    `json:"@attr"`
	Artist     AlbumArtist     `json:"artist"`
}

type TrackInfoAlbum struct {
	Artist string  `json:"artist"`
	Title  string  `json:"title"`
	MBID   string  `json:"mbid"`
	URL    string  `json:"url"`
	Image  []Image `json:"image"`
	Attr   PosAttr `json:"@attr"`
}

type TrackRank struct {
	Rank string `json:"rank"`
}

type TrackRankInt struct {
	Rank int `json:"rank"`
}

type SimpleArtistAttr struct {
	Artist string `json:"artist"`
}

type ArtistTrackAttr struct {
	Artist string `json:"artist"`
	Track  string `json:"track"`
}

/*
artist (Required) : The artist name to correct.

track (Required) : The track name to correct.

api_key (Required) : A Last.fm API key.
*/
func (c *Client) TrackGetCorrection(track, artist string) (TrackCorrection, error) {
	// http://ws.audioscrobbler.com/2.0/?method=track.getcorrection&artist=guns and roses&track=Mrbrownstone&api_key=YOUR_API_KEY&format=json
	lastfmURL := fmt.Sprintf("%s&method=track.getcorrection", c.baseApiURL)

	var opts []RequestOption

	opts = append(opts, ArtistOpt(artist), TrackOpt(track))

	values := processOptions(opts...).urlParams

	if query := values.Encode(); query != "" {
		lastfmURL += "&" + query
	}

	type correction struct {
		TrackCorrection TrackCorrection `json:"correction"`
	}

	var corrections struct {
		Corrections correction `json:"corrections"`
	}

	err := c.get(lastfmURL, &corrections)

	if err != nil {
		return TrackCorrection{}, err
	}

	return corrections.Corrections.TrackCorrection, nil
}

/*
mbid (Optional) : The musicbrainz id for the track

track (Required (unless mbid)] : The track name

artist (Required (unless mbid)] : The artist name

username (Optional) : The username for the context of the request. If supplied, the user's playcount for this track and whether they have loved the track is included in the response.

autocorrect[0|1] (Optional) : Transform misspelled artist and track names into correct artist and track names, returning the correct version instead. The corrected artist and track name will be returned in the response.

api_key (Required) : A Last.fm API key.
*/
func (c *Client) TrackGetInfo(track, artist string, opts ...RequestOption) (TrackInfo, error) {
	// http://ws.audioscrobbler.com/2.0/?method=track.getInfo&api_key=YOUR_API_KEY&artist=cher&track=believe&format=json
	lastfmURL := fmt.Sprintf("%s&method=track.getInfo", c.baseApiURL)

	opts = append(opts, ArtistOpt(artist), TrackOpt(track))

	values := processOptions(opts...).urlParams

	if query := values.Encode(); query != "" {
		lastfmURL += "&" + query
	}

	var trackInfo struct {
		Track TrackInfo `json:"track"`
	}

	err := c.get(lastfmURL, &trackInfo)

	if err != nil {
		return TrackInfo{}, err
	}

	return trackInfo.Track, nil
}

/*
track (Required (unless mbid)] : The track name

artist (Required (unless mbid)] : The artist name

mbid (Optional) : The musicbrainz id for the track

autocorrect[0|1] (Optional) : Transform misspelled artist and track names into correct artist and track names, returning the correct version instead. The corrected artist and track name will be returned in the response.

limit (Optional) : Maximum number of similar tracks to return

api_key (Required) : A Last.fm API key.
*/
func (c *Client) TrackGetSimilar(track, artist string, opts ...RequestOption) (SimilarTracks, error) {
	// http://ws.audioscrobbler.com/2.0/?method=track.getsimilar&artist=cher&track=believe&api_key=YOUR_API_KEY&format=json
	lastfmURL := fmt.Sprintf("%s&method=track.getsimilar&track=%s&artist=%s", c.baseApiURL, track, artist)

	values := processOptions(opts...).urlParams

	if query := values.Encode(); query != "" {
		lastfmURL += "&" + query
	}

	var similarTrack struct {
		SimilarTracks SimilarTracks `json:"similartracks"`
	}

	err := c.get(lastfmURL, &similarTrack)

	if err != nil {
		fmt.Println(err.Error())
		return SimilarTracks{}, err
	}

	return similarTrack.SimilarTracks, nil
}

/*
artist (Required (unless mbid)] : The artist name

track (Required (unless mbid)] : The track name

mbid (Optional) : The musicbrainz id for the track

autocorrect[0|1] (Optional) : Transform misspelled artist and track names into correct artist and track names, returning the correct version instead. The corrected artist and track name will be returned in the response.

user (Optional) : If called in non-authenticated mode you must specify the user to look up

api_key (Required) : A Last.fm API key.
*/
func (c *Client) TrackGetTags(artist, track, user string, opts ...RequestOption) (TrackTags, error) {
	// http://ws.audioscrobbler.com/2.0/?method=track.getTags&api_key=YOUR_API_KEY&artist=AC/DC&track=Hells+Bells&user=RJ&format=json
	lastfmURL := fmt.Sprintf("%s&method=track.getTags&user=%s", c.baseApiURL, user)

	opts = append(opts, ArtistOpt(artist), TrackOpt(track))

	values := processOptions(opts...).urlParams

	if query := values.Encode(); query != "" {
		lastfmURL += "&" + query
	}

	fmt.Println(lastfmURL)

	var trackTags struct {
		Tags TrackTags `json:"tags"`
	}

	err := c.get(lastfmURL, &trackTags)

	if err != nil {
		fmt.Println(err.Error())
		return TrackTags{}, err
	}

	return trackTags.Tags, nil
}

/*
track (Required (unless mbid)] : The track name

artist (Required (unless mbid)] : The artist name

mbid (Optional) : The musicbrainz id for the track

autocorrect[0|1] (Optional) : Transform misspelled artist and track names into correct artist and track names, returning the correct version instead. The corrected artist and track name will be returned in the response.

api_key (Required) : A Last.fm API key.
*/
func (c *Client) TrackGetTopTags(artist, track, user string, opts ...RequestOption) (TrackTopTags, error) {
	// http://ws.audioscrobbler.com/2.0/?method=track.gettoptags&artist=radiohead&track=paranoid+android&api_key=YOUR_API_KEY&format=json
	lastfmURL := fmt.Sprintf("%s&method=track.gettoptags&user=%s", c.baseApiURL, user)

	opts = append(opts, ArtistOpt(artist), TrackOpt(track))

	values := processOptions(opts...).urlParams

	if query := values.Encode(); query != "" {
		lastfmURL += "&" + query
	}

	var trackTags struct {
		Tags TrackTopTags `json:"toptags"`
	}

	err := c.get(lastfmURL, &trackTags)

	if err != nil {
		fmt.Println(err.Error())
		return TrackTopTags{}, err
	}

	return trackTags.Tags, nil
}

/*
limit (Optional) : The number of results to fetch per page. Defaults to 30.

page (Optional) : The page number to fetch. Defaults to first page.

track (Required) : The track name

artist (Optional) : Narrow your search by specifying an artist.

api_key (Required) : A Last.fm API key.
*/
func (c *Client) TrackSearch(track string, opts ...RequestOption) (TrackSearchRes, error) {
	// http://ws.audioscrobbler.com/2.0/?method=track.search&track=Believe&api_key=YOUR_API_KEY&format=json
	lastfmURL := fmt.Sprintf("%s&method=track.search", c.baseApiURL)

	opts = append(opts, TrackOpt(track))

	values := processOptions(opts...).urlParams

	if query := values.Encode(); query != "" {
		lastfmURL += "&" + query
	}

	var searchRes struct {
		SearchResults TrackSearchRes `json:"results"`
	}

	err := c.get(lastfmURL, &searchRes)

	if err != nil {
		return TrackSearchRes{}, err
	}

	return searchRes.SearchResults, nil
}
