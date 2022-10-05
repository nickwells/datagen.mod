package datagen

// NewConstGen returns a new Gen of type T which always generates the given
// string and doesn't change the value.
func NewConstGen[T any](s string) *Gen[T] {
	return NewGen(GenSetStringMaker[T](ConstMakeString[T]{Str: s}))
}
