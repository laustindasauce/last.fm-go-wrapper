package lastfm

import (
	"net/http"
	"os"
	"reflect"
	"testing"
	"time"
)

var (
	hClient http.Client = http.Client{Timeout: time.Duration(1) * time.Second}
	client  *Client     = New(&hClient, os.Getenv("LAST_FM_KEY"), os.Getenv("LAST_FM_SECRET"))
)

func getStringField(obj interface{}, field string) string {
	r := reflect.ValueOf(obj)
	f := reflect.Indirect(r).FieldByName(field)
	return f.String()
}

// func getIntField(obj interface{}, field string) int64 {
// 	r := reflect.ValueOf(obj)
// 	f := reflect.Indirect(r).FieldByName(field)
// 	return f.Int()
// }

// Testing the getNoAuthURL and the encodeParams funcs
func TestGetNoAuthURL(t *testing.T) {
	var (
		hClient http.Client = http.Client{Timeout: time.Duration(1) * time.Second}
		client  *Client     = New(&hClient, os.Getenv("LAST_FM_KEY"), os.Getenv("LAST_FM_SECRET"))
	)

	var tests = []struct {
		input    []string
		expected string
	}{
		{[]string{"method.album.getinfo", "artist.Cher", "album.believe"}, "https://ws.audioscrobbler.com/2.0/?album=believe&api_key=" + os.Getenv("LAST_FM_KEY") + "&artist=Cher&format=json&method=album.getinfo"},
		{[]string{"artist.anime", "method.artist.getinfo"}, "https://ws.audioscrobbler.com/2.0/?api_key=" + os.Getenv("LAST_FM_KEY") + "&artist=anime&format=json&method=artist.getinfo"},
	}

	for _, test := range tests {
		if output := client.getNoAuthURL(test.input...); output != test.expected {
			t.Errorf("Test Failed: %s inputted, %s expected, received: %s", test.input, test.expected, output)
		}
	}
}
