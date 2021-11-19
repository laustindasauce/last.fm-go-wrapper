package lastfm

import (
	"errors"

	"github.com/biter777/countries"
)

type GeoTopArtists struct {
	Artist []Artist `json:"artist"`
	Attr   GeoAttr  `json:"@attr"`
}

type GeoTopTracks struct {
	Track []GeoTrack `json:"track"`
	Attr  GeoAttr    `json:"@attr"`
}

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

type GeoAttr struct {
	Country    string `json:"country"`
	Page       string `json:"page"`
	PerPage    string `json:"perPage"`
	TotalPages string `json:"totalPages"`
	Total      string `json:"total"`
}

func (c *Client) GeoGetTopArtists(country, limit, page string) (GeoTopArtists, error) {
	// http://ws.audioscrobbler.com/2.0/?method=chart.gettoptracks&api_key=YOUR_API_KEY&format=json

	// Check if the country is defined by the ISO 3166-1 country names standard
	thisCountry := countries.ByName(country)

	if thisCountry.String() == "Unknown" {
		return GeoTopArtists{}, errors.New("country param invalid")
	}

	lastfmURL := c.getNoAuthURL("method.geo.gettopartists", "country."+thisCountry.String(), "limit."+limit, "page."+page)

	var topArtistsRes struct {
		TopArtists GeoTopArtists `json:"topartists"`
	}

	err := c.get(lastfmURL, &topArtistsRes)

	if err != nil {
		return GeoTopArtists{}, err
	}

	return topArtistsRes.TopArtists, nil
}

func (c *Client) GeoGetTopTracks(country, location, limit, page string) (GeoTopTracks, error) {
	// http://ws.audioscrobbler.com/2.0/?method=geo.gettoptracks&country=spain&api_key=YOUR_API_KEY&format=json
	thisCountry := countries.ByName(country)

	if thisCountry.String() == "Unknown" {
		return GeoTopTracks{}, errors.New("country param invalid")
	}

	lastfmURL := c.getNoAuthURL("method.geo.gettoptracks", "country."+thisCountry.String(), "location."+location, "limit."+limit, "page."+page)

	var topTrackRes struct {
		TopTracks GeoTopTracks `json:"tracks"`
	}

	err := c.get(lastfmURL, &topTrackRes)

	if err != nil {
		return GeoTopTracks{}, err
	}

	return topTrackRes.TopTracks, nil
}
