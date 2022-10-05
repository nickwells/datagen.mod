package datagen

// Generator describes the methods that a field generator must provide.
type Generator interface {
	Generate() string
	Next()
}

// TypedVal represents a method that returns the internal value of the field
// rather than the string representation.
type TypedVal[T any] interface {
	Value() T
}

// TypedGenerator adds the TypedVal interface to a Generator.
type TypedGenerator[T any] interface {
	Generator
	TypedVal[T]
}
