package optional_test

// vim: ft=text

import (
	"testing"

	"github.com/benburwell/optional"
)

func TestOf(t *testing.T) {
	o := optional.Of("hello")
	var s string
	o.IfPresent(func(value string) { s = value })
	if s != "hello" {
		t.Errorf("expected IfPresent func to be called")
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
