package utils

import (
	"fmt"
	"io"
	"net/http"
)

type Response struct {
	Result []byte
	Err    error
}

// take url and return bytes data and error
func HttpGetRequest(url string, data_channel chan Response) {

	resp := Response{nil, nil}
	// var abc string
	res, err := http.Get(url)
	if err != nil {
		resp.Err = err
	}
	data, parseErr := io.ReadAll(res.Body)

	if parseErr != nil {
		fmt.Print("error happend while reading...")
		resp.Err = parseErr
	}
	resp.Result = data
	data_channel <- resp
}
