package lastfm

type TopArtists struct {
	Artist []ChartArtist `json:"artist"`
	Attr ChartAttr `json:"@attr"`
}

type ChartArtist struct {
	Name string `json:"name"`
	Playcount string `json:"playcount"`
	Listeners string `json:"listeners"`
	MBID string `json:"mbid"`
	URL string `json:"url"`
	Streamable string `json:"streamable"`
	Image []Image `json:"image"`
}

type ChartAttr struct {
	Page string `json:"page"`
	PerPage string `json:"perPage"`
	TotalPages string `json:"totalPages"`
	Total string `json:"total"`
}

func (c *Client) ChartGetTopArtists() (TopArtists, error) {
	// http://ws.audioscrobbler.com/2.0/?method=chart.gettopartists&api_key=YOUR_API_KEY&format=json
	lastfmURL := c.getNoAuthURL("method.chart.gettopartists")

	var topArtistsRes struct {
		TopArtists TopArtists `json:"artists"`
	}

	err := c.get(lastfmURL, &topArtistsRes)

	if err != nil {
		return TopArtists{}, err
	}

	return topArtistsRes.TopArtists, nil
}

func (c *Client) ChartGetTopTags() (string, error) {
	// http://ws.audioscrobbler.com/2.0/?method=chart.gettoptags&api_key=YOUR_API_KEY&format=json
	lastfmURL := c.getNoAuthURL("method.chart.gettoptags")
	return lastfmURL, nil
}

func (c *Client) ChartGetTopTracks() (string, error) {
	// http://ws.audioscrobbler.com/2.0/?method=chart.gettoptracks&api_key=YOUR_API_KEY&format=json
	lastfmURL := c.getNoAuthURL("method.chart.gettoptracks")
	return lastfmURL, nil
}