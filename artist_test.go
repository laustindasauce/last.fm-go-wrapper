package lastfm

import (
	"testing"
)

func TestArtistGetInfo(t *testing.T) {
	res, err := client.ArtistGetInfo("cher")

	if err != nil {
		t.Error(err)
	}

	if res.Name != "Cher" {
		t.Fatalf("artist.getinfo returned the wrong artist. Expected: %s Received: %s", "Cher", res.Name)
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

	if res.Attr.Artist != "Cher" {
		t.Fatalf("artist.getsimilar returned the wrong artist. Expected: %s Received: %s", "Cher", res.Attr.Artist)
	}

	if count := len(res.SimilarArtists); count != 100 {
		t.Fatalf("Got %d artist similar artists, wanted 100\n", count)
	}
}

func TestArtistGetTopAlbums(t *testing.T) {
	res, err := client.ArtistGetTopAlbums("cher")

	if err != nil {
		t.Error(err)
	}

	if res.Attr.Artist != "Cher" {
		t.Fatalf("artist.gettopalbums returned the wrong artist. Expected: %s Received: %s", "Cher", res.Attr.Artist)
	}

	if count := len(res.Album); count != 50 {
		t.Fatalf("Got %d artist top albums, wanted 50\n", count)
	}
}

func TestArtistGetTopTags(t *testing.T) {
	res, err := client.ArtistGetTopTags("cher")

	if err != nil {
		t.Error(err)
	}

	if res.Attr.Artist != "Cher" {
		t.Fatalf("artist.gettoptags returned the wrong artist. Expected: %s Received: %s", "Cher", res.Attr.Artist)
	}

	if count := len(res.Tag); count < 45 {
		t.Fatalf("Got %d artist top tags, wanted 45\n", count)
	}
}

func TestArtistGetTopTracks(t *testing.T) {
	res, err := client.ArtistGetTopTracks("cher")

	if err != nil {
		t.Error(err)
	}

	if res.Attr.Artist != "Cher" {
		t.Fatalf("artist.gettoptracks returned the wrong artist. Expected: %s Received: %s", "Cher", res.Attr.Artist)
	}

	if count := len(res.Track); count != 50 {
		t.Fatalf("Got %d artist top tracks, wanted 50\n", count)
	}
}

func TestArtistSearch(t *testing.T) {
	res, err := client.ArtistSearch("cher")

	if err != nil {
		t.Error(err)
	}

	if res.Attr.For != "cher" {
		t.Fatalf("artist.search returned the wrong artist. Expected: %s Received: %s", "cher", res.Attr.For)
	}

	if count := len(res.ArtistMatches.Artist); count < 30 {
		t.Fatalf("Got %d artist matches, wanted 30\n", count)
	}
}
