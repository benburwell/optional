package optional_test

import (
	"testing"

	"github.com/benburwell/optional"
)

func TestOf(t *testing.T) {
	o := optional.Of("hello")
	if !o.Present() {
		t.Errorf("expected value to be present, but was not")
	}
	val, ok := o.Value()
	if !ok {
		t.Errorf("expected getting value to be OK, but none returned")
	}
	if val != "hello" {
		t.Errorf("expected value to be 'hello', but got %q", val)
	}
}

func TestOrElse(t *testing.T) {
	t.Run("present", func(t *testing.T) {
		actual := optional.Of("a").OrElse("b")
		if actual != "a" {
			t.Errorf("expected 'a', but got: %s", actual)
		}
	})
	t.Run("absent", func(t *testing.T) {
		var s *string
		actual := optional.OfNillable(s).OrElse("b")
		if actual != "b" {
			t.Errorf("expected 'a', but got: %s", actual)
		}
	})
}

func TestOfNillable(t *testing.T) {
	var s *string
	if optional.OfNillable(s).Present() {
		t.Errorf("nil should not be present")
	}
}

func TestEmpty(t *testing.T) {
	if optional.Empty[string]().Present() {
		t.Errorf("expected empty optional to not be present")
	}
}
