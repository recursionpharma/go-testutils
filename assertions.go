package testutils

import (
	"fmt"

	. "github.com/smartystreets/goconvey/convey"
)

func ShouldHaveErrorMessageWithSubstring(actual interface{}, expected ...interface{}) string {
	if len(expected) != 1 {
		return fmt.Sprintf("Expected exactly 1 value for expected but got '%d'", len(expected))
	}

	err, valueIsError := actual.(error)
	msg, valueIsString := expected[0].(string)

	if !valueIsError {
		return fmt.Sprintf("Expected actual to be an error but got '%+v'", actual)
	}
	if !valueIsString {
		return fmt.Sprintf("Expected expected to be a string slice with exactly 1 element but got '%+v'", expected)
	}
	if err == nil {
		return fmt.Sprintf("Expected error to not be nil but got '%+v'", err)
	}
	if msg == "" {
		return fmt.Sprintf("Expected message to not be non-empty but got '%+v'", msg)
	}

	return ShouldContainSubstring(err.Error(), msg)
}
