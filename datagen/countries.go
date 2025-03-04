//nolint:mnd
package datagen

// Country records details about a country. These include the name, the ISO
// 3166 code, the number format used and the currency.
type Country struct {
	name string
	code string
	nf   NumFmt
	ccy  Currency
}

// Name returns the country name
func (c Country) Name() string {
	return c.name
}

// Code returns the country's ISO 3166 code
func (c Country) Code() string {
	return c.code
}

// NF returns the country's number format details
func (c Country) NF() NumFmt {
	return c.nf
}

// Ccy returns the country's currency details
func (c Country) Ccy() Currency {
	return c.ccy
}

// Countries is a map giving the country details of the top 10 countries by
// GDP (as of August 2022). The map keys are the ISO 3166 codes for the
// countries.
var Countries = map[string]Country{
	"US": {
		name: "United States of America",
		code: "US",
		nf: NumFmt{
			decimalSep:  ".",
			digitGrpSep: ",",
			sepCount:    []int{3},
		},
		ccy: Currency{
			name:     "dollar",
			symbol:   "$",
			symPlace: CcySymBefore,
			code:     "USD",
			decimals: 2,
		},
	},
	"CN": {
		name: "People's Republic of China",
		code: "CN",
		nf: NumFmt{
			decimalSep:  ".",
			digitGrpSep: ",",
			sepCount:    []int{4},
		},
		ccy: Currency{
			name:     "yuan",
			symbol:   "¥",
			symPlace: CcySymBefore,
			code:     "CNY",
			decimals: 2,
		},
	},
	"JP": {
		name: "Japan",
		code: "JP",
		nf: NumFmt{
			decimalSep:  ".",
			digitGrpSep: ",",
			sepCount:    []int{3},
		},
		ccy: Currency{
			name:     "yen",
			symbol:   "¥",
			symPlace: CcySymBefore,
			code:     "JPY",
			decimals: 0,
		},
	},
	"DE": {
		name: "Federal Republic of Germany",
		code: "DE",
		nf: NumFmt{
			decimalSep:  ",",
			digitGrpSep: ".",
			sepCount:    []int{3},
		},
		ccy: Currency{
			name:     "euro",
			symbol:   "€",
			symPlace: CcySymAfter,
			code:     "EUR",
			decimals: 2,
		},
	},
	"IN": {
		name: "Republic of India",
		code: "IN",
		nf: NumFmt{
			decimalSep:  ".",
			digitGrpSep: ",",
			sepCount:    []int{3, 2},
		},
		ccy: Currency{
			name:     "rupee",
			symbol:   "₹",
			symPlace: CcySymBefore,
			code:     "INR",
			decimals: 0,
		},
	},
	"GB": {
		name: "United Kingdom",
		code: "GB",
		nf: NumFmt{
			decimalSep:  ".",
			digitGrpSep: ",",
			sepCount:    []int{3},
		},
		ccy: Currency{
			name:     "pound",
			symbol:   "£",
			symPlace: CcySymBefore,
			code:     "GBP",
			decimals: 2,
		},
	},
	"FR": {
		name: "French Republic",
		code: "FR",
		nf: NumFmt{
			decimalSep:  ",",
			digitGrpSep: ".",
			sepCount:    []int{3},
		},
		ccy: Currency{
			name:     "euro",
			symbol:   "€",
			symPlace: CcySymAfter,
			code:     "EUR",
			decimals: 2,
		},
	},
	"BR": {
		name: "Federative Republic of Brazil",
		code: "BR",
		nf: NumFmt{
			decimalSep:  ",",
			digitGrpSep: ".",
			sepCount:    []int{3},
		},
		ccy: Currency{
			name:     "real",
			symbol:   "R$",
			symPlace: CcySymBefore,
			code:     "BRL",
			decimals: 2,
		},
	},
	"IT": {
		name: "Italian Republic",
		code: "IT",
		nf: NumFmt{
			decimalSep:  ",",
			digitGrpSep: ".",
			sepCount:    []int{3},
		},
		ccy: Currency{
			name:     "euro",
			symbol:   "€",
			symPlace: CcySymAfter,
			code:     "EUR",
			decimals: 2,
		},
	},
	"CA": {
		name: "Canada",
		code: "CA",
		nf: NumFmt{
			decimalSep:  ".",
			digitGrpSep: ",",
			sepCount:    []int{3},
		},
		ccy: Currency{
			name:     "dollar",
			symbol:   "$",
			symPlace: CcySymBefore,
			code:     "CAD",
			decimals: 2,
		},
	},
	"RU": {
		name: "Russian Federation",
		code: "RU",
		nf: NumFmt{
			decimalSep:  ",",
			digitGrpSep: " ",
			sepCount:    []int{3},
		},
		ccy: Currency{
			name:     "ruble",
			symbol:   "₽",
			symPlace: CcySymAfter,
			code:     "RUR",
			decimals: 2,
		},
	},
}
