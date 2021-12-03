package lastfm

import "fmt"

// Artists ...
type Artists struct {
	Artist []BareArtist `json:"artist"`
}

// Artist Stats identifies stats associated with an artist
type ArtistStats struct {
	Listeners string `json:"listeners"`
	Playcount string `json:"playcount"`
}

// ArtistAlbums ...
type ArtistAlbums struct {
	Album []Album    `json:"album"`
	Attr  ArtistAttr `json:"@attr"`
}

// AlbumArtist ...
type AlbumArtist struct {
	Name string `json:"name"`
	MBID string `json:"mbid"`
	URL  string `json:"url"`
}

// BareArtist ...
type BareArtist struct {
	Name  string  `json:"name"`
	URL   string  `json:"url"`
	Image []Image `json:"image"`
}

// Artist ...
type Artist struct {
	Name       string  `json:"name"`
	Listeners  string  `json:"listeners"`
	MBID       string  `json:"mbid"`
	Match      string  `json:"match,omitempty"`
	URL        string  `json:"url"`
	Image      []Image `json:"image"`
	Streamable string  `json:"streamable"`
}

// FullArtist ...
type FullArtist struct {
	Name       string      `json:"name"`
	MBID       string      `json:"mbid"`
	URL        string      `json:"url"`
	Image      []Image     `json:"image"`
	Streamable string      `json:"streamable"`
	OnTour     string      `json:"on_tour"`
	Stats      ArtistStats `json:"stats"`
	Similar    Artists     `json:"similar"`
	Tags       Tags        `json:"tags"`
	Bio        Bio         `json:"bio"`
}

// ArtistRanked ...
type ArtistRanked struct {
	Streamable string  `json:"streamable"`
	Image      []Image `json:"image"`
	MBID       string  `json:"mbid"`
	URL        string  `json:"url"`
	Playcount  string  `json:"playcount"`
	Attr       Rank    `json:"@attr"`
	Name       string  `json:"name"`
}

// WeeklyArtistChart ...
type WeeklyArtistChart struct {
	MBID      string `json:"mbid"`
	URL       string `json:"url"`
	Name      string `json:"name"`
	Attr      Rank   `json:"@attr"`
	Playcount string `json:"playcount"`
}

// ArtistRecent ...
type ArtistRecent struct {
	MBID  string  `json:"mbid"`
	Text  string  `json:"#text,omitempty"`
	URL   string  `json:"url,omitempty"`
	Name  string  `json:"name,omitempty"`
	Image []Image `json:"image,omitempty"`
}

// SimilarArtists ...
type SimilarArtists struct {
	SimilarArtists []Artist      `json:"artist"`
	Attr           ArtistTagAttr `json:"@attr"`
}

// ArtistAttr ...
type ArtistAttr struct {
	Artist     string `json:"artist"`
	Page       string `json:"page"`
	PerPage    string `json:"perPage"`
	TotalPages string `json:"totalPages"`
	Total      string `json:"total"`
}

// ArtistTagAttr ...
type ArtistTagAttr struct {
	Artist string `json:"artist"`
}

// ArtistSearchRes ...
type ArtistSearchRes struct {
	Query             OpenSearchQuery `json:"opensearch:Query"`
	QueryTotalResults string          `json:"opensearch:totalResults"`
	QueryStartIndex   string          `json:"opensearch:startIndex"`
	QueryItemsPerPage string          `json:"opensearch:itemsPerPage"`
	ArtistMatches     ArtistMatches   `json:"artistmatches"`
	Attr              SearchAttr      `json:"@attr"`
}

// ArtistMatches ...
type ArtistMatches struct {
	Artist []Artist `json:"artist"`
}

/*
artist (Required (unless mbid)] : The artist name

mbid (Optional) : The musicbrainz id for the artist

lang (Optional) : The language to return the biography in, expressed as an ISO 639 alpha-2 code.

autocorrect[0|1] (Optional) : Transform misspelled artist names into correct artist names, returning the correct version instead. The corrected artist name will be returned in the response.

username (Optional) : The username for the context of the request. If supplied, the user's playcount for this artist is included in the response.

api_key (Required) : A Last.fm API key.
*/
func (c *Client) ArtistGetInfo(artist string, opts ...RequestOption) (FullArtist, error) {
	// http://ws.audioscrobbler.com/2.0/?method=artist.getinfo&artist=Cher&api_key=YOUR_API_KEY&format=json
	lastfmURL := fmt.Sprintf("%s&method=artist.getinfo", c.baseApiURL)

	opts = append(opts, ArtistOpt(artist))

	values := processOptions(opts...).urlParams

	if query := values.Encode(); query != "" {
		lastfmURL += "&" + query
	}

	var artistInfo struct {
		Artist FullArtist `json:"artist"`
	}

	err := c.get(lastfmURL, &artistInfo)

	if err != nil {
		return FullArtist{}, err
	}

	return artistInfo.Artist, nil
}

