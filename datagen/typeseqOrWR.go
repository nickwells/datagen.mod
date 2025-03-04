package datagen

// seqOrRandType control whether values are generated as random values or
// sequentially
type seqOrRandType int

const (
	Random seqOrRandType = iota
	Sequential
)

// IsValid is a method on the seqOrRandType type that can be used
// to check a received parameter for validity. It compares
// the value against the boundary values for the type
// and returns false if it is outside the valid range
func (v seqOrRandType) IsValid() bool {
	if v < Random {
		return false
	}

	if v > Sequential {
		return false
	}

	return true
}
