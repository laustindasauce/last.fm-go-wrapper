package lastfm

import "testing"

func TestTagGetInfo(t *testing.T) {
	res, err := client.TagGetInfo("disco")

	if err != nil {
		t.Error(err)
	}

	var tests = []struct {
		test     string
		expected string
	}{
		{"Name", "disco"},
	}

	for _, test := range tests {
		if output := getStringField(res, test.test); output != test.expected {
			t.Errorf("Test Failed: %s expected, received: %s", test.expected, output)
		}
	}
}

func TestTagGetTopAlbums(t *testing.T) {
	res, err := client.TagGetTopAlbums("disco", LimitOpt(2))

	if err != nil {
		t.Error(err)
	}

	if count := len(res.Albums); count != 2 {
		t.Fatalf("Got %d top albums, wanted 2\n", count)
	}

	var tests = []struct {
		test     string
		expected string
	}{
		{"Tag", "disco"},
		{"Page", "1"},
		{"PerPage", "2"},
	}

	for _, test := range tests {
		if output := getStringField(res.Attr, test.test); output != test.expected {
			t.Errorf("Test Failed: %s expected, received: %s", test.expected, output)
		}
	}
}

func TestTagGetTopArtists(t *testing.T) {
	res, err := client.TagGetTopArtists("pop")

	if err != nil {
		t.Error(err)
	}

	if count := len(res.Artists); count != 50 {
		t.Fatalf("Got %d recent tracks, wanted 50\n", count)
	}

	var tests = []struct {
		test     string
		expected string
	}{
		{"Tag", "pop"},
		{"Page", "1"},
		{"PerPage", "50"},
	}

	for _, test := range tests {
		if output := getStringField(res.Attr, test.test); output != test.expected {
			t.Errorf("Test Failed: %s expected, received: %s", test.expected, output)
		}
	}
}

func TestTagGetTopTags(t *testing.T) {
	res, err := client.TagGetTopTags()

	if err != nil {
		t.Error(err)
	}

	if count := len(res.Tag); count != 50 {
		t.Fatalf("Got %d top tags, wanted 50\n", count)
	}
}

func TestTagGetTopTracks(t *testing.T) {
	res, err := client.TagGetTopTracks("pop")

	if err != nil {
		t.Error(err)
	}

	if count := len(res.Tracks); count != 50 {
		t.Fatalf("Got %d recent tracks, wanted 50\n", count)
	}

	var tests = []struct {
		test     string
		expected string
	}{
		{"Tag", "pop"},
		{"Page", "1"},
		{"PerPage", "50"},
	}

	for _, test := range tests {
		if output := getStringField(res.Attr, test.test); output != test.expected {
			t.Errorf("Test Failed: %s expected, received: %s", test.expected, output)
		}
	}
}

func TestTagGetWeeklyChartlist(t *testing.T) {
	res, err := client.TagGetWeeklyChartlist("rock")

	if err != nil {
		t.Error(err)
	}

	if len(res.Charts) == 0 {
		t.Fatal("tag.getweeklychartlist returned empty list")
	}

	if res.Attr.Tag != "rock" {
		t.Errorf("Got incorrect tag. Expected %s, got %s", "rock", res.Attr.Tag)
	}
}
