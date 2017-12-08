# go-testutils

Golang helpers for testing.

## Assertions

Contains custom assertions that work with [goconvey](https://github.com/smartystreets/goconvey).

For example, if we have some code that returns an error:

    package thehulk

    import (
        "fmt"
    )

    type Hulk struct {
        isHulked boolean
    }

    func (h *Hulk) Hulkify(angerLevel int) error {
        if angerLevel < 3 {
            return fmt.Errorf("Not angry enough: %d", angerLevel)
        }
        h.isHulked = true
        return nil
    }

We might test it like this:

    package thehulk_test

    import (
        "fmt"
        "testing"

        . "github.com/smartystreets/goconvey/convey"
        . "github.com/recursionpharma/go-testutils/assertions"
    )

    func TestHulkify (t *testing.T) {
        t.Parallel()

        Convey("Given an anger level", t, func() {

            Convey("If the anger is too low, an error should be returned", func() {
                angerLevel := 2
                err := h.Hulkify(angerLevel)
                So(err, ShouldHaveErrorMessageWithSubstring, fmt.Sprintf("%d", angerLevel))
            })

            Convey("If the anger is high enough, no error should be returned", func() {
                err := h.Hulkify(5)
                So(err, ShouldBeNil)
            })
        })
    }


## HTTP Client

A stub HTTP client for use in testing.

Suppose we have some code that makes calls to the platelet API:

    package platelet

    import (
        "net/http"

        "github.com/recursionphamrma/go-httpclient"
    )

    type Client struct {
        URL string
        client httpclient.Client
    }

    func (c *Client) Get(path string) (*http.Response, error) {
        return c.client.Get(c.URL + "/" + path)
    }

Then our test might look like:

    package platelet_test

    import (
        "fmt"
        "http"
        "testing"

        . "github.com/smartystreets/goconvey/convey"
        . "github.com/recursionpharma/go-testutils/httptest"
    )

    func TestHTTPGetter(t *testing.T) {
        t.Parallel()

        Convey("Given a path", t, func() {

            Convey("If there's an error, an error should be returned" func() {
                client := &Client{
                    client: NewStubClient(&StubResponse{Err: fmt.Errorf("Oh noes!"}),
                }
                resp, err := client.Get("some-path")
                So(resp, ShouldBeNil)
                So(err, ShouldNotBeNil)
            })

            Convey("If there's no error, a response should be returned" func() {
                client := &Client{
                    client: NewStubClient(&StubResponse{Resp: http.Response{StatusCode: http.StatusOK}}),
                }
                resp, err := client.Get("some-path")
                So(resp, ShouldNotBeNil)
                So(err, ShouldBeNil)
            })
        })
    }

## This repo

    go-testutils/
    |-- assertions
    |   contains custom assertions that work with goconvey
    |-- .editorconfig
    |   configures default for displaying files (e.g. in github PRs)
    |-- .gitignore
    |   git ignored files
    |-- Gopkg.lock
    |   go dep file
    |-- Gopkg.toml
    |   go dep file
    |-- httptest
    |   utilities for testing http
    |-- README.md
    |   this file
    `-- .travis.yml
        travis configurations

The above file tree was generated with `tree -a -L 1 --charset ascii`.
