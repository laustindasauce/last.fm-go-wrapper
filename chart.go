package lastfm

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

func (c *Client) ChartGetTopArtists(limit, page string) (TopArtists, error) {
	// http://ws.audioscrobbler.com/2.0/?method=chart.gettopartists&api_key=YOUR_API_KEY&format=json
	allOpts := []string{"method.chart.gettopartists", "limit." + limit, "page." + page}

	lastfmURL := c.getNoAuthURL(allOpts...)

	var topArtistsRes struct {
		TopArtists TopArtists `json:"artists"`
	}

	err := c.get(lastfmURL, &topArtistsRes)

	if err != nil {
		return TopArtists{}, err
	}

	return topArtistsRes.TopArtists, nil
}

func (c *Client) ChartGetTopTags(limit, page string) (TopTags, error) {
	// http://ws.audioscrobbler.com/2.0/?method=chart.gettoptags&api_key=YOUR_API_KEY&format=json
	allOpts := []string{"method.chart.gettoptags", "limit." + limit, "page." + page}

	lastfmURL := c.getNoAuthURL(allOpts...)

	var topTagsRes struct {
		TopTags TopTags `json:"tags"`
	}

	err := c.get(lastfmURL, &topTagsRes)

	if err != nil {
		return TopTags{}, err
	}

	return topTagsRes.TopTags, nil
}

func (c *Client) ChartGetTopTracks(limit, page string) (ChartTracks, error) {
	// http://ws.audioscrobbler.com/2.0/?method=chart.gettoptracks&api_key=YOUR_API_KEY&format=json
	allOpts := []string{"method.chart.gettoptracks", "limit." + limit, "page." + page}

	lastfmURL := c.getNoAuthURL(allOpts...)

	var topTracksRes struct {
		TopTracks ChartTracks `json:"tracks"`
	}

	err := c.get(lastfmURL, &topTracksRes)

	if err != nil {
		return ChartTracks{}, err
	}

	return topTracksRes.TopTracks, nil
}
