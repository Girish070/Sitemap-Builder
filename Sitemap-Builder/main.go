package main

import (
	"flag"
	"fmt"
	"github.com/peterhellberg/link"
	"io"
	"net/http"
)

func main() {
	urlFlag := flag.String("url", "https://www.facebook.com", "The Url that you want to build a siteMap for")
	flag.Parse()

	fmt.Println(*urlFlag)
	resp, err := http.Get(*urlFlag)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	links,_:= link.Parse(resp.Body)
	for _, l := range links{
		fmt.Println(l)
	}

}