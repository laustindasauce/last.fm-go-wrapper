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