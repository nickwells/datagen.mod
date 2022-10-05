package datagen

import (
	"fmt"
	"math/rand"
)

// WeightedString records a string and an associated weight. When there are
// multiple such entries the WStringGen will generate strings in the
// appropriate proportions.
type WeightedString struct {
	Str    string
	Weight int
}

// sgWeightedString holds the internal representation of a weighted
// string. It includes the cumulative weight of all the prior weighted
// strings.
type sgWeightedString struct {
	WeightedString
	cumWeight int
}

// WStringGen is used to generate strings selected from a population of
// weighted values.
type WStringGen struct {
	seqOrRand seqOrRandType
	r         *rand.Rand
	idx       int

	strings   []sgWeightedString
	totWeight int
}

// addWeightedString adds the weighted string to the WStringGen, calculating
// the new total weight as it does so. It will panic if the weight is <= 0.
func (sg *WStringGen) addWeightedString(ws WeightedString) {
	if ws.Weight <= 0 {
		panic(fmt.Sprintf(
			"The weight (%d) for string %q is <= 0",
			ws.Weight, ws.Str))
	}
	sg.totWeight += ws.Weight
	sg.strings = append(sg.strings, sgWeightedString{
		WeightedString: ws,
		cumWeight:      sg.totWeight,
	})
}

// NewWStringGen creates a new WStringGen object and returns it. It will panic
// if any of the weights is <= 0.
//
// The seqOrRand value can be set to Random to cause the values to be chosen
// randomly from the supplied values. Or it can be set to Sequential and the
// values are returned in the order supplied. In either case the values are
// returned in proportions reflecting the weights.
func NewWStringGen(seqOrRand seqOrRandType,
	ws WeightedString, strs ...WeightedString,
) *WStringGen {
	sg := &WStringGen{
		seqOrRand: seqOrRand,
		strings:   make([]sgWeightedString, 0, len(strs)+1),
	}

	sg.addWeightedString(ws)
	for _, ws := range strs {
		sg.addWeightedString(ws)
	}
	if sg.totWeight == 0 {
		return sg
	}

	if seqOrRand == Random {
		// TODO: The generated Rand value is not thread safe. It should be
		//       locked if it is expected to be used concurrently
		sg.r = NewRand()
		sg.idx = sg.r.Intn(sg.totWeight)
	}

	return sg
}

// Next moves the string to it's next value
func (sg *WStringGen) Next() {
	if sg.seqOrRand == Random {
		sg.idx = sg.r.Intn(sg.totWeight)
	} else {
		sg.idx++
		if sg.idx >= sg.totWeight {
			sg.idx = 0
		}
	}
}

// Generate returns the next string
func (sg WStringGen) Generate() string {
	for _, ws := range sg.strings {
		if sg.idx < ws.cumWeight {
			return ws.Str
		}
	}
	return ""
}

// Value returns the next string
func (sg WStringGen) Value() string {
	return sg.Generate()
}
