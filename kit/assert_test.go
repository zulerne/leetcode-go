package kit

import (
	"testing"
)

func TestAssertEqual(t *testing.T) {
	t.Run("equal integers", func(t *testing.T) {
		AssertEqual(t, 5, 5)
	})

	t.Run("equal slices", func(t *testing.T) {
		AssertEqual(t, []int{1, 2, 3}, []int{1, 2, 3})
	})

	t.Run("equal strings", func(t *testing.T) {
		AssertEqual(t, "hello", "hello")
	})
}

func TestAssertTrue(t *testing.T) {
	t.Run("condition is true", func(t *testing.T) {
		AssertTrue(t, true)
	})

	t.Run("with message", func(t *testing.T) {
		condition := true
		AssertTrue(t, condition, "math should work")
	})
}

func TestAssertFalse(t *testing.T) {
	t.Run("condition is false", func(t *testing.T) {
		AssertFalse(t, false)
	})
}

func TestAssertNil(t *testing.T) {
	t.Run("nil slice", func(t *testing.T) {
		var s []int
		AssertNil(t, s)
	})

	t.Run("nil pointer", func(t *testing.T) {
		var p *int
		AssertNil(t, p)
	})
}

func TestAssertNotNil(t *testing.T) {
	t.Run("non-nil slice", func(t *testing.T) {
		s := []int{1, 2, 3}
		AssertNotNil(t, s)
	})

	t.Run("non-nil pointer", func(t *testing.T) {
		v := 42
		AssertNotNil(t, &v)
	})
}
