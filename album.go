package lastfm

// Album struct
type Album struct {
	Name string `json:"name"`
	Playcount int `json:"playcount"`
	MBID string `json:"mbid"`
	URL string `json:"url"`
	Artist AlbumArtist `json:"artist"`
	Image []Image `json:"image"`
	Attr ArtistAttr `json:"@attr"`
}

type SimpleAlbum struct {
	Name string `json:"name"`
	Artist string `json:"artist"`
	URL string `json:"url"`
	Image []Image `json:"image"`
	Streamable string `json:"streamable"`
	MBID string `json:"mbid"`
}

type FullAlbum struct {
	Artist string `json:"artist"`
	MBID string `json:"mbid"`
	Tags Tags `json:"tags"`
	Playcount string `json:"playcount"`
	Image []Image `json:"image"`
	Tracks AlbumTracks `json:"tracks"`
	URL string `json:"url"`
	Name string `json:"name"`
	Listeners string `json:"listeners"`
	Wiki Wiki `json:"wiki"`
}

type AlbumTracks struct {
	Track []TrackAlbum `json:"track"`
}

type AlbumTopTags struct {
	Tag []TagWithCount `json:"tag"`
	Attr AlbumAttr `json:"@attr"`
}

type AlbumAttr struct {
	Artist string `json:"artist"`
	Album string `json:"album"`
}

type AlbumSearchRes struct {
	Query OpenSearchQuery `json:"opensearch:Query"`
	QueryTotalResults string `json:"opensearch:totalResults"`
	QueryStartIndex string `json:"opensearch:startIndex"`
	QueryItemsPerPage string `json:"opensearch:itemsPerPage"`
	AlbumMatches AlbumMatches `json:"albummatches"`
	Attr SearchAttr `json:"@attr"`
}

type AlbumMatches struct {
	Album []SimpleAlbum `json:"album"`
}

func (c *Client) AlbumGetInfo(album, artist string) (FullAlbum, error) {
	// http://ws.audioscrobbler.com/2.0/?method=album.getinfo&api_key=YOUR_API_KEY&artist=Cher&album=Believe&format=json
	lastfmURL := c.getNoAuthURL("method.album.getinfo", "album."+album, "artist."+artist)

	var albumInfo struct {
		FullAlbum FullAlbum `json:"album"`
	}

	err := c.get(lastfmURL, &albumInfo)

	if err != nil {
		return FullAlbum{}, err
	}

	return albumInfo.FullAlbum, nil
}

// func (c *Client) AlbumGetTags(album, artist string) (string, error) {
// 	// http://ws.audioscrobbler.com/2.0/?method=album.gettags&artist=cher&album=believe&api_key=YOUR_API_KEY&format=json
// 	lastfmURL := c.getNoAuthURL("method.album.gettags", "album."+album, "artist."+artist)
// 	return lastfmURL, nil
// }

func (c *Client) AlbumGetTopTags(album, artist string) (AlbumTopTags, error) {
	// http://ws.audioscrobbler.com/2.0/?method=album.gettoptags&artist=radiohead&album=the%20bends&api_key=YOUR_API_KEY&format=json
	lastfmURL := c.getNoAuthURL("method.album.gettoptags", "album."+album, "artist."+artist)

	var topTags struct {
		TopTags AlbumTopTags `json:"toptags"`
	}

	err := c.get(lastfmURL, &topTags)

	if err != nil {
		return AlbumTopTags{}, err
	}

	return topTags.TopTags, nil
}

func (c *Client) AlbumSearch(album string) (AlbumSearchRes, error) {
	// http://ws.audioscrobbler.com/2.0/?method=album.search&album=believe&api_key=YOUR_API_KEY&format=json
	lastfmURL := c.getNoAuthURL("method.album.search", "album."+album)

	var searchRes struct {
		SearchResults AlbumSearchRes `json:"results"`
	}

	err := c.get(lastfmURL, &searchRes)

	if err != nil {
		return AlbumSearchRes{}, err
	}
	return searchRes.SearchResults, nil
}