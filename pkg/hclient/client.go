package hclient

import (
	"crypto/tls"
	"errors"
	"io"
	"net/http"
	"net/url"
)

var (
	NoCookieJarErr = errors.New("no cookie jar in client")
)

type Client struct {
	client *http.Client

	LastResponse *Response
}

// NewClient creates a new http client
// Takes in the optional arguments: proxy, servername
func NewClient( parameters ...string) (*Client, error) {
	tlsClientConfig := &tls.Config{
		InsecureSkipVerify: true,
	}

	// parameters[0] = proxy
	// parameters[1] = sni
	if len(parameters) > 1 && len(parameters[1]) > 0 {
		tlsClientConfig.ServerName = parameters[1]
	}

	transport := &http.Transport{
		ForceAttemptHTTP2: true,
		TLSClientConfig: tlsClientConfig,
	}

	if len(parameters) > 0 && len(parameters[0]) > 0 {
		proxyUrl, _ := url.Parse(parameters[0])

		transport.Proxy = http.ProxyURL(proxyUrl)
	}

	return &Client{
		client: &http.Client{
			Transport: transport,
		},
		LastResponse: &Response{},
	}, nil
}

// NewRequest creates a new request under a specified http client
func (c *Client) NewRequest() *Request {
	return newRequest(c)
}

// AddCookie adds a new cookie to the request client cookie jar
func (c *Client) AddCookie(u *url.URL, cookie *http.Cookie) error {
	if c.client.Jar == nil {
		return NoCookieJarErr
	}

	c.client.Jar.SetCookies(u, append(c.client.Jar.Cookies(u), cookie))

	return nil
}

// RemoveCookie removes the specified cookie from the request client cookie jar
func (c *Client) RemoveCookie(u *url.URL, cookie string) error {
	if c.client.Jar == nil {
		return NoCookieJarErr
	}

	newCookie := &http.Cookie{
		Name: cookie,
		Value: "",
	}

	c.client.Jar.SetCookies(u, append(c.client.Jar.Cookies(u), newCookie))

	return nil
}

// Do will send the specified request
func (c *Client) Do(r *http.Request) (*Response, error) {
	resp, err := c.client.Do(r)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	_ = resp.Body.Close()
	if err != nil {
		return nil, err
	}

	r.Close = true

	response :=  &Response{
		headers: resp.Header,
		body: body,
		status:     resp.Status,
		statusCode: resp.StatusCode,
	}

	c.LastResponse = response

	return response, nil
}