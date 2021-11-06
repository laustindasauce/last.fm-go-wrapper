package lastfm

import (
	"reflect"
	"testing"
)

func getStringField(obj interface{}, field string) string {
	r := reflect.ValueOf(obj)
	f := reflect.Indirect(r).FieldByName(field)
	return f.String()
}

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

	var tests = []struct {
		test     string
		expected string
	}{
		{"Artist", "Cher"},
		{"Album", "Believe"},
	}

	for _, test := range tests {
		if output := getStringField(res.Attr, test.test); output != test.expected {
			t.Errorf("Test Failed: %s expected, received: %s", test.expected, output)
		}
	}
}

func TestAlbumSearch(t *testing.T) {
	res, err := client.AlbumSearch("Believe")

	if err != nil {
		t.Error(err)
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
