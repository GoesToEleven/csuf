package main

import (
	"fmt"
	"net/url"
)

func main() {
	baseURL := "http://api.hostip.info/get_json.php?position=true&ip="
	ip := "198 252 210 32"

	// QueryEscape escapes the ip string so
	// it can be safely placed inside a URL query
	safeIp := url.QueryEscape(ip)
	url := baseURL + safeIp

	fmt.Println(url)
}
