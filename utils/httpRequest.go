package myutils

import (
	"fmt"
	"io"
	"net/http"
)

// take url and return bytes data and error
func HttpGetRequest(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Print("Some error happend while loading data from server")
		return nil, err
	}
	data, parseErr := io.ReadAll(res.Body)
	if parseErr != nil {
		fmt.Print("error happend while reading...")
		return nil, parseErr
	}
	return data, nil
}
