package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	lastfm "github.com/austinbspencer/last.fm-go-wrapper"
)

func testLastfmFunc(client *lastfm.Client) {
	res, err := client.GeoGetTopTracks("usa")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}

func main() {
	hClient := http.Client{Timeout: time.Duration(1) * time.Second}
	client := lastfm.New(&hClient, os.Getenv("LAST_FM_KEY"), os.Getenv("LAST_FM_SECRET"))

	// client.SetUser("RJ")
	res, err := client.TrackSearch("Believe", "Cher", "", "")

	// res, err := client.TrackGetTags("Hells Bells", "AC/DC", "")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
