package datagen

import (
	"math/rand/v2"

	"golang.org/x/exp/constraints"
)

// ValSetter is the interface wrapping the SetVal method
type ValSetter[T any] interface {
	SetVal(*T)
}

// ===================================================================

// IncrementingValSetter implements a ValSetter that will increment the
// passed value by the incr amount
type IncrementingValSetter[T constraints.Integer | constraints.Float] struct {
	incr T
}

// NewIncrementingValSetter creates and returns an IncrementingValSetter
func NewIncrementingValSetter[T constraints.Integer | constraints.Float](
	i T,
) *IncrementingValSetter[T] {
	return &IncrementingValSetter[T]{incr: i}
}

// SetVal increments the given value by the incr amount
func (vs IncrementingValSetter[T]) SetVal(v *T) {
	*v += vs.incr
}

// ===================================================================

// NormValSetter implements a ValSetter that will set the passed value to a
// normally distributed value constrained by the min, max, mean and sd values
type NormValSetter[T constraints.Integer | constraints.Float] struct {
	r        *rand.Rand
	min, max T
	mean, sd float64
}

// NewNormValSetter creates and returns a NormValSetter
func NewNormValSetter[T constraints.Integer | constraints.Float](
	min, max T,
	mean, sd float64,
) *NormValSetter[T] {
	return &NormValSetter[T]{
		r:    NewRand(),
		min:  min,
		max:  max,
		mean: mean,
		sd:   sd,
	}
}

// SetVal increments the given value by the incr amount
func (vs NormValSetter[T]) SetVal(v *T) {
	trial := T(vs.r.NormFloat64()*vs.sd + vs.mean)

	if trial > vs.max {
		*v = vs.max
		return
	}

	if trial < vs.min {
		*v = vs.min
		return
	}

	*v = trial
}
