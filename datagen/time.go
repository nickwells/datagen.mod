package datagen

import (
	"math"
	"math/rand/v2"
	"time"
)

const dfltTimeGenLayout = "2006/01/02 15:04:05.000"

type Time2Str struct {
	format string
}

// ToStr returns a string representing the time
func (t2s Time2Str) ToStr(t time.Time) string {
	return t.Format(t2s.format)
}

// NewTime2Str returns a new Time2Str with the format value set
func NewTime2Str(f string) *Time2Str {
	return &Time2Str{format: f}
}

// TimeValSetConstInterval provides a SetVal method that will set the time
// value to a constant interval from its current value.
type TimeValSetConstInterval struct {
	// Ival gives the constant duration used as the interval
	Ival time.Duration
}

// SetVal sets the time to a constant duration from its current value
func (tvs TimeValSetConstInterval) SetVal(t *time.Time) {
	*t = t.Add(tvs.Ival)
}

// TimeValSetGaussianInterval provides a SetVal method that will set the time
// value to a random interval from its current value. The interval used is
// some random value centered on the mean with a standard deviation of sd and
// then scaled by the units.
type TimeValSetGaussianInterval struct {
	// mean gives the mean value of the intervals to be used
	mean float64
	// sd gives the standard deviation of the intervals
	sd float64

	// r is the random number generator
	r *rand.Rand

	// units is the Duration units in which the mean and sd are expressed
	units time.Duration

	// forceGT0 is a flag which if set to true will force any duration to be
	// greater than zero
	forceGT0 bool
}

// NewTimeValSetGaussianInterval constructs a new time value setter with a
// random interval which follows a Gaussian (normal) distribution.
func NewTimeValSetGaussianInterval(mean, sd float64,
	units time.Duration,
	forceGT0 bool,
) *TimeValSetGaussianInterval {
	tvs := &TimeValSetGaussianInterval{
		r:        NewRand(),
		mean:     mean,
		sd:       sd,
		units:    units,
		forceGT0: forceGT0,
	}

	return tvs
}

// SetVal sets the time to a random duration from its current value. This
// interval may be negative unless the forceGT0 flag is set. The interval
// will always be in whole multiples of the units.
func (tvs TimeValSetGaussianInterval) SetVal(t *time.Time) {
	f := rand.NormFloat64() //nolint:gosec
	if f <= 0 && tvs.forceGT0 {
		if f == 0 {
			f = 1.0
		} else {
			f *= -1.0
		}
	}

	ival64 := f*tvs.sd + tvs.mean
	ival := time.Duration(math.Floor(ival64)) * tvs.units

	*t = t.Add(ival)
}

// ==============================

// TimeGenIntervalF is the type of the function that returns the next
// interval between times
type TimeGenIntervalF func(time.Time) time.Duration

// TimeGenConstIntervalF returns an interval func which always returns the
// supplied duration.
func TimeGenConstIntervalF(d time.Duration) TimeGenIntervalF {
	return func(_ time.Time) time.Duration {
		return d
	}
}

// TimeGen records the information needed to generate a time field
type TimeGen struct {
	layout    string
	value     time.Time
	intervalF TimeGenIntervalF
}

type TimeGenOptFunc func(tg *TimeGen) error

// TimeGenSetLayout returns a TimeGen Opt function which sets the layout (the
// format for displaying the time)
func TimeGenSetLayout(layout string) TimeGenOptFunc {
	return func(tg *TimeGen) error {
		tg.layout = layout
		return nil
	}
}

// TimeGenSetInitialTime returns a TimeGen Opt function which sets the
// initial current time. The default value is the current time.
func TimeGenSetInitialTime(t time.Time) TimeGenOptFunc {
	return func(tg *TimeGen) error {
		tg.value = t
		return nil
	}
}

// TimeGenSetIntervalF returns a TimeGen Opt function which sets the
// function which generates the interval between subsequent times
func TimeGenSetIntervalF(f TimeGenIntervalF) TimeGenOptFunc {
	return func(tg *TimeGen) error {
		tg.intervalF = f
		return nil
	}
}

// NewTimeGen creates a new TimeGen object. It will panic if any of the
// option functions returns an error.
func NewTimeGen(opts ...TimeGenOptFunc) *TimeGen {
	tg := &TimeGen{
		layout:    dfltTimeGenLayout,
		value:     time.Now(),
		intervalF: TimeGenConstIntervalF(time.Second),
	}

	for _, o := range opts {
		if err := o(tg); err != nil {
			panic(err)
		}
	}

	return tg
}

// Generate generates a formatted time string
func (tg TimeGen) Generate() string {
	return tg.value.Format(tg.layout)
}

// Value returns the current value as a time
func (tg TimeGen) Value() time.Time {
	return tg.value
}

// Next moves the time on to its next value
func (tg *TimeGen) Next() {
	tg.value = tg.value.Add(tg.intervalF(tg.value))
}
