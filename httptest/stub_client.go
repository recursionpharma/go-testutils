package httptest

import (
	"io"
	"net/http"
)

// StubResponse stubs the return value of the functions in the github.com/recursionpharma/go-httpclient Client interface.
type StubResponse struct {
	Resp *http.Response
	Err  error
}

// StubClient impelments the github.com/recursionpharma/go-httpclient Client interface.
// It returns each response in the order of the Responses slice.
type StubClient struct {
	count     int
	Responses []*StubResponse
}

// NewStubClient creates a new StubClient with the passed responses.
func NewStubClient(resps ...*StubResponse) *StubClient {
	return &StubClient{
		Responses: resps,
	}
}

// Do returns the next response in the StubClient.
func (c *StubClient) Do(request *http.Request) (*http.Response, error) {
	return c.next()
}

// Get returns the next response in the StubClient.
func (c *StubClient) Get(url string) (*http.Response, error) {
	return c.next()
}

// Head returns the next response in the StubClient.
func (c *StubClient) Head(url string) (*http.Response, error) {
	return c.next()
}

// Post returns the next response in the StubClient.
func (c *StubClient) Post(url, contentType string, body io.Reader) (*http.Response, error) {
	return c.next()
}

func (c *StubClient) next() (*http.Response, error) {
	// Yes, we want to panic if we have an array out of bounds.
	r := c.Responses[c.count]
	c.count++
	return r.Resp, r.Err
}
