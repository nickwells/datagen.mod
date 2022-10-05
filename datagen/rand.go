package datagen

import (
	"math/rand"
	"time"
)

// randCount is used to ensure uniqueness of the sources
var randCount int64

// NewRand returns a new rand.Rand with its own unique source, suitable for
// generating random variables. It generates the seed for the new value from
// the current time (in nano seconds) plus a monotonically increasing value.
func NewRand() *rand.Rand {
	now := time.Now()
	randCount++
	return rand.New(rand.NewSource(int64(now.Nanosecond()) + randCount))
}
