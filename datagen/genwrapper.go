package datagen

import "errors"

// GWStringer represents an object having a String method
type GWStringer[T any] interface {
	String(TypedGenerator[T]) string
}

// GWIterator represents an object having a Next method
type GWIterator[T any] interface {
	Next(TypedGenerator[T])
}

// GWValue represents an object having a Value method
type GWValue[T any] interface {
	Value(TypedGenerator[T]) T
}

// dfltOperations just calls the methods on the underlying TypedGenerator
type dfltOperations[T any] struct{}

// String calls the String function on the underlying TypedGenerator
func (dfltOperations[T]) String(wrapped TypedGenerator[T]) string {
	return wrapped.String()
}

// Next[T any] calls the Next function on the underlying
// TypedGenerator
func (dfltOperations[T]) Next(wrapped TypedGenerator[T]) {
	wrapped.Next()
}

// Value[T any] calls the Value function on the underlying
// TypedGenerator
func (dfltOperations[T]) Value(wrapped TypedGenerator[T]) T {
	return wrapped.Value()
}

// GenWrapper wraps a TypedGenerator and allows the Next, String and Value
// methods to be overridden.
type GenWrapper[T any] struct {
	wrapped TypedGenerator[T]

	stringer GWStringer[T]
	value    GWValue[T]
	iterator GWIterator[T]
}

// GenWraperOptFunc is the type of a function that can be passed to
// NewGenWrapper to set various values.
type GenWraperOptFunc[T any] func(*GenWrapper[T]) error

// GenWrapperSetStringer sets the generator to the supplied value
func GenWrapperSetStringer[T any](g GWStringer[T]) GenWraperOptFunc[T] {
	return func(gw *GenWrapper[T]) error {
		if g == nil {
			return errors.New("the GWStringer cannot be nil")
		}

		gw.stringer = g

		return nil
	}
}

// GenWrapperSetIterator sets the next to the supplied value
func GenWrapperSetIterator[T any](n GWIterator[T]) GenWraperOptFunc[T] {
	return func(gw *GenWrapper[T]) error {
		if n == nil {
			return errors.New("the GWIterator cannot be nil")
		}

		gw.iterator = n

		return nil
	}
}

// GenWrapperSetValue sets the value to the supplied value
func GenWrapperSetValue[T any](v GWValue[T]) GenWraperOptFunc[T] {
	return func(gw *GenWrapper[T]) error {
		if v == nil {
			return errors.New("the GWValue cannot be nil")
		}

		gw.value = v

		return nil
	}
}

// NewGenWrapper returns a GenWrapper, a type which allows the String,
// Value and Next functions to be overridden. Note that this function must
// take at least one option function to override the default operations.
func NewGenWrapper[T any](wrapped TypedGenerator[T],
	opt GenWraperOptFunc[T],
	opts ...GenWraperOptFunc[T],
) (*GenWrapper[T], error) {
	gw := GenWrapper[T]{
		wrapped:  wrapped,
		stringer: dfltOperations[T]{},
		iterator: dfltOperations[T]{},
		value:    dfltOperations[T]{},
	}

	err := opt(&gw)
	if err != nil {
		return nil, err
	}

	for _, opt := range opts {
		err := opt(&gw)
		if err != nil {
			return nil, err
		}
	}

	return &gw, nil
}

// String calls the stringer member's String func and returns whatever that
// returns
func (gw GenWrapper[T]) String() string {
	return gw.stringer.String(gw.wrapped)
}

// Value calls the value member func and returns whatever that returns
func (gw GenWrapper[T]) Value() T {
	return gw.value.Value(gw.wrapped)
}

// Next calls the next member func and returns whatever that returns
func (gw GenWrapper[T]) Next() {
	gw.iterator.Next(gw.wrapped)
}
