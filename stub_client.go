package testutils

import (
	"io"
	"net/http"
)

type StubResponse struct {
	Resp *http.Response
	Err  error
}

type StubClient struct {
	count     int
	Responses []*StubResponse
}

func NewStubClient(resps ...*StubResponse) *StubClient {
	return &StubClient{
		Responses: resps,
	}
}

func (c *StubClient) Do(request *http.Request) (*http.Response, error) {
	return c.next()
}

func (c *StubClient) Get(url string) (*http.Response, error) {
	return c.next()
}

func (c *StubClient) Head(url string) (*http.Response, error) {
	return c.next()
}

func (c *StubClient) Post(url, contentType string, body io.Reader) (*http.Response, error) {
	return c.next()
}

func (c *StubClient) next() (*http.Response, error) {
	// Yes, we want to panic if we have an array out of bounds.
	r := c.Responses[c.count]
	c.count += 1
	return r.Resp, r.Err
}
