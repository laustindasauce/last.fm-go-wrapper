package lastfm

import (
	"net/url"
	"strconv"
)

type RequestOption func(*requestOptions)

type requestOptions struct {
	urlParams url.Values
}

// LimitOpt – The number of results to fetch per page. Defaults to 50. Maximum is 200.
func LimitOpt(amount int) RequestOption {
	return func(o *requestOptions) {
		o.urlParams.Set("limit", strconv.Itoa(amount))
	}
}

// PageOpt – The page number to fetch. Defaults to first page.
func PageOpt(page int) RequestOption {
	return func(o *requestOptions) {
		o.urlParams.Set("page", strconv.Itoa(page))
	}
}

// ArtistOpt – The artist name
func ArtistOpt(artist string) RequestOption {
	return func(o *requestOptions) {
		o.urlParams.Set("artist", artist)
	}
}

// AlbumOpt – The album name
func AlbumOpt(album string) RequestOption {
	return func(o *requestOptions) {
		o.urlParams.Set("album", album)
	}
}

// TrackOpt – The artist name
func TrackOpt(track string) RequestOption {
	return func(o *requestOptions) {
		o.urlParams.Set("track", track)
	}
}

// MbidOpt – The musicbrainz id for the album
func MbidOpt(id string) RequestOption {
	return func(o *requestOptions) {
		o.urlParams.Set("mbid", id)
	}
}

// Username – The username for the context of the request. If supplied, the user's playcount for this album is included in the response.
func UsernameOpt(user string) RequestOption {
	return func(o *requestOptions) {
		o.urlParams.Set("username", user)
	}
}

// Lang – The language to return the biography in, expressed as an ISO 639 alpha-2 code.
func LangOpt(code string) RequestOption {
	return func(o *requestOptions) {
		o.urlParams.Set("username", code)
	}
}

// CountryOpt – A country name, as defined by the ISO 3166-1 country names standard
func CountryOpt(code string) RequestOption {
	return func(o *requestOptions) {
		o.urlParams.Set("country", code)
	}
}

// LocationOpt – A metro name, to fetch the charts for (must be within the country specified)
func LocationOpt(local string) RequestOption {
	return func(o *requestOptions) {
		o.urlParams.Set("location", local)
	}
}

// TagOpt – The tag name
func TagOpt(tag string) RequestOption {
	return func(o *requestOptions) {
		o.urlParams.Set("tag", tag)
	}
}

// FromOpt – The date at which the chart should start from. See User.getChartsList for more.
// integer number of seconds since 00:00:00, January 1st 1970 UTC). This must be in the UTC time zone.
func FromOpt(from string) RequestOption {
	return func(o *requestOptions) {
		o.urlParams.Set("from", from)
	}
}

// ToOpt – End timestamp of a range - only display scrobbles before this time, in UNIX timestamp format
// integer number of seconds since 00:00:00, January 1st 1970 UTC). This must be in the UTC time zone.
func ToOpt(to string) RequestOption {
	return func(o *requestOptions) {
		o.urlParams.Set("to", to)
	}
}

type Binary string

const (
	// 0 means do not use this opt
	Zero Binary = "0"
	// 1 means do use this opt
	One Binary = "1"
)

// AutocorrectOpt – Transform misspelled artist names into correct artist names, returning the correct version instead. The corrected artist name will be returned in the response.
func AutocorrectOpt(autocorrect Binary) RequestOption {
	return func(o *requestOptions) {
		o.urlParams.Set("autocorrect", string(autocorrect))
	}
}

// ExtendedOpt – Includes extended data in each artist, and whether or not the user has loved each track
func ExtendedOpt(extend Binary) RequestOption {
	return func(o *requestOptions) {
		o.urlParams.Set("extended", string(extend))
	}
}

type Tagging string

const (
	AlbumTag  Tagging = "album"
	ArtistTag Tagging = "artist"
	TrackTag  Tagging = "track"
)

// TaggingTypeOpt – The type of items which have been tagged
func TaggingTypeOpt(tag Tagging) RequestOption {
	return func(o *requestOptions) {
		o.urlParams.Set("taggingtype", string(tag))
	}
}

func processOptions(options ...RequestOption) requestOptions {
	o := requestOptions{
		urlParams: url.Values{},
	}
	for _, opt := range options {
		opt(&o)
	}

	return o
}
