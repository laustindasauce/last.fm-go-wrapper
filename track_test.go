package lastfm

import (
	"testing"
)

func TestTrackGetInfo(t *testing.T) {
	res, err := client.TrackGetInfo("believe", "cher")

	if err != nil {
		t.Error(err)
	}

	var tests = []struct {
		test     string
		expected string
	}{
		{"Name", "Believe"},
		{"MBID", "32ca187e-ee25-4f18-b7d0-3b6713f24635"},
		{"URL", "https://www.last.fm/music/Cher/_/Believe"},
		{"Duration", "240000"},
	}

	for _, test := range tests {
		if output := getStringField(res, test.test); output != test.expected {
			t.Errorf("Test Failed: %s expected, received: %s", test.expected, output)
		}
	}
}

func TestTrackGetCorrection(t *testing.T) {
	res, err := client.TrackGetCorrection("believe", "cher")

	if err != nil {
		t.Error(err)
	}

	var tests = []struct {
		test     string
		expected string
	}{
		{"Name", "Believe"},
		{"MBID", "32ca187e-ee25-4f18-b7d0-3b6713f24635"},
		{"URL", "https://www.last.fm/music/Cher/_/Believe"},
	}

	for _, test := range tests {
		if output := getStringField(res.Track, test.test); output != test.expected {
			t.Errorf("Test Failed: %s expected, received: %s", test.expected, output)
		}
	}
}

func TestTrackGetSimilar(t *testing.T) {
	const limit = 2
	res, err := client.TrackGetSimilar("believe", "cher", LimitOpt(limit))

	if err != nil {
		t.Error(err)
	}

	if count := len(res.Tracks); count != limit {
		t.Fatalf("Got %d similar tracks, wanted %d\n", count, limit)
	}

	var tests = []struct {
		test     string
		expected string
	}{
		{"Artist", "Cher"},
	}

	for _, test := range tests {
		if output := getStringField(res.Attr, test.test); output != test.expected {
			t.Errorf("Test Failed: %s expected, received: %s", test.expected, output)
		}
	}
}

func TestTrackGetTags(t *testing.T) {
	client.SetUser("RJ")

	res, err := client.TrackGetTags("AC/DC", "Hells Bells", client.User)

	if err != nil {
		t.Error(err)
	}

	if count := len(res.Tags); count == 0 {
		t.Fatalf("Got %d tags, wanted more\n", count)
	}

	if res.Attr.Track != "Hells Bells" {
		t.Errorf("Got incorrect Track name. Expected %s, got %s", "Hells Bells", res.Attr.Track)
	}

	if res.Attr.Artist != "AC/DC" {
		t.Errorf("Got incorrect Artist name. Expected %s, got %s", "AC/DC", res.Attr.Artist)
	}

	var tests = []struct {
		test     string
		expected string
	}{
		{"Name", "guitar"},
		{"URL", "https://www.last.fm/tag/guitar"},
	}

	for _, test := range tests {
		if output := getStringField(res.Tags[0], test.test); output != test.expected {
			t.Errorf("Test Failed: %s expected, received: %s", test.expected, output)
		}
	}
}

func TestTrackGetTopTags(t *testing.T) {
	client.SetUser("RJ")
	res, err := client.TrackGetTopTags("Cher", "Believe", client.User)

	if err != nil {
		t.Error(err)
	}

	if count := len(res.Tags); count != 100 {
		t.Fatalf("Got %d recent tracks, wanted 100\n", count)
	}

	if res.Attr.Track != "Believe" {
		t.Errorf("Got incorrect Track name. Expected %s, got %s", "Believe", res.Attr.Track)
	}

	if res.Attr.Artist != "Cher" {
		t.Errorf("Got incorrect Artist name. Expected %s, got %s", "Cher", res.Attr.Artist)
	}
}

func TestTrackSearch(t *testing.T) {
	res, err := client.TrackSearch("Believe", ArtistOpt("Cher"))

	if err != nil {
		t.Error(err)
	}

	if len(res.TrackMatches.Tracks) == 0 {
		t.Fatal("track.search returned empty list")
	}

	var tests = []struct {
		test     string
		expected string
	}{
		{"Name", "Believe"},
		{"Artist", "Cher"},
		{"URL", "https://www.last.fm/music/Cher/_/Believe"},
		{"MBID", "32ca187e-ee25-4f18-b7d0-3b6713f24635"},
	}

	for _, test := range tests {
		if output := getStringField(res.TrackMatches.Tracks[0], test.test); output != test.expected {
			t.Errorf("Test Failed: %s expected, received: %s", test.expected, output)
		}
	}
}
