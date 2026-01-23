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

// AssertTrue checks if the condition is true.
func AssertTrue(t *testing.T, condition bool, msg ...string) {
	t.Helper()

	if !condition {
		if len(msg) > 0 {
			t.Errorf("Expected true, but got false: %s", msg[0])
		} else {
			t.Error("Expected true, but got false")
		}
	}
}

// AssertFalse checks if the condition is false.
func AssertFalse(t *testing.T, condition bool, msg ...string) {
	t.Helper()

	if condition {
		if len(msg) > 0 {
			t.Errorf("Expected false, but got true: %s", msg[0])
		} else {
			t.Error("Expected false, but got true")
		}
	}
}

// AssertNil checks if the value is nil.
func AssertNil(t *testing.T, got any) {
	t.Helper()

	if got != nil && !reflect.ValueOf(got).IsNil() {
		t.Errorf("Expected nil, but got: %v", got)
	}
}

// AssertNotNil checks if the value is not nil.
func AssertNotNil(t *testing.T, got any) {
	t.Helper()

	if got == nil || reflect.ValueOf(got).IsNil() {
		t.Error("Expected not nil, but got nil")
	}
}
