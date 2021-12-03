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
		t.Fatal(err)
	}

	if count := len(res.Users); count != 2 {
		t.Fatalf("Got %d users, wanted 2\n", count)
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
		t.Fatal(err)
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
		t.Fatal(err)
	}

	if count := len(res.Tracks); count != 2 {
		t.Fatalf("Got %d loved tracks, wanted 2\n", count)
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
		t.Fatal(err)
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
		t.Fatal(err)
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
		t.Fatal(err)
	}

	if count := len(res.TrackTags.Tracks); count != 2 {
		t.Fatalf("Got %d track tags, wanted 2\n", count)
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

func TestUserGetRecentTracks(t *testing.T) {
	err := client.SetUser("rj")

	if err != nil {
		t.Error(err)
	}

	res, err := client.UserGetRecentTracks(client.User, LimitOpt(2))

	if err != nil {
		t.Fatal(err)
	}

	if count := len(res.Tracks); count != 2 {
		t.Fatalf("Got %d recent tracks, wanted 2\n", count)
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

func TestUserGetTopAlbums(t *testing.T) {
	err := client.SetUser("abspen1")

	if err != nil {
		t.Error(err)
	}

	res, err := client.UserGetTopAlbums(client.User, LimitOpt(2))

	if err != nil {
		t.Fatal(err)
	}

	if count := len(res.Albums); count != 2 {
		t.Fatalf("Got %d top albums, wanted 2\n", count)
	}

	if res.Albums[0].Name != "~how i'm feeling~" {
		t.Errorf("Got incorrect top album. Expected %s, got %s", "~how i'm feeling~", res.Albums[0].Name)
	}

	var tests = []struct {
		test     string
		expected string
	}{
		{"Page", "1"},
		{"User", "Abspen1"},
		{"PerPage", "2"},
	}

	for _, test := range tests {
		if output := getStringField(res.Attr, test.test); output != test.expected {
			t.Errorf("Test Failed: %s expected, received: %s", test.expected, output)
		}
	}
}

func TestUserGetTopArtists(t *testing.T) {
	err := client.SetUser("abspen1")

	if err != nil {
		t.Error(err)
	}

	res, err := client.UserGetTopArtists(client.User, LimitOpt(2))

	if err != nil {
		t.Fatal(err)
	}

	if count := len(res.Artists); count != 2 {
		t.Fatalf("Got %d top artists, wanted 2\n", count)
	}

	if res.Artists[0].Name != "Lauv" {
		t.Errorf("Got incorrect top artist. Expected %s, got %s", "Lauv", res.Artists[0].Name)
	}

	var tests = []struct {
		test     string
		expected string
	}{
		{"Page", "1"},
		{"User", "Abspen1"},
		{"PerPage", "2"},
	}

	for _, test := range tests {
		if output := getStringField(res.Attr, test.test); output != test.expected {
			t.Errorf("Test Failed: %s expected, received: %s", test.expected, output)
		}
	}
}

func TestUserGetTopTags(t *testing.T) {
	err := client.SetUser("abspen1")

	if err != nil {
		t.Error(err)
	}

	res, err := client.UserGetTopTags(client.User, LimitOpt(2))

	if err != nil {
		t.Fatal(err)
	}

	if count := len(res.Tags); count != 2 {
		t.Fatalf("Got %d top tags, wanted 2\n", count)
	}

	if res.Attr.User != "Abspen1" {
		t.Errorf("Got incorrect top tag. Expected %s, got %s", "Abspen1", res.Tags[0].Name)
	}

	var tests = []struct {
		test     string
		expected string
	}{
		{"Name", "pop"},
		{"Count", "1"},
		{"URL", "https://www.last.fm/tag/pop"},
	}

	for _, test := range tests {
		if output := getStringField(res.Tags[0], test.test); output != test.expected {
			t.Errorf("Test Failed: %s expected, received: %s", test.expected, output)
		}
	}
}

func TestUserGetTopTracks(t *testing.T) {
	err := client.SetUser("abspen1")

	if err != nil {
		t.Error(err)
	}

	res, err := client.UserGetTopTracks(client.User, LimitOpt(2))

	if err != nil {
		t.Fatal(err)
	}

	if count := len(res.Tracks); count != 2 {
		t.Fatalf("Got %d top tracks, wanted 2\n", count)
	}

	if res.Tracks[0].Name != "i'm so tired..." {
		t.Errorf("Got incorrect top track. Expected %s, got %s", "i'm so tired...", res.Tracks[0].Name)
	}

	var tests = []struct {
		test     string
		expected string
	}{
		{"User", "Abspen1"},
		{"Page", "1"},
		{"PerPage", "2"},
	}

	for _, test := range tests {
		if output := getStringField(res.Attr, test.test); output != test.expected {
			t.Errorf("Test Failed: %s expected, received: %s", test.expected, output)
		}
	}
}

func TestUserGetWeeklyAlbumChart(t *testing.T) {
	err := client.SetUser("abspen1")

	if err != nil {
		t.Error(err)
	}

	res, err := client.UserGetWeeklyAlbumChart(client.User, LimitOpt(2))

	if err != nil {
		t.Fatal(err)
	}

	if count := len(res.Albums); count != 2 {
		t.Fatalf("Got %d weekly albums, wanted 2\n", count)
	}

	if res.Attr.User != "Abspen1" {
		t.Errorf("Got incorrect user. Expected %s, got %s", "Abspen1", res.Attr.User)
	}
}

func TestUserGetWeeklyArtistChart(t *testing.T) {
	err := client.SetUser("abspen1")

	if err != nil {
		t.Error(err)
	}

	res, err := client.UserGetWeeklyArtistChart(client.User, LimitOpt(2))

	if err != nil {
		t.Fatal(err)
	}

	if count := len(res.Artists); count != 2 {
		t.Fatalf("Got %d weekly artists, wanted 2\n", count)
	}

	if res.Attr.User != "Abspen1" {
		t.Errorf("Got incorrect user. Expected %s, got %s", "Abspen1", res.Attr.User)
	}
}

func TestUserGetWeeklyChartList(t *testing.T) {
	err := client.SetUser("abspen1")

	if err != nil {
		t.Error(err)
	}

	res, err := client.UserGetWeeklyChartList(client.User)

	if err != nil {
		t.Fatal(err)
	}

	if count := len(res.Charts); count < 876 {
		t.Fatalf("Got %d weekly artists, wanted >= 876\n", count)
	}

	if res.Attr.User != "Abspen1" {
		t.Errorf("Got incorrect user. Expected %s, got %s", "Abspen1", res.Attr.User)
	}
}

func TestUserGetWeeklyTrackChart(t *testing.T) {
	err := client.SetUser("abspen1")

	if err != nil {
		t.Error(err)
	}

	res, err := client.UserGetWeeklyTrackChart(client.User, LimitOpt(2))

	if err != nil {
		t.Fatal(err)
	}

	if count := len(res.Tracks); count != 2 {
		t.Fatalf("Got %d weekly tracks, wanted 2\n", count)
	}

	if res.Attr.User != "Abspen1" {
		t.Errorf("Got incorrect user. Expected %s, got %s", "Abspen1", res.Attr.User)
	}
}
