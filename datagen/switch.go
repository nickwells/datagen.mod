package datagen

// SwitchGen[T any] records the set of cases and the default value. The Value
// used is that for the first case that passes. If none pass then the default
// value is used.
type SwitchGen[T any] struct {
	cases   []*Case[T]
	dfltVal TypedGenerator[T]
}

// NewSwitchGen returns a new switcher of type T. Note that the first
// parameter gives the default value. This will be used if none of the cases
// pass. The supplied cases are evaluated in the order they are given and the
// first one to pass is used to supply the generated value.
func NewSwitchGen[T any](dfltVal TypedGenerator[T], cases ...*Case[T]) *SwitchGen[T] {
	return &SwitchGen[T]{
		cases:   cases,
		dfltVal: dfltVal,
	}
}

// Next moves the values on to their next value
func (sg *SwitchGen[T]) Next() {
	sg.dfltVal.Next()
	for _, c := range sg.cases {
		c.v.Next()
	}
}

// Generate generates and returns the next value as a string
func (sg SwitchGen[T]) Generate() string {
	for _, c := range sg.cases {
		if c.vCk.Passes() {
			return c.v.Generate()
		}
	}

	return sg.dfltVal.Generate()
}

// Value returns the next value
func (sg SwitchGen[T]) Value() T {
	for _, c := range sg.cases {
		if c.vCk.Passes() {
			return c.v.Value()
		}
	}

	return sg.dfltVal.Value()
}
