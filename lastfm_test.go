package lastfm

import (
	"net/http"
	"os"
	"reflect"
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
