package datagen

// Case[T any] represents a case in a switch
type Case[T any] struct {
	vCk *ValCk
	v   TypedGenerator[T]
}

// NewCase returns a new Case of type T
func NewCase[T any](fc *ValCk, v TypedGenerator[T]) *Case[T] {
	return &Case[T]{
		vCk: fc,
		v:   v,
	}
}
