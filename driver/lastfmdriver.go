package driver

import (
	"log"
	"os"
)

var (
	LAST_FM_KEY = os.Getenv("LAST_FM_KEY")
)

func init() {
	if LAST_FM_KEY == "" {
		log.Fatal("Please set the LAST_FM_KEY environment variable to your Last.fm API key!")
	}
}