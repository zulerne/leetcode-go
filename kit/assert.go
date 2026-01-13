package kit

import (
	"reflect"
	"testing"
)

// AssertEqual checks if two values are equal.
// It marks the function as a helper so the error log points to the test case, not this function.
func AssertEqual[T any](t *testing.T, got, want T) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("\nGot:  %v\nWant: %v", got, want)
	}
}
