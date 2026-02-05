package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"regexp"
)

func fetch(url string) string {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	r := regexp.MustCompile(`"browseId"\s*:\s*"(U[^"]+)"`)
	browseId := r.FindStringSubmatch(string(body))[1]

	return browseId
}

var urlFlag string

func init() {
	flag.StringVar(&urlFlag, "u", "", "url for a youtube channel")
	flag.StringVar(&urlFlag, "url", "", "url for a youtube channel")
}

func convert(browseId string) string {
	prefix := "https://www.youtube.com/feeds/videos.xml?channel_id="
	feed := prefix + browseId
	return feed
}

func main() {
	flag.Parse()

	log.Printf("YT RSS Feed: %s\n", convert(fetch(urlFlag)))
}
