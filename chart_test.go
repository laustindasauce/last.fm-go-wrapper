package lastfm

import "testing"

func TestChartGetTopArtists(t *testing.T) {
	res, err := client.ChartGetTopArtists()

	if err != nil {
		t.Error(err)
	}

	if len(res.Artist) == 0 {
		t.Error("artist.gettoptags returned an empty array")
	}
}

func TestChartGetTopTags(t *testing.T) {
	res, err := client.ChartGetTopTags()

	if err != nil {
		t.Error(err)
	}

	if len(res.Tag) == 0 {
		t.Error("artist.gettoptags returned an empty array")
	}
}

func TestChartGetTopTracks(t *testing.T) {
	res, err := client.ChartGetTopTracks()

	if err != nil {
		t.Error(err)
	}

	if len(res.Track) == 0 {
		t.Error("artist.gettoptags returned an empty array")
	}
}