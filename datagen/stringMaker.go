package datagen

// StringMaker is the interface wrapping the MakeString method
type StringMaker[T any] interface {
	MakeString(T) string
}

// ConstMakeString implements a const MakeString method always returning the same string
type ConstMakeString[T any] struct {
	Str string
}

// MakeString returns the const string
func (s ConstMakeString[T]) MakeString(_ T) string {
	return s.Str
}
