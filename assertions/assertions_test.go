package assertions_test

import (
	"fmt"
	"testing"

	. "github.com/recursionpharma/go-testutils/assertions"
	. "github.com/smartystreets/goconvey/convey"
)

func TestAtLeast(t *testing.T) {
	Convey("Given a slice", t, func() {
		testSlice := []int{1, 2, 3, 4}

		Convey("It ensures that at least N elements pass an assertion", func() {
			So(testSlice, AtLeast(0, ShouldBeGreaterThan), 2)
			So(testSlice, AtLeast(1, ShouldBeGreaterThan), 2)
			So(testSlice, AtLeast(2, ShouldBeGreaterThan), 2)
			failureMessage := AtLeast(3, ShouldBeGreaterThan)(testSlice, 2)
			So(failureMessage, ShouldEqual, "Expected the collection (length 4) to contain at least 3 passing elements, but it contained 2.\nFailures:\n\nExpected '1' to be greater than '2' (but it wasn't)!\nExpected '2' to be greater than '2' (but it wasn't)!")
			failureMessage = AtLeast(1, ShouldBeGreaterThan)(testSlice, 5)
			So(failureMessage, ShouldEqual, "Expected the collection (length 4) to contain at least 1 passing element, but it contained 0.\nFailures:\n\nExpected '1' to be greater than '5' (but it wasn't)!\nExpected '2' to be greater than '5' (but it wasn't)!\nExpected '3' to be greater than '5' (but it wasn't)!\nExpected '4' to be greater than '5' (but it wasn't)!")
		})
	})
}

func TestAtMost(t *testing.T) {
	Convey("Given a slice", t, func() {
		testSlice := []int{1, 2, 3, 4}

		Convey("It ensures that at most N elements pass an assertion", func() {
			So(testSlice, AtMost(4, ShouldBeGreaterThan), 2)
			So(testSlice, AtMost(3, ShouldBeGreaterThan), 2)
			So(testSlice, AtMost(2, ShouldBeGreaterThan), 2)
			failureMessage := AtMost(1, ShouldBeGreaterThan)(testSlice, 2)
			So(failureMessage, ShouldEqual, "Expected the collection (length 4) to contain at most 1 passing element, but it contained 2.")
			failureMessage = AtMost(0, ShouldBeGreaterThan)(testSlice, 2)
			So(failureMessage, ShouldEqual, "Expected the collection (length 4) to contain at most 0 passing elements, but it contained 2.")
		})
	})
}

func TestExactly(t *testing.T) {
	Convey("Given a slice", t, func() {
		testSlice := []int{1, 2, 3, 4}

		Convey("It ensures that exactly N elements pass an assertion", func() {
			So(testSlice, Exactly(2, ShouldBeLessThan), 3)
			So(testSlice, Exactly(1, ShouldBeLessThan), 2)
			So(testSlice, Exactly(0, ShouldBeGreaterThan), 4)
			failureMessage := Exactly(2, ShouldBeLessThan)(testSlice, 2)
			So(failureMessage, ShouldEqual, "Expected the collection (length 4) to contain exactly 2 passing elements, but it contained 1.\nFailures:\n\nExpected '2' to be less than '2' (but it wasn't)!\nExpected '3' to be less than '2' (but it wasn't)!\nExpected '4' to be less than '2' (but it wasn't)!")
			failureMessage = Exactly(2, ShouldBeLessThan)(testSlice, 4)
			So(failureMessage, ShouldEqual, "Expected the collection (length 4) to contain exactly 2 passing elements, but it contained 3.")
		})
	})
}

func TestAny(t *testing.T) {
	Convey("Given a slice", t, func() {
		testSlice := []int{1, 2, 3, 4}

		Convey("It ensures that at least 1 element passes an assertion", func() {
			So(testSlice, Any(ShouldEqual), 3)
			So(testSlice, Any(ShouldBeLessThan), 3)
			failureMessage := Any(ShouldEqual)(testSlice, 5)
			expectedMessage := `Expected the collection (length 4) to contain at least 1 passing element, but it contained 0.
Failures:

Expected: '5'
Actual:   '1'
(Should be equal)
Expected: '5'
Actual:   '2'
(Should be equal)
Expected: '5'
Actual:   '3'
(Should be equal)
Expected: '5'
Actual:   '4'
(Should be equal)`
			So(failureMessage, ShouldEqual, expectedMessage)
		})
	})
}

func TestAll(t *testing.T) {
	Convey("Given a slice", t, func() {
		testSlice := []int{1, 2, 3, 4}

		Convey("It ensures that all elements pass an assertion", func() {
			So(testSlice, All(ShouldBeGreaterThan), 0)
			failureMessage := All(ShouldBeGreaterThan)(testSlice, 1)
			So(failureMessage, ShouldEqual, "Expected the collection (length 4) to contain exactly 4 passing elements, but it contained 3.\nFailures:\n\nExpected '1' to be greater than '1' (but it wasn't)!")
		})
	})
}

func TestNone(t *testing.T) {
	Convey("Given a slice", t, func() {
		testSlice := []int{1, 2, 3, 4}

		Convey("It ensures that no elements pass an assertion", func() {
			So(testSlice, None(ShouldEqual), 5)
			failureMessage := None(ShouldEqual)(testSlice, 4)
			So(failureMessage, ShouldEqual, "Expected the collection (length 4) to contain exactly 0 passing elements, but it contained 1.")
		})
	})
}

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

func TestJoinComparisons(t *testing.T) {
	Convey("Given a struct with properties", t, func() {

		type testStruct struct {
			Name string
			Id   int
		}

		testSlice := []testStruct{
			testStruct{
				Name: "Alice",
				Id:   1,
			},
			testStruct{
				Name: "Bob",
				Id:   2,
			},
		}

		Convey("Properties can be tested dependently", func() {
			So(testSlice, ShouldHaveLength, 2)
			So(testSlice, Exactly(1, func(actual interface{}, expected ...interface{}) string {
				return JoinComparisons([]string{
					ShouldEqual(actual.(testStruct).Name, "Alice"),
					ShouldEqual(actual.(testStruct).Id, 1),
				})
			}))
			failureMessage := Exactly(2, func(actual interface{}, expected ...interface{}) string {
				return JoinComparisons([]string{
					ShouldEqual(actual.(testStruct).Name, "Bob"),
					ShouldEqual(actual.(testStruct).Id, 2),
				})
			})(testSlice)
			So(failureMessage, ShouldEqual, "Expected the collection (length 2) to contain exactly 2 passing elements, but it contained 1.\nFailures:\n\nExpected: 'Bob'\nActual:   'Alice'\n(Should be equal)")
		})
	})
}
