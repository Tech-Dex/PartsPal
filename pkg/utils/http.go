package utils

import "net/http"

var Headers = &http.Header{
	"Accept":     {"text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8"},
	"User-Agent": {"Mozilla/5.0 (X11; Linux x86_64; rv:60.0) Gecko/20100101 Firefox/60.0 FirePHP/0.7.4"},
}

var Client = &http.Client{}

func HttpGet(url string) (*http.Response, error) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header = *Headers
	return Client.Do(req)
}