/*
limit (Optional) : Limit the number of similar artists returned

artist (Required (unless mbid)] : The artist name

autocorrect[0|1] (Optional) : Transform misspelled artist names into correct artist names, returning the correct version instead. The corrected artist name will be returned in the response.

mbid (Optional) : The musicbrainz id for the artist

api_key (Required) : A Last.fm API key.
*/
func (c *Client) ArtistGetSimilar(artist string, opts ...RequestOption) (SimilarArtists, error) {
	// http://ws.audioscrobbler.com/2.0/?method=artist.getsimilar&artist=cher&api_key=YOUR_API_KEY&format=json
	lastfmURL := fmt.Sprintf("%s&method=artist.getsimilar", c.baseApiURL)

	opts = append(opts, ArtistOpt(artist))

	values := processOptions(opts...).urlParams

	if query := values.Encode(); query != "" {
		lastfmURL += "&" + query
	}

	var similar struct {
		SimilarRes SimilarArtists `json:"similarartists"`
	}

	err := c.get(lastfmURL, &similar)

	if err != nil {
		return SimilarArtists{}, err
	}

	return similar.SimilarRes, nil
}

/*
artist (Required (unless mbid)] : The artist name

mbid (Optional) : The musicbrainz id for the artist

autocorrect[0|1] (Optional) : Transform misspelled artist names into correct artist names, returning the correct version instead. The corrected artist name will be returned in the response.

page (Optional) : The page number to fetch. Defaults to first page.

limit (Optional) : The number of results to fetch per page. Defaults to 50.

api_key (Required) : A Last.fm API key.
*/
func (c *Client) ArtistGetTopAlbums(artist string, opts ...RequestOption) (ArtistAlbums, error) {
	// method=artist.gettopalbums&artist=cher&api_key=YOUR_API_KEY&format=json
	lastfmURL := fmt.Sprintf("%s&method=artist.gettopalbums", c.baseApiURL)

	opts = append(opts, ArtistOpt(artist))

	values := processOptions(opts...).urlParams

	if query := values.Encode(); query != "" {
		lastfmURL += "&" + query
	}

	var topAlbums struct {
		Albums ArtistAlbums `json:"topalbums"`
	}

	err := c.get(lastfmURL, &topAlbums)

	if err != nil {
		return ArtistAlbums{}, err
	}

	return topAlbums.Albums, nil
}

/*
artist (Required (unless mbid)] : The artist name

mbid (Optional) : The musicbrainz id for the artist

autocorrect[0|1] (Optional) : Transform misspelled artist names into correct artist names, returning the correct version instead. The corrected artist name will be returned in the response.

api_key (Required) : A Last.fm API key.
*/
func (c *Client) ArtistGetTopTags(artist string, opts ...RequestOption) (TagsWithCount, error) {
	//http://ws.audioscrobbler.com/2.0/?method=artist.gettoptags&artist=cher&api_key=YOUR_API_KEY&format=json
	lastfmURL := fmt.Sprintf("%s&method=artist.gettoptags", c.baseApiURL)

	opts = append(opts, ArtistOpt(artist))

	values := processOptions(opts...).urlParams

	if query := values.Encode(); query != "" {
		lastfmURL += "&" + query
	}

	var topTags struct {
		Tags TagsWithCount `json:"toptags"`
	}

	err := c.get(lastfmURL, &topTags)

	if err != nil {
		return TagsWithCount{}, err
	}

	return topTags.Tags, nil
}

/*
artist (Required (unless mbid)] : The artist name

mbid (Optional) : The musicbrainz id for the artist

autocorrect[0|1] (Optional) : Transform misspelled artist names into correct artist names, returning the correct version instead. The corrected artist name will be returned in the response.

page (Optional) : The page number to fetch. Defaults to first page.

limit (Optional) : The number of results to fetch per page. Defaults to 50.

api_key (Required) : A Last.fm API key.
*/
func (c *Client) ArtistGetTopTracks(artist string, opts ...RequestOption) (Tracks, error) {
	//http://ws.audioscrobbler.com/2.0/?method=artist.gettoptracks&artist=cher&api_key=YOUR_API_KEY&format=json
	lastfmURL := fmt.Sprintf("%s&method=artist.gettoptracks", c.baseApiURL)

	opts = append(opts, ArtistOpt(artist))

	values := processOptions(opts...).urlParams

	if query := values.Encode(); query != "" {
		lastfmURL += "&" + query
	}

	var topTracks struct {
		TopTracks Tracks `json:"toptracks"`
	}

	err := c.get(lastfmURL, &topTracks)

	if err != nil {
		return Tracks{}, err
	}

	return topTracks.TopTracks, nil
}

/*
limit (Optional) : The number of results to fetch per page. Defaults to 30.

page (Optional) : The page number to fetch. Defaults to first page.

artist (Required) : The artist name

api_key (Required) : A Last.fm API key.
*/
func (c *Client) ArtistSearch(artist string, opts ...RequestOption) (ArtistSearchRes, error) {
	//http://ws.audioscrobbler.com/2.0/?method=artist.search&artist=cher&api_key=YOUR_API_KEY&format=json
	lastfmURL := fmt.Sprintf("%s&method=artist.search", c.baseApiURL)

	opts = append(opts, ArtistOpt(artist))

	values := processOptions(opts...).urlParams

	if query := values.Encode(); query != "" {
		lastfmURL += "&" + query
	}

	var searchRes struct {
		SearchResults ArtistSearchRes `json:"results"`
	}

	err := c.get(lastfmURL, &searchRes)

	if err != nil {
		return ArtistSearchRes{}, err
	}

	return searchRes.SearchResults, nil
}
