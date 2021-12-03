package lastfm

import "fmt"

// TagInfo ...
type TagInfo struct {
	Name  string `json:"name"`
	Total int    `json:"total"`
	Reach int    `json:"reach"`
	Wiki  Wiki   `json:"wiki"`
}

// TagTopTags ...
type TagTopTags struct {
	Attr TopTagAttr `json:"@attr"`
	Tag  []TopTag   `json:"tag"`
}

// TagTopAlbums ...
type TagTopAlbums struct {
	Albums []TagAlbum `json:"album"`
	Attr   TagAttr    `json:"@attr"`
}

// TagTopArtists ...
type TagTopArtists struct {
	Artists []TagArtist `json:"artist"`
	Attr    TagAttr     `json:"@attr"`
}

// TagTopTracks ...
type TagTopTracks struct {
	Tracks []TagTrack `json:"track"`
	Attr   TagAttr    `json:"@attr"`
}

// TagWeeklyChartlist ...
type TagWeeklyChartlist struct {
	Charts []TagChartlist `json:"chart"`
	Attr   TagChartAttr   `json:"@attr"`
}

// TagAlbum ...
type TagAlbum struct {
	Name   string      `json:"name"`
	MBID   string      `json:"mbid"`
	URL    string      `json:"url"`
	Artist AlbumArtist `json:"artist"`
	Image  []Image     `json:"image"`
	Attr   Rank        `json:"@attr"`
}

// TagArtist ...
type TagArtist struct {
	Name       string  `json:"name"`
	MBID       string  `json:"mbid"`
	URL        string  `json:"url"`
	Streamable string  `json:"streamable"`
	Image      []Image `json:"image"`
	Attr       Rank    `json:"@attr"`
}

// TagTrack ...
type TagTrack struct {
	Name       string          `json:"name"`
	Duration   string          `json:"duration"`
	MBID       string          `json:"mbid"`
	URL        string          `json:"url"`
	Streamable StreamableTrack `json:"streamable"`
	Artist     AlbumArtist     `json:"artist"`
	Image      []Image         `json:"image"`
	Attr       Rank            `json:"@attr"`
}

// TagWithStrCount ...
type TagWithStrCount struct {
	Name  string `json:"name"`
	URL   string `json:"url"`
	Count string `json:"count"`
}

// TopTag ...
type TopTag struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
	Reach int    `json:"reach"`
}

// TagChartlist ...
type TagChartlist struct {
	Text string `json:"#text"`
	From string `json:"from"`
	To   string `json:"to"`
}

// TopTagAttr ...
type TopTagAttr struct {
	Offset int `json:"offset"`
	NumRes int `json:"num_res"`
	Total  int `json:"total"`
}

// TagChartAttr ...
type TagChartAttr struct {
	Tag string `json:"tag"`
}

// TagAttr ...
type TagAttr struct {
	Tag        string `json:"tag"`
	Page       string `json:"page"`
	PerPage    string `json:"perPage"`
	TotalPages string `json:"totalPages"`
	Total      string `json:"total"`
}

/*
Get the metadata for a tag

lang (Optional) : The language to return the wiki in, expressed as an ISO 639 alpha-2 code.

tag (Required) : The tag name

api_key (Required) : A Last.fm API key.
*/
func (c *Client) TagGetInfo(tag string, opts ...RequestOption) (TagInfo, error) {
	// http://ws.audioscrobbler.com/2.0/?method=tag.getinfo&tag=disco&api_key=YOUR_API_KEY&format=json
	lastfmURL := fmt.Sprintf("%s&method=tag.getinfo&tag=%s", c.baseApiURL, tag)

	values := processOptions(opts...).urlParams

	if query := values.Encode(); query != "" {
		lastfmURL += "&" + query
	}

	var tagInfo struct {
		Tag TagInfo `json:"tag"`
	}

	err := c.get(lastfmURL, &tagInfo)

	if err != nil {
		return TagInfo{}, err
	}

	return tagInfo.Tag, nil
}

