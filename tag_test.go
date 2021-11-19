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
	res, err := client.TagGetTopAlbums("disco", "", "")

	if err != nil {
		t.Error(err)
	}

	if len(res.Albums) == 0 {
		t.Error("tag.gettopalbums returned empty list")
	}

	var tests = []struct {
		test     string
		expected string
	}{
		{"Tag", "disco"},
		{"Page", "1"},
		{"PerPage", "50"},
	}

	for _, test := range tests {
		if output := getStringField(res.Attr, test.test); output != test.expected {
			t.Errorf("Test Failed: %s expected, received: %s", test.expected, output)
		}
	}
}

func TestTagGetTopArtists(t *testing.T) {
	res, err := client.TagGetTopArtists("pop", "", "")

	if err != nil {
		t.Error(err)
	}

	if len(res.Artists) == 0 {
		t.Error("tag.gettopartists returned empty list")
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

	if len(res.Tag) == 0 {
		t.Error("tag.getTopTags returned empty list")
	}
}

func TestTagGetTopTracks(t *testing.T) {
	res, err := client.TagGetTopTracks("pop", "", "")

	if err != nil {
		t.Error(err)
	}

	if len(res.Tracks) == 0 {
		t.Error("tag.gettoptracks returned empty list")
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
		t.Error("tag.getweeklychartlist returned empty list")
	}

	var tests = []struct {
		test     string
		expected string
	}{
		{"Tag", "rock"},
	}

	for _, test := range tests {
		if output := getStringField(res.Attr, test.test); output != test.expected {
			t.Errorf("Test Failed: %s expected, received: %s", test.expected, output)
		}
	}
}
