package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	lastfm "github.com/austinbspencer/last.fm-go-wrapper"
)

func main() {
	hClient := http.Client{Timeout: time.Duration(1) * time.Second}
	client := lastfm.New(&hClient, os.Getenv("LAST_FM_KEY"), os.Getenv("LAST_FM_SECRET"))

	err := client.SetUser("abspen1")

	if err != nil {
		log.Fatal(err)
	}

	res, err := client.UserGetPersonalAlbumTags("love", "", "")

	// res, err := client.TrackGetTags("Hells Bells", "AC/DC", "")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
