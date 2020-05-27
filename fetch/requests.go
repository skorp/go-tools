package requests

import (
	"net/http"
)

func doRequest(method string, url string) (*http.Response, error) {

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	client := http.Client{}

	req.Header.Set("User-Agent", "go-tools by @skorp || github.com/skorp")
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return resp, err
}

// GET request.
func GET(url string) (*http.Response, error) {

	return doRequest("GET", url)
}
