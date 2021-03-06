// Package optional contains utilities for working with possibly-empty values.
package optional

// An Optional is a wrapper around a possibly-present value.
type Optional[T any] struct {
	present bool
	val     T
}

// OrElse returns the value of the optional, if it is present, or the provided
// fallback if the optional is empty.
func (o Optional[T]) OrElse(fallback T) T {
	if o.present {
		return o.val
	}
	return fallback
}

// Present returns true if and only if the optional contains a value.
func (o Optional[T]) Present() bool {
	return o.present
}

// Value returns the optional's value, or the zero value for the optional type,
// and a boolean indicating whether the value was present or not.
func (o Optional[T]) Value() (T, bool) {
	return o.val, o.present
}

// Of returns an Optional representing the provided present value.
//
// To create an optional from a possibly-nil pointer, use OfNillable instead.
func Of[T any](val T) *Optional[T] {
	return &Optional[T]{val: val, present: true}
}

// OfNillable returns an Optional representing the provided value. If the value
// is nil, then the resulting optional will be empty. Otherwise, the value of
// the resulting optional will be *val.
func OfNillable[T any](val *T) Optional[T] {
	if val == nil {
		return Optional[T]{}
	}
	return Optional[T]{val: *val, present: true}
}

// Empty returns an empty optional of the specified type.
func Empty[T any]() Optional[T] {
	return Optional[T]{}
}
