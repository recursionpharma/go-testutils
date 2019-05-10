package assertions

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/smartystreets/goconvey/convey"
)

type Assertion = func(actual interface{}, expected ...interface{}) string

// TODO: Support other collection types
func countPassingElements(assertion Assertion, collection interface{}, expected ...interface{}) (length int, passingCounter int, failures []string, err error) {
	failures = []string{}
	switch reflect.TypeOf(collection).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(collection)
		length = s.Len()
		for i := 0; i < length; i++ {
			each := s.Index(i)
			message := assertion(each.Interface(), expected...)
			if message == "" {
				passingCounter++
			} else {
				// goconvey does not consistently structure failure message strings, some look like this:
				// Expected '2' to be less than '2' (but it wasn't)!
				// others look like this:
				// {"Message":"Expected: '5'\nActual:   '1'\n(Should be equal)","Expected":"5","Actual":"1"}
				type Failure struct {
					Message  string
					Expected string
					Actual   string
				}
				var f Failure
				err := json.Unmarshal([]byte(message), &f)
				if err == nil {
					failures = append(failures, f.Message)
				} else {
					failures = append(failures, message)
				}

			}
		}
	default:
		return 0, 0, []string{}, fmt.Errorf("The collection is not a valid slice.")
	}
	return length, passingCounter, failures, nil
}

func pluralize(num int, word string) string {
	if num == 1 {
		return word
	}
	return fmt.Sprintf("%ss", word)
}

func AtLeast(required int, subAssertion Assertion) Assertion {
	return func(collection interface{}, expected ...interface{}) string {
		total, passing, failures, err := countPassingElements(subAssertion, collection, expected...)
		if err != nil {
			return err.Error()
		}
		if passing < required {
			combinedMessage := fmt.Sprintf("Expected the collection (length %d) to contain at least %d passing %s, but it contained %d.", total, required, pluralize(required, "element"), passing)
			combinedMessage = combinedMessage + "\nFailures:\n"
			for _, message := range failures {
				combinedMessage = combinedMessage + "\n" + message
			}
			return combinedMessage
		}
		return ""
	}
}

func AtMost(required int, subAssertion Assertion) Assertion {
	return func(collection interface{}, expected ...interface{}) string {
		total, passing, _, err := countPassingElements(subAssertion, collection, expected...)
		if err != nil {
			return err.Error()
		}
		if passing > required {
			return fmt.Sprintf("Expected the collection (length %d) to contain at most %d passing %s, but it contained %d.", total, required, pluralize(required, "element"), passing)
		}
		return ""
	}
}

func exactly(required int, all bool, subAssertion Assertion) Assertion {
	return func(collection interface{}, expected ...interface{}) string {
		total, passing, failures, err := countPassingElements(subAssertion, collection, expected...)
		if err != nil {
			return err.Error()
		}
		if all {
			required = total
		}
		if passing != required {
			combinedMessage := fmt.Sprintf("Expected the collection (length %d) to contain exactly %d passing %s, but it contained %d.", total, required, pluralize(required, "element"), passing)
			if passing < required {
				combinedMessage = combinedMessage + "\nFailures:\n"
				for _, message := range failures {
					combinedMessage = combinedMessage + "\n" + message
				}
			}
			return combinedMessage
		}
		return ""
	}
}

func Any(subAssertion Assertion) Assertion {
	return AtLeast(1, subAssertion)
}

func All(subAssertion Assertion) Assertion {
	return exactly(0, true, subAssertion)
}

func Exactly(required int, subAssertion Assertion) Assertion {
	return exactly(required, false, subAssertion)
}

func None(subAssertion Assertion) Assertion {
	return Exactly(0, subAssertion)
}

// ShouldHaveErrorMessageWithSubstring takes an error and a message
// and checks that the error is not nil and contains the message.
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

	return convey.ShouldContainSubstring(err.Error(), msg)
}

func JoinComparisons(comparisons []string) string {
	for _, comparison := range comparisons {
		if comparison != "" {
			return comparison
		}
	}
	return ""
}
