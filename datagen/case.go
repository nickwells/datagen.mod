package datagen

// Case represents a case in a switch
type Case[T any] struct {
	p Passer
	v TypedGenerator[T]
}

// NewCase returns a new Case of type T
func NewCase[T any](p Passer, v TypedGenerator[T]) *Case[T] {
	return &Case[T]{
		p: p,
		v: v,
	}
}
