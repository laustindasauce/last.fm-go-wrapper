package lastfm

import "fmt"

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

func (c *Client) LibraryGetArtists() (LibraryArtists, error) {
	// http://ws.audioscrobbler.com/2.0/?method=library.getartists&api_key=YOUR_API_KEY&user=joanofarctan&format=json
	lastfmURL := c.getNoAuthURL("method.library.getartists", "user."+c.User)
	fmt.Println(lastfmURL)

	var artists struct {
		Artists LibraryArtists `json:"artists"`
	}

	err := c.get(lastfmURL, &artists)

	if err != nil {
		return LibraryArtists{}, err
	}

	return artists.Artists, nil
}
