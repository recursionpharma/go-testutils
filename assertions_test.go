package testutils

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestShouldHaveErrorMessageWithSubstring(t *testing.T) {
	Convey("Given actual and expected", t, func() {
		Convey("If expected has more than 1 argument, an message should be returned", func() {
			So(ShouldHaveErrorMessageWithSubstring(fmt.Errorf("foo"), "bar", "baz"), ShouldNotBeBlank)
		})
		Convey("If actual isn't an error, an message should be returned", func() {
			So(ShouldHaveErrorMessageWithSubstring("foo", "bar"), ShouldNotBeBlank)
		})
		Convey("If expected isn't a string, an message should be returned", func() {
			So(ShouldHaveErrorMessageWithSubstring(fmt.Errorf("foo"), 123), ShouldNotBeBlank)
		})
		Convey("If the error is nil, an message should be returned", func() {
			var err error
			So(ShouldHaveErrorMessageWithSubstring(err, "foobar"), ShouldNotBeBlank)
		})
		Convey("If the expected message is empty, an message should be returned", func() {
			So(ShouldHaveErrorMessageWithSubstring(fmt.Errorf("foo"), ""), ShouldNotBeBlank)
		})
		Convey("If the actual error message doesn't contain the substring, an message should be returned", func() {
			So(ShouldHaveErrorMessageWithSubstring(fmt.Errorf("foo"), "bar"), ShouldNotBeBlank)
		})
		Convey("If the actual error message does contain the substring, an empty message should be returned", func() {
			So(ShouldHaveErrorMessageWithSubstring(fmt.Errorf("foobar"), "bar"), ShouldBeBlank)
		})
	})
}
