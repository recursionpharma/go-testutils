# go-testutils

Golang helpers for testing.

*NOTE* This is a public repo.

## Collection Assertions

This repository includes several new assertions: `Any`, `All`, `None`, `AtLeast`, `AtMost`, and `Exactly`. These can be wrapped around existing assertions to perform the assertion on every element of a collection.

Example:

```golang
So([]int{1, 2, 3, 4}, Any(ShouldEqual), 4)
So([]int{1, 2, 3, 4}, Exactly(2, ShouldBeLessThan), 3)
```
This lets you perform an assertion on a collection without making assumptions regarding the order of the elements. For example:

```golang
So(printings, Any(ShouldHaveId), id)
```

## Other Assertions

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

## This repo

    go-testutils/
    |-- assertions/
    |   contains custom assertions that work with goconvey
    |-- .editorconfig
    |   configures default for displaying files (e.g. in github PRs)
    |-- .gitignore
    |   git ignored files
    |-- Gopkg.lock
    |   go dep file
    |-- Gopkg.toml
    |   go dep file
    |-- README.md
    |   this file
    `-- .travis.yml
        travis configuration

The above file tree was generated with `tree -a -F -L 1 --charset ascii`.
