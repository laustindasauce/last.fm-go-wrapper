// Copyright Austin Spencer
// All Rights Reserved

// lastfm package is used to communicate with the scrobbler API on Last.fm
//
// Using docs from https://www.last.fm/api/
package lastfm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Client is a client for working with the Last.fm Web API
// **Requires API key and secret for now, until auth.go is set up
type Client struct {
	http       *http.Client
	key        string
	secret     string
	baseURL    string
	baseApiURL string
	User       string
	jsonOpt    PathOptions
	keyOpt     PathOptions
	secretOpt  PathOptions
}

// PathOptions is the opt for URL parameters
type PathOptions struct {
	key   string
	value string
}

// New returns a client for working with the Spotify Web API.
// The provided httpClient must provide Authentication with the requests.
// The auth package may be used to generate a suitable client.
func New(httpClient *http.Client, key string, secret string) *Client {
	c := &Client{
		http: httpClient,
		// Base for Last.fm API endpoints
		baseURL:    "https://ws.audioscrobbler.com/2.0/?",
		baseApiURL: fmt.Sprintf("https://ws.audioscrobbler.com/2.0/?api_key=%s&format=%s", key, "json"),
		key:        key,
		secret:     secret,
		User:       "",
		jsonOpt:    PathOptions{"format", "json"},
		keyOpt:     PathOptions{"api_key", key},
		secretOpt:  PathOptions{"api_secret", secret},
	}

	if c.key == "" || c.secret == "" {
		log.Fatal("API key or API secret key are missing!")
	}

	return c
}

func (c *Client) SetUser(user string) error {
	// http://ws.audioscrobbler.com/2.0/?method=user.getinfo&user=rj&api_key=YOUR_API_KEY&format=json

	lastfmURL := fmt.Sprintf("%s&method=user.getinfo&user=%s", c.baseApiURL, user)

	var userInfo struct {
		User User `json:"user"`
	}

	err := c.get(lastfmURL, &userInfo)

	if err != nil {
		return err
	}

	// Set the user for the client
	c.User = userInfo.User.Name
	return nil
}

// Image identifies an image associated with an item
type Image struct {
	URL  string `json:"#text"`
	Size string `json:"size"`
}

// RecentDate
type RecentDate struct {
	UTS  string `json:"uts"`
	Text string `json:"#text"`
}

type Wiki struct {
	Published string `json:"published"`
	Summary   string `json:"summary"`
	Content   string `json:"content"`
}

type Bio struct {
	Links     Links  `json:"links"`
	Published string `json:"published"`
	Summary   string `json:"summary"`
	Content   string `json:"content"`
}

type Links struct {
	Link Link `json:"link"`
}

type Link struct {
	Text string `json:"#text"`
	Rel  string `json:"rel"`
	Href string `json:"href"`
}

type Rank struct {
	Rank string `json:"rank"`
}

// Tags identifies an array of type tag
type Tags struct {
	Tag []Tag `json:"tag"`
}

type TagsWithCount struct {
	Tag  []TagWithCount `json:"tag"`
	Attr ArtistTagAttr  `json:"@attr"`
}

// Tag identifies a tag associated with an item
type Tag struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PosAttr struct {
	Position string `json:"position"`
}

// TagWithCount identifies a tag with count
type TagWithCount struct {
	Count int    `json:"count"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}

type OpenSearchQuery struct {
	Text        string `json:"#text"`
	Role        string `json:"role"`
	SearchTerms string `json:"searchTerms"`
	StartPage   string `json:"startPage"`
}

type SearchAttr struct {
	For string `json:"for"`
}

// Error represents an error returned by the Last.fm Web API.
type Error struct {
	// A short description of the error.
	Message string `json:"message"`
	// The HTTP status code.
	Err int `json:"error"`
}

func (e Error) Error() string {
	return e.Message
}

// decodeError decodes an Error from an io.Reader.
func (c *Client) decodeError(resp *http.Response) error {
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if len(responseBody) == 0 {
		return fmt.Errorf("lastfm: HTTP %d: %s (body empty)", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	buf := bytes.NewBuffer(responseBody)

	var e Error
	err = json.NewDecoder(buf).Decode(&e)
	if err != nil {
		return fmt.Errorf("lastfm: couldn't decode error: (%d) [%s]", len(responseBody), responseBody)
	}

	if e.Message == "" {
		// Some errors will result in there being a useful status-code but an
		// empty message, which will confuse the user (who only has access to
		// the message and not the code). An example of this is when we send
		// some of the arguments directly in the HTTP query and the URL ends-up
		// being too long.

		e.Message = fmt.Sprintf("lastfm: unexpected HTTP %d: %s (empty error)",
			resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	return e
}

// get handles the get requests for the client
func (c *Client) get(url string, result interface{}) error {
	for {
		resp, err := http.Get(url)

		if err != nil {
			return err
		}

		// body, err := ioutil.ReadAll(resp.Body)

		// if err != nil {
		// 	return err
		// }

		// fmt.Println(string(body))

		defer resp.Body.Close()

		if resp.StatusCode == http.StatusNoContent {
			return nil
		}
		if resp.StatusCode != http.StatusOK {
			return c.decodeError(resp)
		}

		err = json.NewDecoder(resp.Body).Decode(result)
		if err != nil {
			return err
		}

		break
	}

	return nil
}
