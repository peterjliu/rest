package main

import (
	"fmt"
	"net/url"

	"github.com/peterjliu/rest"
)

type UrlShortenerReq struct {
	LongUrl string `json:"longUrl"`
}

type UrlShortenerResp struct {
	LongUrl string `json:"longUrl"`
	Id      string `json: "id"`
	// if asked to expand a shortUrl, this is the status of it
	Status string `json: "status"`
}

func main() {
	// Get short url for long url
	var answer UrlShortenerResp
	err := rest.Post("https://www.googleapis.com/urlshortener/v1/url",
		UrlShortenerReq{LongUrl: "http://www.googlsse.com"}, &answer)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println("Shorten URL response:")
	fmt.Printf("%+v\n", answer)

	// Get long url for a short url
	v := url.Values{
		"shortUrl": {"http://goo.gl/OYEBE1"},
	}
	u := url.URL{
		Scheme:   "https",
		Host:     "www.googleapis.com",
		Path:     "urlshortener/v1/url",
		RawQuery: v.Encode(),
	}
	fmt.Println("\nExpand URL response:")
	err = rest.Get(u.String(), &answer)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("%+v\n", answer)

	fmt.Println("\nExpand URL response without building URL:")
	err = rest.Get("https://www.googleapis.com/urlshortener/v1/url?shortUrl=http://goo.gl/fbsS", &answer)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("%+v\n", answer)

}
