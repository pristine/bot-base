package hclient

import (
	"io"
	"net/http"
)

type Request struct {
	client *Client

	method, url, host string

	header http.Header

	body io.Reader

	cookies []*http.Cookie
}

func newRequest(c *Client) *Request {
	return &Request{
		client: c,
		header: make(http.Header),
	}
}

// SetURL sets the url of the request
func (r *Request) SetURL(url string) *Request {
	r.url = url
	return r
}

// SetMethod sets the method of the request
func (r *Request) SetMethod(method string) *Request {
	r.method = method
	return r
}

// AddHeader adds a specified header to the request
// If the header already exists, the value will be appended by the new specified value
// If the header does not exist, the header will be set to the specified value
func (r *Request) AddHeader(key, value string) *Request {
	if header, ok := r.header[key]; ok {
		header = append(header, value)
		r.header[key] = header
	} else {
		r.header[key] = []string{value}
	}
	return r
}

// SetHeader sets a specified header to the request
// This overrides any previously set values of the specified header
func (r *Request) SetHeader(key, value string) *Request {
	r.header[key] = []string{value}
	return r
}

// SetHost sets the host of the request
func (r *Request) SetHost(value string) *Request {
	r.host = value
	return r
}

// Do will send the request with all specified request values
func (r *Request) Do() (*Response, error) {
	req, err := http.NewRequest(r.method, r.url, r.body)
	if err != nil {
		return nil, err
	}

	for _, cookie := range r.cookies {
		if cookie != nil {
			req.AddCookie(cookie)
		}
	}

	req.Header = r.header

	if len(r.host) > 0 {
		req.Host = r.host
	}

	if len(r.header.Get("Host")) > 0 {
		req.Host = r.header.Get("Host")

		r.header.Del("Host")
	}

	return r.client.Do(req)
}
