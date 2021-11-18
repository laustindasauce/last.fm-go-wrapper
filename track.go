package lastfm

import (
	"errors"
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

func (c *Client) TrackGetCorrection(track, artist string) (TrackCorrection, error) {
	// http://ws.audioscrobbler.com/2.0/?method=track.getcorrection&artist=guns and roses&track=Mrbrownstone&api_key=YOUR_API_KEY&format=json
	lastfmURL := c.getNoAuthURL("method.track.getcorrection", "track."+track, "artist."+artist)

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

func (c *Client) TrackGetInfo(track, artist string) (TrackInfo, error) {
	// http://ws.audioscrobbler.com/2.0/?method=track.getInfo&api_key=YOUR_API_KEY&artist=cher&track=believe&format=json
	lastfmURL := c.getNoAuthURL("method.track.getInfo", "track."+track, "artist."+artist)

	var trackInfo struct {
		Track TrackInfo `json:"track"`
	}

	err := c.get(lastfmURL, &trackInfo)

	if err != nil {
		return TrackInfo{}, err
	}

	return trackInfo.Track, nil
}

func (c *Client) TrackGetSimilar(track, artist, mbid, limit string) (SimilarTracks, error) {
	// http://ws.audioscrobbler.com/2.0/?method=track.getsimilar&artist=cher&track=believe&api_key=YOUR_API_KEY&format=json
	var lastfmURL string
	var allOpts []string

	allOpts = append(allOpts, "method.track.getsimilar")

	if mbid != "" {
		allOpts = append(allOpts, "mbid."+mbid)
	} else {
		allOpts = append(allOpts, "track."+track, "artist."+artist)
	}

	if limit != "" {
		allOpts = append(allOpts, "limit."+limit)
	}

	lastfmURL = c.getNoAuthURL(allOpts...)

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

func (c *Client) TrackGetTags(track, artist, mbid string) (TrackTags, error) {
	// http://ws.audioscrobbler.com/2.0/?method=track.getTags&api_key=YOUR_API_KEY&artist=AC/DC&track=Hells+Bells&user=RJ&format=json
	var lastfmURL string
	var allOpts []string

	allOpts = append(allOpts, "method.track.getTags")

	if c.User == "" {
		return TrackTags{}, errors.New("empty user... please run set user method first")
	} else {
		allOpts = append(allOpts, "user."+c.User)
	}

	if mbid != "" {
		allOpts = append(allOpts, "mbid."+mbid)
	} else {
		allOpts = append(allOpts, "track."+track, "artist."+artist)
	}

	lastfmURL = c.getNoAuthURL(allOpts...)

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

func (c *Client) TrackGetTopTags(track, artist, mbid string) (TrackTopTags, error) {
	// http://ws.audioscrobbler.com/2.0/?method=track.gettoptags&artist=radiohead&track=paranoid+android&api_key=YOUR_API_KEY&format=json
	var lastfmURL string
	var allOpts []string

	allOpts = append(allOpts, "method.track.gettoptags")

	if mbid != "" {
		allOpts = append(allOpts, "mbid."+mbid)
	} else {
		allOpts = append(allOpts, "track."+track, "artist."+artist)
	}

	lastfmURL = c.getNoAuthURL(allOpts...)

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

func (c *Client) TrackSearch(track, artist, page, limit string) (TrackSearchRes, error) {
	// http://ws.audioscrobbler.com/2.0/?method=track.search&track=Believe&api_key=YOUR_API_KEY&format=json
	var lastfmURL string
	var allOpts []string

	allOpts = append(allOpts, "method.track.search", "track."+track)

	if limit != "" {
		allOpts = append(allOpts, "limit."+limit)
	}

	if artist != "" {
		allOpts = append(allOpts, "artist."+artist)
	}

	if page != "" {
		allOpts = append(allOpts, "page."+page)
	}

	lastfmURL = c.getNoAuthURL(allOpts...)

	var searchRes struct {
		SearchResults TrackSearchRes `json:"results"`
	}

	err := c.get(lastfmURL, &searchRes)

	if err != nil {
		return TrackSearchRes{}, err
	}

	return searchRes.SearchResults, nil
}
