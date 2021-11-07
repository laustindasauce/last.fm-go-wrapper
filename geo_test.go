package lastfm

import "testing"

func TestGeoGetTopArtists(t *testing.T) {
	res, err := client.GeoGetTopArtists("usa")

	if err != nil {
		t.Error(err)
	}

	var tests = []struct {
		test     string
		expected string
	}{
		{"Country", "United States"},
		{"Page", "1"},
		{"PerPage", "50"},
	}

	for _, test := range tests {
		if output := getStringField(res.Attr, test.test); output != test.expected {
			t.Errorf("Test Failed: %s expected, received: %s", test.expected, output)
		}
	}
}
