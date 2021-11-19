package lastfm

import "errors"

type LibraryArtists struct {
	Artists []LibraryArtist `json:"artist"`
	Attr    UserAttr        `json:"@attr"`
}

type LibraryArtist struct {
	TagCount   string  `json:"tagcount"`
	Image      []Image `json:"image"`
	MBID       string  `json:"mbid"`
	URL        string  `json:"url"`
	Playcount  string  `json:"playcount"`
	Name       string  `json:"name"`
	Streamable string  `json:"streamable"`
}

func (c *Client) LibraryGetArtists(limit, page string) (LibraryArtists, error) {
	// http://ws.audioscrobbler.com/2.0/?method=library.getartists&api_key=YOUR_API_KEY&user=joanofarctan&format=json
	if c.User == "" {
		return LibraryArtists{}, errors.New("empty user... please run set user method first")
	}

	allOpts := []string{"method.library.getartists", "user." + c.User, "limit." + limit, "page." + page}

	lastfmURL := c.getNoAuthURL(allOpts...)

	var artists struct {
		Artists LibraryArtists `json:"artists"`
	}

	err := c.get(lastfmURL, &artists)

	if err != nil {
		return LibraryArtists{}, err
	}

	return artists.Artists, nil
}
