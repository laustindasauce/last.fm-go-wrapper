package lastfm

import "fmt"

// LibraryArtists ...
type LibraryArtists struct {
	Artists []LibraryArtist `json:"artist"`
	Attr    UserAttr        `json:"@attr"`
}

// LibraryArtist ...
type LibraryArtist struct {
	TagCount   string  `json:"tagcount"`
	Image      []Image `json:"image"`
	MBID       string  `json:"mbid"`
	URL        string  `json:"url"`
	Playcount  string  `json:"playcount"`
	Name       string  `json:"name"`
	Streamable string  `json:"streamable"`
}

/*
user (Required) : The user whose library you want to fetch.

limit (Optional) : The number of results to fetch per page. Defaults to 50.

page (Optional) : The page number you wish to scan to.

api_key (Required) : A Last.fm API key.
*/
func (c *Client) LibraryGetArtists(user string, opts ...RequestOption) (LibraryArtists, error) {
	// http://ws.audioscrobbler.com/2.0/?method=library.getartists&api_key=YOUR_API_KEY&user=joanofarctan&format=json
	lastfmURL := fmt.Sprintf("%s&method=library.getartists&user=%s", c.baseApiURL, user)

	values := processOptions(opts...).urlParams

	if query := values.Encode(); query != "" {
		lastfmURL += "&" + query
	}

	var artists struct {
		Artists LibraryArtists `json:"artists"`
	}

	err := c.get(lastfmURL, &artists)

	if err != nil {
		return LibraryArtists{}, err
	}

	return artists.Artists, nil
}
