Simple Golang library for making http REST requests slightly easier. In particular, it supports parsing JSON responses into Go structs.

To install:

```
go get github.com/peterjliu/rest
```


For example, 
```
package main

import (
	"fmt"

	"github.com/peterjliu/rest"
)

type UrlShortenerResp struct {
	LongUrl string `json:"longUrl"`
	Id      string `json: "id"`
	Status  string `json: "status"`
}

func main() {
	// Get short url for long url
	var answer UrlShortenerResp
	url := "https://www.googleapis.com/urlshortener/v1/url?shortUrl=http://goo.gl/fbsS"
	err := rest.Get(url, nil, &answer)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("%+v\n", answer)
}
```

Loot at more examples in `examples/`
