package datagen

import (
	"errors"
	"fmt"
)

// CcySymbolPlacement encodes where the currency symbol should appear when
// displaying a currency amount.
//
// CCYSymBefore means that the symbol should come before the number: $1.23
//
// CcySymAfter means that the symbol comes after the number: 1.23$
//
// CcySymAtDecimal means that the symbol appears where the decimal place
// would otherwise appear: 1$23
type CcySymbolPlacement int

const (
	CcySymBefore CcySymbolPlacement = iota
	CcySymAfter
	CcySymAtDecimal
)

// Currency records details about a specific currency. These include the
// name, currency symbol, the ISO 4217 code and how many decimals it uses.
type Currency struct {
	name     string
	symbol   string
	symPlace CcySymbolPlacement
	code     string
	decimals int
}

// NumFmtWithCCY converts a number format into one that will format a number
// as a currency amount (with the currency symbol in the correct place)
func (ccy Currency) NumFmtWithCCY(nf NumFmt) NumFmt {
	switch ccy.symPlace {
	case CcySymBefore:
		nf.prefix = ccy.symbol
	case CcySymAtDecimal:
		nf.decimalSep = ccy.symbol
	case CcySymAfter:
		nf.suffix = ccy.symbol
	}

	return nf
}

// Money represents a monetary amount. It gives the amount and the
// currency. Note that the amount is held as an integer. It represents the
// number of fractional currency units and gives correct results when adding
// amounts. This is in contrast to holding money amounts as a float where
// occaisional rounding errors will give incorrect results. Note that for
// currencies without fractional units in common use (Japanese Yen, Indian
// Rupee, etc.) the amount is in the major unit (Yen, Rupee etc.)
//
// For instance a Money value of 123, with Currency of US dollars would
// represent $1.23 not $123.00.
type Money struct {
	Amt int64
	Ccy Currency
}

// MoneyStringMaker implements the StringMaker interface
type MoneyStringMaker struct {
	sm func(Money) string
}

// NewMoneyStringMaker returns a new MoneyStringMaker with the StringMaker
// function initialised correctly.
func NewMoneyStringMaker(f func(Money) string) *MoneyStringMaker {
	if f == nil {
		panic(errors.New("invalid nil value for the StringMaker func"))
	}

	return &MoneyStringMaker{sm: f}
}

// MakeString returns a string representing the money value
func (m MoneyStringMaker) MakeString(v Money) string {
	return m.sm(v)
}

// MoneyValSetter implements the ValSetter interface
type MoneyValSetter struct {
	vs ValSetter[int64]
}

// NewMoneyValSetter returns a new MoneyValSetter with the ValSetter function
// initialised correctly.
func NewMoneyValSetter(vs ValSetter[int64]) *MoneyValSetter {
	return &MoneyValSetter{vs: vs}
}

// SetVal sets the money amount to the next value
func (m MoneyValSetter) SetVal(v *Money) {
	m.vs.SetVal(&v.Amt)
}

// stripDecimals will take the value, split out the remainder on division
// by the factor and format that with the right number of decimals and the
// appropriate digital separator.
//
// For instance, stripDecimals(123, 100, 2, ".") will return (".23", 1)
func stripDecimals(v, factor int64, decimals int, sep string) (string, int64) {
	if decimals == 0 {
		return "", v
	}
	return fmt.Sprintf("%s%0*d", sep, decimals, v%factor), v / factor
}

// MoneyMkStrFunc returns a to-string function which can be used to format
// a money amount. A money amount is held as an int64 in the fractional
// currency unit (cents, pence etc).
func (ccy Currency) MoneyMkStrFunc(nf *NumFmt) func(int64) string {
	decFactor := makeFactor[int64](ccy.decimals)

	innerNF := *nf
	innerNF.prefix = ""
	innerNF.suffix = ""
	innerNF.useZeroVal = false
	uFunc := UnsignedMkStrFunc[uint64](innerNF)

	return func(v int64) string {
		if v == 0 && nf.useZeroVal {
			return nf.zeroVal
		}

		s := nf.prefix

		isNegative := v < 0
		if isNegative {
			v *= -1
		}

		var decimalPart string
		decimalPart, v = stripDecimals(v, decFactor,
			ccy.decimals, nf.decimalSep)
		nonDecimalPart := uFunc(uint64(v))

		if isNegative && nf.negFmt == NegFmtMinus {
			s += "-"
		}

		s += nonDecimalPart
		s += decimalPart + nf.suffix

		if isNegative && nf.negFmt == NegFmtAccounts {
			return "(" + s + ")"
		}

		return s
	}
}
