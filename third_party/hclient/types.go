package hclient

import (
	"io"
	"net/http"
)

type Client struct {
	client *http.Client

	LatestResponse *Response
}

type Request struct {
	client *Client

	method, url, host string

	header http.Header

	body io.Reader

	cookies []*http.Cookie
}

type Response struct {
	headers http.Header

	body []byte

	status     string
	statusCode int
}