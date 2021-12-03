package lastfm

import (
	"testing"
)

func TestAlbumGetInfo(t *testing.T) {
	res, err := client.AlbumGetInfo("Believe", "Cher")

	if err != nil {
		t.Error(err)
	}

	var tests = []struct {
		test     string
		expected string
	}{
		{"Artist", "Cher"},
		{"Name", "Believe"},
		{"URL", "https://www.last.fm/music/Cher/Believe"},
		{"MBID", "03c91c40-49a6-44a7-90e7-a700edf97a62"},
	}

	for _, test := range tests {
		if output := getStringField(res, test.test); output != test.expected {
			t.Errorf("Test Failed: %s expected, received: %s", test.expected, output)
		}
	}
}

func TestAlbumTopTags(t *testing.T) {
	res, err := client.AlbumGetTopTags("Believe", "Cher")

	if err != nil {
		t.Error(err)
	}

	if count := len(res.Tag); count != 100 {
		t.Fatalf("Got %d artist matches, wanted 100\n", count)
	}

	if res.Attr.Album != "Believe" {
		t.Errorf("Got incorrect Album name. Expected %s, got %s", "Believe", res.Attr.Album)
	}

	if res.Attr.Artist != "Cher" {
		t.Errorf("Got incorrect Artist name. Expected %s, got %s", "Cher", res.Attr.Artist)
	}
}

func TestAlbumSearch(t *testing.T) {
	res, err := client.AlbumSearch("Believe", LimitOpt(2))

	if err != nil {
		t.Error(err)
	}

	if count := len(res.AlbumMatches.Album); count != 2 {
		t.Fatalf("Got %d artist matches, wanted 2\n", count)
	}

	var tests = []struct {
		test     string
		expected string
	}{
		{"Text", ""},
		{"Role", "request"},
		{"SearchTerms", "Believe"},
		{"StartPage", "1"},
	}

	for _, test := range tests {
		if output := getStringField(res.Query, test.test); output != test.expected {
			t.Errorf("Test Failed: %s expected, received: %s", test.expected, output)
		}
	}
}
