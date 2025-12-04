package main

import (
	"fmt"
	"net/url"
)

func main() {

	// [scheme://][userinfo@] host [:port] [/path] [?query] [#fragment]

	rawUrl := "postgres://user:pass@host.com:5432/path?k=v#fragment"

	parsedURL, err := url.Parse(rawUrl)

	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	fmt.Println(parsedURL)

	fmt.Println("scheme:", parsedURL.Scheme)
	fmt.Println("host:", parsedURL.Host)
	fmt.Println("port:", parsedURL.Port())
	fmt.Println("path:", parsedURL.Path)
	fmt.Println("query:", parsedURL.RawQuery)
	fmt.Println("fragment:", parsedURL.Fragment)

	rawUrl2 := "https://example.com/path?name=akshat&age=24"
	parsedURL2, err := url.Parse(rawUrl2)

	if err !=  nil {
		return
	}
	fmt.Println("query", parsedURL2.Query())

	queryParams := parsedURL2.Query()

	fmt.Println( queryParams.Get("age"))

}
