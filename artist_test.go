package lastfm

import (
	"net/http"
	"os"
	"testing"
	"time"
)

var hClient http.Client = http.Client{Timeout: time.Duration(1) * time.Second}
var client *Client = New(&hClient, os.Getenv("LAST_FM_KEY"), os.Getenv("LAST_FM_SECRET"))

func TestArtistGetInfo(t *testing.T) {
	res, err := client.ArtistGetInfo("cher")

	if err != nil {
		t.Error(err)
	}

	if res == nil {
		t.Error("artist.getinfo returned nil")
		return
	}

	if res.Name != "Cher" {
		t.Error("artist.getinfo returned the wrong artist")
	}
}

func TestArtistGetSimilar(t *testing.T) {
	res, err := client.ArtistGetSimilar("cher")

	if err != nil {
		t.Error(err)
	}

	if err != nil {
		t.Error(err)
	}

	if len(res.SimilarArtists) == 0 {
		t.Error("artist.getsimilar returned an empty array")
	}
}

func TestArtistGetTopAlbums(t *testing.T) {
	res, err := client.ArtistGetTopAlbums("cher")

	if err != nil {
		t.Error(err)
	}

	if len(res.Album) == 0 {
		t.Error("artist.gettopalbums returned an empty array")
	}
}

func TestArtistGetTopTags(t *testing.T) {
	res, err := client.ArtistGetTopTags("cher")

	if err != nil {
		t.Error(err)
	}

	if len(res.Tag) == 0 {
		t.Error("artist.gettoptags returned an empty array")
	}
}

func TestArtistGetTopTracks(t *testing.T) {
	res, err := client.ArtistGetTopTracks("cher")

	if err != nil {
		t.Error(err)
	}

	if len(res.Track) == 0 {
		t.Error("artist.gettoptracks returned an empty array")
	}
}

func TestArtistSearch(t *testing.T) {
	res, err := client.ArtistSearch("cher")

	if err != nil {
		t.Error(err)
	}

	if len(res.ArtistMatches.Artist) == 0 {
		t.Error("artist.gettoptracks returned an empty array")
	}
}
