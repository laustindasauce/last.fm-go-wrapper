package lastfm

import "testing"

func TestLibraryGetArtists(t *testing.T) {
	client.SetUser("abspen1")
	res, err := client.LibraryGetArtists(client.User)

	if err != nil {
		t.Error(err)
	}

	var tests = []struct {
		test     string
		expected string
	}{
		{"User", "Abspen1"},
		{"Page", "1"},
		{"PerPage", "50"},
	}

	for _, test := range tests {
		if output := getStringField(res.Attr, test.test); output != test.expected {
			t.Errorf("Test Failed: %s expected, received: %s", test.expected, output)
		}
	}
}
