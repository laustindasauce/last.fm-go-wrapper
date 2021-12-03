package lastfm

import (
	"errors"
	"fmt"

	"github.com/biter777/countries"
)

// GeoTopArtists ...
type GeoTopArtists struct {
	Artist []Artist `json:"artist"`
	Attr   GeoAttr  `json:"@attr"`
}

// GeoTopTracks ...
type GeoTopTracks struct {
	Track []GeoTrack `json:"track"`
	Attr  GeoAttr    `json:"@attr"`
}

// GeoTrack ...
type GeoTrack struct {
	Name       string          `json:"name"`
	Duration   string          `json:"duration"`
	Listeners  string          `json:"listeners"`
	MBID       string          `json:"mbid"`
	URL        string          `json:"url"`
	Streamable StreamableTrack `json:"streamable"`
	Artist     AlbumArtist     `json:"artist"`
	Image      []Image         `json:"image"`
	Attr       TrackRank       `json:"@attr"`
}

// GeoAttr ...
type GeoAttr struct {
	Country    string `json:"country"`
	Page       string `json:"page"`
	PerPage    string `json:"perPage"`
	TotalPages string `json:"totalPages"`
	Total      string `json:"total"`
}

/*
country (Required) : A country name, as defined by the ISO 3166-1 country names standard

limit (Optional) : The number of results to fetch per page. Defaults to 50.

page (Optional) : The page number to fetch. Defaults to first page.

api_key (Required) : A Last.fm API key.
*/
func (c *Client) GeoGetTopArtists(country string, opts ...RequestOption) (GeoTopArtists, error) {
	// http://ws.audioscrobbler.com/2.0/?method=geo.gettopartists&country=spain&api_key=YOUR_API_KEY&format=json

	// Check if the country is defined by the ISO 3166-1 country names standard
	thisCountry := countries.ByName(country)

	if thisCountry.String() == "Unknown" {
		return GeoTopArtists{}, errors.New("country param invalid")
	}

	lastfmURL := fmt.Sprintf("%s&method=geo.gettopartists", c.baseApiURL)

	opts = append(opts, CountryOpt(thisCountry.String()))

	values := processOptions(opts...).urlParams

	if query := values.Encode(); query != "" {
		lastfmURL += "&" + query
	}

	var topArtistsRes struct {
		TopArtists GeoTopArtists `json:"topartists"`
	}

	err := c.get(lastfmURL, &topArtistsRes)

	if err != nil {
		return GeoTopArtists{}, err
	}

	return topArtistsRes.TopArtists, nil
}

/*
country (Required) : A country name, as defined by the ISO 3166-1 country names standard

location (Optional) : A metro name, to fetch the charts for (must be within the country specified)

limit (Optional) : The number of results to fetch per page. Defaults to 50.

page (Optional) : The page number to fetch. Defaults to first page.

api_key (Required) : A Last.fm API key.
*/
func (c *Client) GeoGetTopTracks(country string, opts ...RequestOption) (GeoTopTracks, error) {
	// http://ws.audioscrobbler.com/2.0/?method=geo.gettoptracks&country=spain&api_key=YOUR_API_KEY&format=json
	thisCountry := countries.ByName(country)

	if thisCountry.String() == "Unknown" {
		return GeoTopTracks{}, errors.New("country param invalid")
	}

	lastfmURL := fmt.Sprintf("%s&method=geo.gettoptracks", c.baseApiURL)

	opts = append(opts, CountryOpt(thisCountry.String()))

	values := processOptions(opts...).urlParams

	if query := values.Encode(); query != "" {
		lastfmURL += "&" + query
	}

	var topTrackRes struct {
		TopTracks GeoTopTracks `json:"tracks"`
	}

	err := c.get(lastfmURL, &topTrackRes)

	if err != nil {
		return GeoTopTracks{}, err
	}

	return topTrackRes.TopTracks, nil
}
