package optional

// vim: ft=text

type Optional[T any] struct {
	val *T
}

func (o *Optional[T]) IfPresent(f func(x T)) {
	if o.val != nil {
		f(*o.val)
	}
}

func (o *Optional[T]) OrElse(fallback T) T {
	if o.val != nil {
		return *o.val
	}
	return fallback
}

func (o *Optional[T]) Present() bool {
	return o.val != nil
}

func Of[T any](val T) *Optional[T] {
	return &Optional[T]{val: &val}
}

func OfNillable[T any](val *T) *Optional[T] {
	return &Optional[T]{val: val}
}