/*
Get the top albums tagged by this tag, ordered by tag count.

tag (Required) : The tag name

limit (Optional) : The number of results to fetch per page. Defaults to 50.

page (Optional) : The page number to fetch. Defaults to first page.

api_key (Required) : A Last.fm API key.
*/
func (c *Client) TagGetTopAlbums(tag string, opts ...RequestOption) (TagTopAlbums, error) {
	// http://ws.audioscrobbler.com/2.0/?method=tag.gettopalbums&tag=disco&api_key=YOUR_API_KEY&format=json
	lastfmURL := fmt.Sprintf("%s&method=tag.gettopalbums&tag=%s", c.baseApiURL, tag)

	values := processOptions(opts...).urlParams

	if query := values.Encode(); query != "" {
		lastfmURL += "&" + query
	}

	var tagTopAlbums struct {
		Albums TagTopAlbums `json:"albums"`
	}

	err := c.get(lastfmURL, &tagTopAlbums)

	if err != nil {
		return TagTopAlbums{}, err
	}

	return tagTopAlbums.Albums, nil
}

/*
Get the top artists tagged by this tag, ordered by tag count.

tag (Required) : The tag name

limit (Optional) : The number of results to fetch per page. Defaults to 50.

page (Optional) : The page number to fetch. Defaults to first page.

api_key (Required) : A Last.fm API key.
*/
func (c *Client) TagGetTopArtists(tag string, opts ...RequestOption) (TagTopArtists, error) {
	// http://ws.audioscrobbler.com/2.0/?method=tag.gettopartists&tag=disco&api_key=YOUR_API_KEY&format=json
	lastfmURL := fmt.Sprintf("%s&method=tag.gettopartists&tag=%s", c.baseApiURL, tag)

	values := processOptions(opts...).urlParams

	if query := values.Encode(); query != "" {
		lastfmURL += "&" + query
	}

	var tagTopArtists struct {
		Artists TagTopArtists `json:"topartists"`
	}

	err := c.get(lastfmURL, &tagTopArtists)

	if err != nil {
		return TagTopArtists{}, err
	}

	return tagTopArtists.Artists, nil
}

/*
Fetches the top global tags on Last.fm, sorted by popularity (number of times used)

api_key (Required) : A Last.fm API key.
*/
func (c *Client) TagGetTopTags() (TagTopTags, error) {
	// http://ws.audioscrobbler.com/2.0/?method=tag.getTopTags&api_key=YOUR_API_KEY&format=json
	lastfmURL := fmt.Sprintf("%s&method=tag.getTopTags", c.baseApiURL)

	var tagTopTags struct {
		Tags TagTopTags `json:"toptags"`
	}

	err := c.get(lastfmURL, &tagTopTags)

	if err != nil {
		return TagTopTags{}, err
	}

	return tagTopTags.Tags, nil
}

/*
Get the top tracks tagged by this tag, ordered by tag count.

tag (Required) : The tag name

limit (Optional) : The number of results to fetch per page. Defaults to 50.

page (Optional) : The page number to fetch. Defaults to first page.

api_key (Required) : A Last.fm API key.
*/
func (c *Client) TagGetTopTracks(tag string, opts ...RequestOption) (TagTopTracks, error) {
	// http://ws.audioscrobbler.com/2.0/?method=tag.gettoptracks&tag=disco&api_key=YOUR_API_KEY&format=json
	lastfmURL := fmt.Sprintf("%s&method=tag.gettoptracks&tag=%s", c.baseApiURL, tag)

	values := processOptions(opts...).urlParams

	if query := values.Encode(); query != "" {
		lastfmURL += "&" + query
	}

	var tagTopTracks struct {
		Tracks TagTopTracks `json:"tracks"`
	}

	err := c.get(lastfmURL, &tagTopTracks)

	if err != nil {
		return TagTopTracks{}, err
	}

	return tagTopTracks.Tracks, nil
}

/*
Get a list of available charts for this tag, expressed as date ranges which can be sent to the chart services.

tag (Required) : The tag name

api_key (Required) : A Last.fm API key.
*/
func (c *Client) TagGetWeeklyChartlist(tag string) (TagWeeklyChartlist, error) {
	// http://ws.audioscrobbler.com/2.0/?method=tag.getweeklychartlist&tag=disco&api_key=YOUR_API_KEY&format=json
	lastfmURL := fmt.Sprintf("%s&method=tag.getweeklychartlist&tag=%s", c.baseApiURL, tag)

	var tagChartlist struct {
		Chartlist TagWeeklyChartlist `json:"weeklychartlist"`
	}

	err := c.get(lastfmURL, &tagChartlist)

	if err != nil {
		return TagWeeklyChartlist{}, err
	}

	return tagChartlist.Chartlist, nil
}
