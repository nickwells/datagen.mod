package datagen

import (
	"math/rand/v2"
)

// NewRand returns a new rand.Rand with its own unique source, suitable for
// generating random variables. It generates the seed for the new value from
// the current time (in nano seconds) plus a monotonically increasing value.
//
//nolint:gosec
func NewRand() *rand.Rand {
	return rand.New(rand.NewPCG(rand.Uint64(), rand.Uint64()))
}
