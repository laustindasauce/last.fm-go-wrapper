package lastfm

import (
	"testing"
)

func TestUserGetFriends(t *testing.T) {
	err := client.SetUser("rj")

	if err != nil {
		t.Error(err)
	}

	res, err := client.UserGetFriends(client.User, LimitOpt(2))

	if err != nil {
		t.Error(err)
	}

	var tests = []struct {
		test     string
		expected string
	}{
		{"User", "RJ"},
		{"Page", "1"},
		{"PerPage", "2"},
	}

	for _, test := range tests {
		if output := getStringField(res.Attr, test.test); output != test.expected {
			t.Errorf("Test Failed: %s expected, received: %s", test.expected, output)
		}
	}
}

func TestUserGetInfo(t *testing.T) {
	err := client.SetUser("rj")

	if err != nil {
		t.Error(err)
	}

	res, err := client.UserGetInfo(client.User)

	if err != nil {
		t.Error(err)
	}

	var tests = []struct {
		test     string
		expected string
	}{
		{"Country", "United Kingdom"},
		{"Realname", "Richard Jones "},
		{"URL", "https://www.last.fm/user/RJ"},
		{"Gender", "n"},
		{"Name", "RJ"},
		{"Type", "alum"},
	}

	for _, test := range tests {
		if output := getStringField(res, test.test); output != test.expected {
			t.Errorf("Test Failed: %s expected, received: %s", test.expected, output)
		}
	}
}

func TestUserGetLovedTracks(t *testing.T) {
	err := client.SetUser("rj")

	if err != nil {
		t.Error(err)
	}

	res, err := client.UserGetLovedTracks(client.User, LimitOpt(2))

	if err != nil {
		t.Error(err)
	}

	var tests = []struct {
		test     string
		expected string
	}{
		{"Page", "1"},
		{"User", "RJ"},
		{"PerPage", "2"},
	}

	for _, test := range tests {
		if output := getStringField(res.Attr, test.test); output != test.expected {
			t.Errorf("Test Failed: %s expected, received: %s", test.expected, output)
		}
	}
}

func TestUserGetPersonalAlbumTags(t *testing.T) {
	err := client.SetUser("abspen1")

	if err != nil {
		t.Error(err)
	}

	res, err := client.UserGetPersonalAlbumTags(client.User, "love", LimitOpt(2))

	if err != nil {
		t.Error(err)
	}

	var tests = []struct {
		test     string
		expected string
	}{
		{"Page", "1"},
		{"User", "Abspen1"},
		{"PerPage", "2"},
		{"Tag", "love"},
	}

	for _, test := range tests {
		if output := getStringField(res.Attr, test.test); output != test.expected {
			t.Errorf("Test Failed: %s expected, received: %s", test.expected, output)
		}
	}
}

func TestUserGetPersonalArtistTags(t *testing.T) {
	err := client.SetUser("rj")

	if err != nil {
		t.Error(err)
	}

	res, err := client.UserGetPersonalArtistTags(client.User, "rock", LimitOpt(2))

	if err != nil {
		t.Error(err)
	}

	var tests = []struct {
		test     string
		expected string
	}{
		{"Page", "1"},
		{"User", "RJ"},
		{"PerPage", "2"},
		{"Tag", "rock"},
	}

	for _, test := range tests {
		if output := getStringField(res.Attr, test.test); output != test.expected {
			t.Errorf("Test Failed: %s expected, received: %s", test.expected, output)
		}
	}
}

func TestUserGetPersonalTrackTags(t *testing.T) {
	err := client.SetUser("rj")

	if err != nil {
		t.Error(err)
	}

	res, err := client.UserGetPersonalTrackTags(client.User, "rock", LimitOpt(2))

	if err != nil {
		t.Error(err)
	}

	var tests = []struct {
		test     string
		expected string
	}{
		{"Page", "1"},
		{"User", "RJ"},
		{"PerPage", "2"},
		{"Tag", "rock"},
	}

	for _, test := range tests {
		if output := getStringField(res.Attr, test.test); output != test.expected {
			t.Errorf("Test Failed: %s expected, received: %s", test.expected, output)
		}
	}
}
