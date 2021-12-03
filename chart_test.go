package lastfm

import "testing"

func TestChartGetTopArtists(t *testing.T) {
	res, err := client.ChartGetTopArtists()

	if err != nil {
		t.Error(err)
	}

	if count := len(res.Artist); count != 50 {
		t.Fatalf("Got %d chart top artists, wanted 50\n", count)
	}

	if res.Attr.Page != "1" {
		t.Errorf("Got incorrect Page. Expected %s, got %s", "Believe", res.Attr.Page)
	}

	if res.Attr.PerPage != "50" {
		t.Errorf("Got incorrect PerPage. Expected %s, got %s", "Cher", res.Attr.PerPage)
	}
}

func TestChartGetTopTags(t *testing.T) {
	res, err := client.ChartGetTopTags()

	if err != nil {
		t.Error(err)
	}

	if count := len(res.Tag); count != 50 {
		t.Fatalf("Got %d chart top tags, wanted 50\n", count)
	}

	if res.Attr.Page != "1" {
		t.Errorf("Got incorrect Page. Expected %s, got %s", "Believe", res.Attr.Page)
	}

	if res.Attr.PerPage != "50" {
		t.Errorf("Got incorrect PerPage. Expected %s, got %s", "Cher", res.Attr.PerPage)
	}
}

func TestChartGetTopTracks(t *testing.T) {
	res, err := client.ChartGetTopTracks()

	if err != nil {
		t.Error(err)
	}

	if count := len(res.Track); count != 50 {
		t.Fatalf("Got %d chart top tracks, wanted 50\n", count)
	}

	if res.Attr.Page != "1" {
		t.Errorf("Got incorrect Page. Expected %s, got %s", "Believe", res.Attr.Page)
	}

	if res.Attr.PerPage != "50" {
		t.Errorf("Got incorrect PerPage. Expected %s, got %s", "Cher", res.Attr.PerPage)
	}
}
