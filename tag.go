package lastfm

type TagInfo struct {
	Name  string `json:"name"`
	Total int    `json:"total"`
	Reach int    `json:"reach"`
	Wiki  Wiki   `json:"wiki"`
}

type TagTopTags struct {
	Attr TopTagAttr `json:"@attr"`
	Tag  []TopTag   `json:"tag"`
}

type TagTopAlbums struct {
	Albums []TagAlbum `json:"album"`
	Attr   TagAttr    `json:"@attr"`
}

type TagTopArtists struct {
	Artists []TagArtist `json:"artist"`
	Attr    TagAttr     `json:"@attr"`
}

type TagTopTracks struct {
	Tracks []TagTrack `json:"track"`
	Attr   TagAttr    `json:"@attr"`
}

type TagWeeklyChartlist struct {
	Charts []TagChartlist `json:"chart"`
	Attr   TagChartAttr   `json:"@attr"`
}

type TagAlbum struct {
	Name   string      `json:"name"`
	MBID   string      `json:"mbid"`
	URL    string      `json:"url"`
	Artist AlbumArtist `json:"artist"`
	Image  []Image     `json:"image"`
	Attr   Rank        `json:"@attr"`
}

type TagArtist struct {
	Name       string  `json:"name"`
	MBID       string  `json:"mbid"`
	URL        string  `json:"url"`
	Streamable string  `json:"streamable"`
	Image      []Image `json:"image"`
	Attr       Rank    `json:"@attr"`
}

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

type TopTag struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
	Reach int    `json:"reach"`
}

type TagChartlist struct {
	Text string `json:"#text"`
	From string `json:"from"`
	To   string `json:"to"`
}

type TopTagAttr struct {
	Offset int `json:"offset"`
	NumRes int `json:"num_res"`
	Total  int `json:"total"`
}

type TagChartAttr struct {
	Tag string `json:"tag"`
}

type TagAttr struct {
	Tag        string `json:"tag"`
	Page       string `json:"page"`
	PerPage    string `json:"perPage"`
	TotalPages string `json:"totalPages"`
	Total      string `json:"total"`
}

func (c *Client) TagGetInfo(tag string) (TagInfo, error) {
	// http://ws.audioscrobbler.com/2.0/?method=library.getartists&api_key=YOUR_API_KEY&user=joanofarctan&format=json
	lastfmURL := c.getNoAuthURL("method.tag.getinfo", "tag."+tag)

	var tagInfo struct {
		Tag TagInfo `json:"tag"`
	}

	err := c.get(lastfmURL, &tagInfo)

	if err != nil {
		return TagInfo{}, err
	}

	return tagInfo.Tag, nil
}

func (c *Client) TagGetTopAlbums(tag string) (TagTopAlbums, error) {
	// http://ws.audioscrobbler.com/2.0/?method=library.getartists&api_key=YOUR_API_KEY&user=joanofarctan&format=json
	lastfmURL := c.getNoAuthURL("method.tag.gettopalbums", "tag."+tag)

	var tagTopAlbums struct {
		Albums TagTopAlbums `json:"albums"`
	}

	err := c.get(lastfmURL, &tagTopAlbums)

	if err != nil {
		return TagTopAlbums{}, err
	}

	return tagTopAlbums.Albums, nil
}

func (c *Client) TagGetTopArtists(tag string) (TagTopArtists, error) {
	// http://ws.audioscrobbler.com/2.0/?method=library.getartists&api_key=YOUR_API_KEY&user=joanofarctan&format=json
	lastfmURL := c.getNoAuthURL("method.tag.gettopartists", "tag."+tag)

	var tagTopArtists struct {
		Artists TagTopArtists `json:"topartists"`
	}

	err := c.get(lastfmURL, &tagTopArtists)

	if err != nil {
		return TagTopArtists{}, err
	}

	return tagTopArtists.Artists, nil
}

func (c *Client) TagGetTopTags() (TagTopTags, error) {
	// http://ws.audioscrobbler.com/2.0/?method=library.getartists&api_key=YOUR_API_KEY&user=joanofarctan&format=json
	lastfmURL := c.getNoAuthURL("method.tag.getTopTags")

	var tagTopTags struct {
		Tags TagTopTags `json:"toptags"`
	}

	err := c.get(lastfmURL, &tagTopTags)

	if err != nil {
		return TagTopTags{}, err
	}

	return tagTopTags.Tags, nil
}

func (c *Client) TagGetTopTracks(tag string) (TagTopTracks, error) {
	// http://ws.audioscrobbler.com/2.0/?method=library.getartists&api_key=YOUR_API_KEY&user=joanofarctan&format=json
	lastfmURL := c.getNoAuthURL("method.tag.gettoptracks", "tag."+tag)

	var tagTopTracks struct {
		Tracks TagTopTracks `json:"tracks"`
	}

	err := c.get(lastfmURL, &tagTopTracks)

	if err != nil {
		return TagTopTracks{}, err
	}

	return tagTopTracks.Tracks, nil
}

func (c *Client) TagGetWeeklyChartlist(tag string) (TagWeeklyChartlist, error) {
	// http://ws.audioscrobbler.com/2.0/?method=library.getartists&api_key=YOUR_API_KEY&user=joanofarctan&format=json
	lastfmURL := c.getNoAuthURL("method.tag.getweeklychartlist", "tag."+tag)

	var tagChartlist struct {
		Chartlist TagWeeklyChartlist `json:"weeklychartlist"`
	}

	err := c.get(lastfmURL, &tagChartlist)

	if err != nil {
		return TagWeeklyChartlist{}, err
	}

	return tagChartlist.Chartlist, nil
}
