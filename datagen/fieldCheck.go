package datagen

import "github.com/nickwells/check.mod/v2/check"

// Passer
type Passer interface {
	Passes() bool
}

// ValCk represents a check to be performed on a value
type ValCk struct {
	Passer
}

// NewValCk[T any] constructs a simple value check that can check whether or
// not the supplied value passes the supplied check
func NewValCk[T any](ck check.ValCk[T], val TypedVal[T]) *ValCk {
	return &ValCk{
		passer[T]{
			ck:  ck,
			val: val,
		},
	}
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

// ===================================================================

// passerAnd represents a typed value and a test to be performed on it. It
// implements the Passer interface.
type passerAnd struct {
	vCks []*ValCk
}

// Passes returns true if all the checks pass and false if any of them fails
func (p passerAnd) Passes() bool {
	for _, vCk := range p.vCks {
		if !vCk.Passes() {
			return false
		}
	}
	return true
}

// NewValCkAnd constructs a value check that will pass if all the supplied
// checks pass and fail otherwise.
func NewValCkAnd(c *ValCk, checks ...*ValCk) *ValCk {
	p := passerAnd{
		vCks: make([]*ValCk, 0, len(checks)+1),
	}

	p.vCks = append(p.vCks, c)
	p.vCks = append(p.vCks, checks...)

	return &ValCk{p}
}

// ===================================================================

// passerOr represents a typed value and a test to be performed on it. It
// implements the Passer interface.
type passerOr struct {
	vCks []*ValCk
}

// Passes returns true if any of the checks passes and false if they all fail
func (p passerOr) Passes() bool {
	for _, vCk := range p.vCks {
		if vCk.Passes() {
			return true
		}
	}
	return false
}

// NewValCkOr constructs a value check that will pass if any of the supplied
// checks pass and fail if none of them do.
func NewValCkOr(c *ValCk, checks ...*ValCk) *ValCk {
	p := passerOr{
		vCks: make([]*ValCk, 0, len(checks)+1),
	}

	p.vCks = append(p.vCks, c)
	p.vCks = append(p.vCks, checks...)

	return &ValCk{p}
}
