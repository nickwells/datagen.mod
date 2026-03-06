package datagen

import "github.com/nickwells/check.mod/v2/check"

// Passer is the interface describing the Passes method
type Passer interface {
	Passes() bool
}

// passer represents a typed value and a test to be performed on it. It
// implements the Passer interface.
type passer[T any] struct {
	ck  check.ValCk[T]
	val TypedVal[T]
}

// Passes returns true if the check passes (returns a nil error)
func (p passer[T]) Passes() bool {
	return p.ck(p.val.Value()) == nil
}

// MakePasser constructs a simple Passer that checks whether the supplied
// value passes the supplied check.
func MakePasser[T any](ck check.ValCk[T], val TypedVal[T]) Passer {
	return &passer[T]{
		ck:  ck,
		val: val,
	}
}

// ===================================================================

// boolPasser is a Passer that simply returns the current value of the given
// bool.
type boolPasser struct {
	val TypedVal[bool]
}

// Passes returns the current value of the val
func (bp boolPasser) Passes() bool {
	return bp.val.Value()
}

// MakeBoolPasser constructs a bool Passer that returns the current value of
// the bool TypedValue.
func MakeBoolPasser(val TypedVal[bool]) Passer {
	return boolPasser{
		val: val,
	}
}

// ===================================================================

// multiPasser holds a list of Passers and the function to be performed on
// it. It implements the Passer interface.
type multiPasser struct {
	passers  []Passer
	passFunc func([]Passer) bool
}

// makeMultiPasser constructs a multiPasser. At least one Passer must be given
func makeMultiPasser(
	f func([]Passer) bool, c Passer, checks ...Passer,
) Passer {
	p := multiPasser{
		passers:  make([]Passer, 0, len(checks)+1),
		passFunc: f,
	}

	p.passers = append(p.passers, c)
	p.passers = append(p.passers, checks...)

	return &p
}

// Passes returns whatever the passFunc returns
func (p multiPasser) Passes() bool {
	return p.passFunc(p.passers)
}

// orPassFunc returns true if any of the checks passes and false if they all
// fail.
func orPassFunc(ps []Passer) bool {
	for _, p := range ps {
		if p.Passes() {
			return true
		}
	}

	return false
}

// andPassFunc returns true if all the checks pass and false if any of them
// fails
func andPassFunc(ps []Passer) bool {
	for _, p := range ps {
		if !p.Passes() {
			return false
		}
	}

	return true
}

// MakeAndPasser constructs a Passer that will pass if all the supplied
// checks pass and fail otherwise.
func MakeAndPasser(c Passer, checks ...Passer) Passer {
	return makeMultiPasser(andPassFunc, c, checks...)
}

// MakeOrPasser constructs a Passer that will pass if any of the supplied
// checks pass and fail if none of them do.
func MakeOrPasser(c Passer, checks ...Passer) Passer {
	return makeMultiPasser(orPassFunc, c, checks...)
}

// NotPasser wraps a Passer and its Passes method will return the logical NOT
// of the wrapped Passer
type NotPasser struct {
	p Passer
}

// Passes returns the logical NOT of the wrapped Passer's Passes method
func (np NotPasser) Passes() bool {
	return !np.p.Passes()
}
