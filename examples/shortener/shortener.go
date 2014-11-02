package main

import (
	"fmt"

	"github.com/peterjliu/rest"
)

type UrlShortenerReq struct {
	LongUrl string `json:"longUrl"`
}

type UrlShortenerResp struct {
	LongUrl string `json:"longUrl"`
	Id      string `json: "id"`
	// if asked to expand a shortUrl, this is the status of it
	Status string `json: "status,omitempty"`
}

func main() {
	// Get short url for long url
	var answer UrlShortenerResp
	err := rest.Post(
		"https://www.googleapis.com/urlshortener/v1/url",
		nil,
		UrlShortenerReq{LongUrl: "http://www.google.com"},
		&answer)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println("Shorten URL response:")
	fmt.Printf("%+v\n\n", answer)

	// Get long url for a short url
	answer = UrlShortenerResp{}
	err = rest.Get(
		"https://www.googleapis.com/urlshortener/v1/url?shortUrl=http://goo.gl/fbsS",
		nil,
		&answer)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println("Expand URL response:")
	fmt.Printf("%+v\n\n", answer)

	// Do a GET by defining a Request
	answer = UrlShortenerResp{}
	r := rest.Request{
		Method: rest.GET,
		Url:    "https://www.googleapis.com/urlshortener/v1/url?shortUrl=http://goo.gl/fbsS",
	}
	err = r.Do(&answer)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println("Expand URL response, using Request:")
	fmt.Printf("%+v\n\n", answer)
}
