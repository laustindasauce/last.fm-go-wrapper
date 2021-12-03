package lastfm

import "fmt"

type TopArtists struct {
	Artist []ChartArtist `json:"artist"`
	Attr   ChartAttr     `json:"@attr"`
}

type ChartArtist struct {
	Name       string  `json:"name"`
	Playcount  string  `json:"playcount"`
	Listeners  string  `json:"listeners"`
	MBID       string  `json:"mbid"`
	URL        string  `json:"url"`
	Streamable string  `json:"streamable"`
	Image      []Image `json:"image"`
}

type TopTags struct {
	Tag  []ChartTag `json:"tag"`
	Attr ChartAttr  `json:"@attr"`
}

type ChartTag struct {
	Name       string `json:"name"`
	URL        string `json:"url"`
	Reach      string `json:"reach"`
	Taggings   string `json:"taggings"`
	Streamable string `json:"streamable"`
	Wiki       Wiki   `json:"wiki"`
}

type ChartTracks struct {
	Track []ChartTrack `json:"track"`
	Attr  ChartAttr    `json:"@attr"`
}

type ChartTrack struct {
	Name       string          `json:"name"`
	Duration   string          `json:"duration"`
	Playcount  string          `json:"playcount"`
	Listeners  string          `json:"listeners"`
	MBID       string          `json:"mbid"`
	URL        string          `json:"url"`
	Streamable StreamableTrack `json:"streamable"`
	Artist     AlbumArtist     `json:"artist"`
	Image      []Image         `json:"image"`
}

type ChartAttr struct {
	Page       string `json:"page"`
	PerPage    string `json:"perPage"`
	TotalPages string `json:"totalPages"`
	Total      string `json:"total"`
}

/*
page (Optional) : The page number to fetch. Defaults to first page.

limit (Optional) : The number of results to fetch per page. Defaults to 50.

api_key (Required) : A Last.fm API key.
*/
func (c *Client) ChartGetTopArtists(opts ...RequestOption) (TopArtists, error) {
	// http://ws.audioscrobbler.com/2.0/?method=chart.gettopartists&api_key=YOUR_API_KEY&format=json
	lastfmURL := fmt.Sprintf("%s&method=chart.gettopartists", c.baseApiURL)

	values := processOptions(opts...).urlParams

	if query := values.Encode(); query != "" {
		lastfmURL += "&" + query
	}

	var topArtistsRes struct {
		TopArtists TopArtists `json:"artists"`
	}

	err := c.get(lastfmURL, &topArtistsRes)

	if err != nil {
		return TopArtists{}, err
	}

	return topArtistsRes.TopArtists, nil
}

/*
page (Optional) : The page number to fetch. Defaults to first page.

limit (Optional) : The number of results to fetch per page. Defaults to 50.

api_key (Required) : A Last.fm API key.
*/
func (c *Client) ChartGetTopTags(opts ...RequestOption) (TopTags, error) {
	// http://ws.audioscrobbler.com/2.0/?method=chart.gettoptags&api_key=YOUR_API_KEY&format=json
	lastfmURL := fmt.Sprintf("%s&method=chart.gettoptags", c.baseApiURL)

	values := processOptions(opts...).urlParams

	if query := values.Encode(); query != "" {
		lastfmURL += "&" + query
	}

	var topTagsRes struct {
		TopTags TopTags `json:"tags"`
	}

	err := c.get(lastfmURL, &topTagsRes)

	if err != nil {
		return TopTags{}, err
	}

	return topTagsRes.TopTags, nil
}

/*
page (Optional) : The page number to fetch. Defaults to first page.

limit (Optional) : The number of results to fetch per page. Defaults to 50.

api_key (Required) : A Last.fm API key.
*/
func (c *Client) ChartGetTopTracks(opts ...RequestOption) (ChartTracks, error) {
	// http://ws.audioscrobbler.com/2.0/?method=chart.gettoptracks&api_key=YOUR_API_KEY&format=json
	lastfmURL := fmt.Sprintf("%s&method=chart.gettoptracks", c.baseApiURL)

	values := processOptions(opts...).urlParams

	if query := values.Encode(); query != "" {
		lastfmURL += "&" + query
	}

	var topTracksRes struct {
		TopTracks ChartTracks `json:"tracks"`
	}

	err := c.get(lastfmURL, &topTracksRes)

	if err != nil {
		return ChartTracks{}, err
	}

	return topTracksRes.TopTracks, nil
}
