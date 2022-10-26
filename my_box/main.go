package my_box

type Box[T any] struct {
	v T
}

func NewBox[T any](t T) Box[T] {
	return Box[T]{
		v: t,
	}
}

// Zero is not included in a tag, and cannot be used without
// a workspace (or a pseudo-version tag must be used).
func (b *Box[T]) Zero() *Box[T] {
	var t T
	b.v = t
	return b
}
