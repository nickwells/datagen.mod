package datagen

import (
	"errors"
	"fmt"
)

// dfltGenImpl implements default SetVal and MakeString methods
type dfltGenImpl[T any] struct{}

// MakeString returns the Go string representation of the value
func (dfltGenImpl[T]) MakeString(v T) string {
	return fmt.Sprint(v)
}

// SetVal leaves the value unchanged
func (dfltGenImpl[T]) SetVal(_ *T) {
}

// Gen[T any] records the information needed to generate a field. It
// implements the TypedGenerator interface.
type Gen[T any] struct {
	sm        StringMaker[T]
	valSetter ValSetter[T]
	value     T
}

// Generate returns the string form of the value
func (s Gen[T]) Generate() string {
	return s.sm.MakeString(s.value)
}

// Next calculates the next value
func (s *Gen[T]) Next() {
	s.valSetter.SetVal(&s.value)
}

// Value returns the internal form of the value
func (s Gen[T]) Value() T {
	return s.value
}

type GenOptFunc[T any] func(s *Gen[T]) error

// GenSetStringMaker returns an option func that will set the function used
// to generate the string value.
func GenSetStringMaker[T any](sm StringMaker[T]) GenOptFunc[T] {
	return func(s *Gen[T]) error {
		if sm == nil {
			return errors.New("a nil string maker has been supplied")
		}

		s.sm = sm

		return nil
	}
}

// GenSetValue returns an option func that will set the value on a
// Gen to the supplied value
func GenSetValue[T any](v T) GenOptFunc[T] {
	return func(s *Gen[T]) error {
		s.value = v
		return nil
	}
}

// GenSetValSetter returns an option func that will set the valSetter on a
// Gen to the supplied value
func GenSetValSetter[T any](vs ValSetter[T]) GenOptFunc[T] {
	return func(s *Gen[T]) error {
		if vs == nil {
			return errors.New("a nil value setter has been supplied")
		}

		s.valSetter = vs

		return nil
	}
}

// NewGen constructs and returns a new generator of the supplied generic type
func NewGen[T any](opts ...GenOptFunc[T]) *Gen[T] {
	s := &Gen[T]{
		sm:        dfltGenImpl[T]{},
		valSetter: dfltGenImpl[T]{},
	}

	for _, o := range opts {
		if err := o(s); err != nil {
			panic(err)
		}
	}

	return s
}
