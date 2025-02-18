package datagen

// ComputedValSetter applies the aggregator to the values. The aggregator is
// expected to apply the result of the aggregation to the passed value.
type ComputedValSetter[T any] struct {
	values     []TypedGenerator[T]
	aggregator func(*T, ...TypedGenerator[T])
}

// SetVal applies the aggregator to the passed value and the list of values.
func (vs ComputedValSetter[T]) SetVal(v *T) {
	vs.aggregator(v, vs.values...)
}

// NewComputedValSetter builds and returns a ComputedValSetter of the given
// type.
func NewComputedValSetter[T any](aggregator func(*T, ...TypedGenerator[T]),
	val0 TypedGenerator[T], vals ...TypedGenerator[T],
) *ComputedValSetter[T] {
	vs := &ComputedValSetter[T]{
		values:     make([]TypedGenerator[T], 0, len(vals)+1),
		aggregator: aggregator,
	}
	vs.values = append(vs.values, val0)
	vs.values = append(vs.values, vals...)

	return vs
}
