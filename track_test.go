package lastfm

import (
	"strconv"
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
	const limit = "2"
	res, err := client.TrackGetSimilar("believe", "cher", "", limit)

	if err != nil {
		t.Error(err)
	}

	if len(res.Tracks) == 0 {
		t.Error("track.getsimilar returned empty list")
	}

	// Convert the limit string variable to int
	intVar, _ := strconv.Atoi(limit)

	if len(res.Tracks) != intVar {
		t.Error("track.getsimilar returned the wrong amount of tracks")
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

	res, err := client.TrackGetTags("Hells Bells", "AC/DC", "")

	if err != nil {
		t.Error(err)
	}

	if len(res.Tags) == 0 {
		t.Error("track.getTags returned empty list")
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
	res, err := client.TrackGetTopTags("Believe", "Cher", "")

	if err != nil {
		t.Error(err)
	}

	if len(res.Tags) == 0 {
		t.Error("track.gettoptags returned empty list")
	}

	if len(res.Tags) != 100 {
		t.Error("track.gettoptags didn't return 100 results")
	}

	var tests = []struct {
		test     string
		expected string
	}{
		{"Artist", "Cher"},
		{"Track", "Believe"},
	}

	for _, test := range tests {
		if output := getStringField(res.Attr, test.test); output != test.expected {
			t.Errorf("Test Failed: %s expected, received: %s", test.expected, output)
		}
	}
}

func TestTrackSearch(t *testing.T) {
	res, err := client.TrackSearch("Believe", "Cher", "", "")

	if err != nil {
		t.Error(err)
	}

	if len(res.TrackMatches.Tracks) == 0 {
		t.Error("track.search returned empty list")
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
