package lastfm

type Artists struct {
	Artist []BareArtist `json:"artist"`
}

// Artist Stats identifies stats associated with an artist
type ArtistStats struct {
	Listeners string `json:"listeners"`
	Playcount string `json:"playcount"`
}

type ArtistAlbums struct {
	Album []Album `json:"album"`
	Attr ArtistAttr `json:"@attr"`
}

type AlbumArtist struct {
	Name string `json:"name"`
	MBID string `json:"mbid"`
	URL string 	`json:"url"`
}

type BareArtist struct {
	Name string `json:"name"`
	URL string 	`json:"url"`
	Image []Image `json:"image"`
}

type Artist struct {
	Name string `json:"name"`
	Listeners string `json:"listeners"`
	MBID string `json:"mbid"`
	Match string `json:"match"`
	URL string 	`json:"url"`
	Image []Image `json:"image"`
	Streamable string `json:"streamable"`
}

type FullArtist struct {
	Name string `json:"name"`
	MBID string `json:"mbid"`
	URL string 	`json:"url"`
	Image []Image `json:"image"`
	Streamable string `json:"streamable"`
	OnTour string `json:"on_tour"`
	Stats ArtistStats `json:"stats"`
	Similar Artists `json:"similar"`
	Tags Tags `json:"tags"`
	Bio Bio `json:"bio"`
}

type SimilarArtists struct {
	SimilarArtists []Artist `json:"artist"`
	Attr ArtistTagAttr `json:"@attr"`
}

type ArtistAttr struct {
	Artist string `json:"artist"`
	Page string `json:"page"`
	PerPage string `json:"perPage"`
	TotalPages string `json:"totalPages"`
	Total string `json:"total"`
}

type ArtistTagAttr struct {
	Artist string `json:"artist"`
}

type ArtistSearchRes struct {
	Query OpenSearchQuery `json:"opensearch:Query"`
	QueryTotalResults string `json:"opensearch:totalResults"`
	QueryStartIndex string `json:"opensearch:startIndex"`
	QueryItemsPerPage string `json:"opensearch:itemsPerPage"`
	ArtistMatches ArtistMatches `json:"artistmatches"`
	Attr SearchAttr `json:"@attr"`
}

type ArtistMatches struct {
	Artist []Artist `json:"artist"`
}

// ArtistGetInfo uses the artist.getinfo method with scrobbler API to return the specifed
// full artist information
func (c *Client) ArtistGetInfo(artist string) (*FullArtist, error) {
	lastfmURL := c.getNoAuthURL("method.artist.getinfo", "artist."+artist)

	var artistInfo struct {
		Artist FullArtist `json:"artist"`
	}

	err := c.get(lastfmURL, &artistInfo)

	if err != nil {
		return nil, err
	}

	return &artistInfo.Artist, nil
}

func (c *Client) ArtistGetSimilar(artist string) (SimilarArtists, error) {
	lastfmURL := c.getNoAuthURL("method.artist.getsimilar", "artist."+artist)

	var similar struct {
		SimilarRes SimilarArtists `json:"similarartists"`
	}

	err := c.get(lastfmURL, &similar)

	if err != nil {
		return SimilarArtists{}, err
	}

	return similar.SimilarRes, nil
}

func (c *Client) ArtistGetTopAlbums(artist string) (ArtistAlbums, error) {
	// method=artist.gettopalbums&artist=cher&api_key=YOUR_API_KEY&format=json
	lastfmURL := c.getNoAuthURL("method.artist.gettopalbums", "artist."+artist)

	var topAlbums struct {
		Albums ArtistAlbums `json:"topalbums"`
	}

	err := c.get(lastfmURL, &topAlbums)

	if err != nil {
		return ArtistAlbums{}, err
	}

	return topAlbums.Albums, nil
}

func (c *Client) ArtistGetTopTags(artist string) (TagsWithCount, error) {
	//http://ws.audioscrobbler.com/2.0/?method=artist.gettoptags&artist=cher&api_key=YOUR_API_KEY&format=json
	lastfmURL := c.getNoAuthURL("method.artist.gettoptags", "artist."+artist)
	
	var topTags struct {
		Tags TagsWithCount `json:"toptags"`
	}

	err := c.get(lastfmURL, &topTags)

	if err != nil {
		return TagsWithCount{}, err
	}

	return topTags.Tags, nil
}


func (c *Client) ArtistGetTopTracks(artist string) (Tracks, error) {
	//http://ws.audioscrobbler.com/2.0/?method=artist.gettoptracks&artist=cher&api_key=YOUR_API_KEY&format=json
	lastfmURL := c.getNoAuthURL("method.artist.gettoptracks", "artist."+artist)


	var topTracks struct {
		TopTracks Tracks `json:"toptracks"`
	}

	err := c.get(lastfmURL, &topTracks)

	if err != nil {
		return Tracks{}, err
	}
	
	return topTracks.TopTracks, nil
}

func (c *Client) ArtistSearch(artist string) (ArtistSearchRes, error) {
	//http://ws.audioscrobbler.com/2.0/?method=artist.search&artist=cher&api_key=YOUR_API_KEY&format=json
	lastfmURL := c.getNoAuthURL("method.artist.search", "artist."+artist)


	var searchRes struct {
		SearchResults ArtistSearchRes `json:"results"`
	}

	err := c.get(lastfmURL, &searchRes)

	if err != nil {
		return ArtistSearchRes{}, err
	}
	
	return searchRes.SearchResults, nil
}