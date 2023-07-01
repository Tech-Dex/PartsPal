package utils

import "net/http"

var Headers = &http.Header{
	"Accept":     {"text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8"},
	"User-Agent": {"Mozilla/5.0 (Windows NT 10.0; rv:114.0) Gecko/20100101 Firefox/114.0"},
	"Connection": {"keep-alive"},
}

var Client = &http.Client{}

func HttpGet(url string) (*http.Response, error) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header = *Headers
	return Client.Do(req)
}
