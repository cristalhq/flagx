package flagx

import (
	"flag"
	"io"
	"time"
)

// FlagSet represents a set of flags.
// Flag names must be unique within a FlagSet.
// An attempt to define a flag whose name is already in use will cause a panic.
type FlagSet struct {
	fs *flag.FlagSet
}

// NewFlagSet returns new FlagSet.
func NewFlagSet(name string, output io.Writer) *FlagSet {
	fs := flag.NewFlagSet(name, flag.ContinueOnError)
	fs.SetOutput(output)
	return &FlagSet{fs: fs}
}

// AsStdlib returns *flag.FlagSet with all flags.
// Note: alias are duplicated.
func (f *FlagSet) AsStdlib() *flag.FlagSet {
	return f.fs
}

// NFlag returns the number of flags that have been set.
func (f *FlagSet) NFlag() int { return f.fs.NFlag() }

// NArg is the number of arguments remaining after flags have been processed.
func (f *FlagSet) NArg() int { return f.fs.NArg() }

// Arg returns the i'th argument. Arg(0) is the first remaining argument
// after flags have been processed. Arg returns an empty string if the
// requested element does not exist.
func (f *FlagSet) Arg(i int) string { return f.fs.Arg(i) }

// Args returns the non-flag arguments.
func (f *FlagSet) Args() []string { return f.fs.Args() }

// IsParsed reports whether f.Parse has been called.
func (f *FlagSet) IsParsed() bool { return f.fs.Parsed() }

// Parse parses flag definitions from the argument list, which should not
// include the command name. Must be called after all flags in the FlagSet
// are defined and before flags are accessed by the program.
// The return value will be ErrHelp if -help or -h were set but not defined.
func (f *FlagSet) Parse(arguments []string) error { return f.fs.Parse(arguments) }

func (f *FlagSet) PrintDefaults()                { f.fs.PrintDefaults() }
func (f *FlagSet) VisitAll(fn func(*flag.Flag))  { f.fs.VisitAll(fn) }
func (f *FlagSet) Visit(fn func(*flag.Flag))     { f.fs.Visit(fn) }
func (f *FlagSet) Lookup(name string) *flag.Flag { return f.fs.Lookup(name) }
func (f *FlagSet) Set(name, value string) error  { return f.fs.Set(name, value) }

//  defines a flag with the specified name and usage string. The type and
// value of the flag are represented by the first argument, of type Value, which
// typically holds a user-defined implementation of Value. For instance, the
// caller could create a flag that turns a comma-separated string into a slice
// of strings by giving the slice the methods of Value; in particular, Set would
// decompose the comma-separated string into the slice.
// Empty string for alias means no alias will be created.
func (f *FlagSet) Var(value flag.Value, name, alias, usage string) {
	f.fs.Var(value, name, usage)
	if alias != "" {
		f.fs.Var(value, alias, usage)
	}
}

// Func defines a flag with the specified name and usage string.
// Each time the flag is seen, fn is called with the value of the flag.
// If fn returns a non-nil error, it will be treated as a flag value parsing error.
// Empty string for alias means no alias will be created.
func (f *FlagSet) Func(name, alias, usage string, fn func(string) error) {
	f.fs.Func(name, usage, fn)
	if alias != "" {
		f.fs.Func(name, alias, fn)
	}
}

// Bool defines a bool flag with specified name, alias, default value, and usage string.
// The argument p points to a bool variable in which to store the value of the flag.
// Empty string for alias means no alias will be created.
func (f *FlagSet) Bool(p *bool, name, alias string, value bool, usage string) {
	f.fs.BoolVar(p, name, value, usage)
	if alias != "" {
		f.fs.BoolVar(p, alias, value, usage)
	}
}

// Int defines an int flag with specified name, alias, default value, and usage string.
// The argument p points to an int variable in which to store the value of the flag.
// Empty string for alias means no alias will be created.
func (f *FlagSet) Int(p *int, name, alias string, value int, usage string) {
	f.fs.IntVar(p, name, value, usage)
	if alias != "" {
		f.fs.IntVar(p, alias, value, usage)
	}
}

// Int64 defines an int64 flag with specified name, alias, default value, and usage string.
// The argument p points to an int64 variable in which to store the value of the flag.
// Empty string for alias means no alias will be created.
func (f *FlagSet) Int64(p *int64, name, alias string, value int64, usage string) {
	f.fs.Int64Var(p, name, value, usage)
	if alias != "" {
		f.fs.Int64Var(p, alias, value, usage)
	}
}

// Uint defines a uint flag with specified name, alias, default value, and usage string.
// The argument p points to a uint variable in which to store the value of the flag.
// Empty string for alias means no alias will be created.
func (f *FlagSet) Uint(p *uint, name, alias string, value uint, usage string) {
	f.fs.UintVar(p, name, value, usage)
	if alias != "" {
		f.fs.UintVar(p, alias, value, usage)
	}
}

// Uint64 defines a uint64 flag with specified name, alias, default value, and usage string.
// The argument p points to a uint64 variable in which to store the value of the flag.
// Empty string for alias means no alias will be created.
func (f *FlagSet) Uint64(p *uint64, name, alias string, value uint64, usage string) {
	f.fs.Uint64Var(p, name, value, usage)
	if alias != "" {
		f.fs.Uint64Var(p, alias, value, usage)
	}
}

// String defines a string flag with specified name, alias, default value, and usage string.
// The argument p points to a string variable in which to store the value of the flag.
// Empty string for alias means no alias will be created.
func (f *FlagSet) String(p *string, name, alias, value, usage string) {
	f.fs.StringVar(p, name, value, usage)
	if alias != "" {
		f.fs.StringVar(p, alias, value, usage)
	}
}

// Float64 defines a float64 flag with specified name, alias, default value, and usage string.
// The argument p points to a float64 variable in which to store the value of the flag.
// Empty string for alias means no alias will be created.
func (f *FlagSet) Float64(p *float64, name, alias string, value float64, usage string) {
	f.fs.Float64Var(p, name, value, usage)
	if alias != "" {
		f.fs.Float64Var(p, alias, value, usage)
	}
}

// Duration defines a time.Duration flag with specified name, alias, default value, and usage string.
// The argument p points to a time.Duration variable in which to store the value of the flag.
// The flag accepts a value acceptable to time.ParseDuration.
// Empty string for alias means no alias will be created.
func (f *FlagSet) Duration(p *time.Duration, name, alias string, value time.Duration, usage string) {
	f.fs.DurationVar(p, name, value, usage)
	if alias != "" {
		f.fs.DurationVar(p, alias, value, usage)
	}
}

// BoolSlice defines a slice of bool flag with specified name, alias, default value, and usage string.
// The argument p points to a bool variable in which to store the value of the flag.
// Empty string for alias means no alias will be created.
func (f *FlagSet) BoolSlice(p *[]bool, name, alias string, value []bool, usage string) {
	var sb SliceOfBool
	*p = []bool(sb)
	f.fs.Var(&sb, name, usage)
	if alias != "" {
		f.fs.Var(&sb, alias, usage)
	}
}

// IntSlice defines a slice of int flag with specified name, alias, default value, and usage string.
// The argument p points to an int variable in which to store the value of the flag.
// Empty string for alias means no alias will be created.
func (f *FlagSet) IntSlice(p *[]int, name, alias string, value []int, usage string) {
	var si SliceOfInt = value
	*p = []int(si)
	f.fs.Var(&si, name, usage)
	if alias != "" {
		f.fs.Var(&si, alias, usage)
	}
}

// Int64Slice defines a slice of int64 flag with specified name, alias, default value, and usage string.
// The argument p points to an int64 variable in which to store the value of the flag.
// Empty string for alias means no alias will be created.
func (f *FlagSet) Int64Slice(p *[]int64, name, alias string, value []int64, usage string) {
	var si SliceOfInt64 = value
	*p = []int64(si)
	f.fs.Var(&si, name, usage)
	if alias != "" {
		f.fs.Var(&si, alias, usage)
	}
}

// UintSlice defines a slice of uint flag with specified name, alias, default value, and usage string.
// The argument p points to a uint variable in which to store the value of the flag.
// Empty string for alias means no alias will be created.
func (f *FlagSet) UintSlice(p *[]uint, name, alias string, value []uint, usage string) {
	var su SliceOfUint = value
	*p = []uint(su)
	f.fs.Var(&su, name, usage)
	if alias != "" {
		f.fs.Var(&su, alias, usage)
	}
}

// Uint64Slice defines a slice of uint64 flag with specified name, alias, default value, and usage string.
// The argument p points to a uint64 variable in which to store the value of the flag.
// Empty string for alias means no alias will be created.
func (f *FlagSet) Uint64Slice(p *[]uint64, name, alias string, value []uint64, usage string) {
	var su SliceOfUint64 = value
	*p = []uint64(su)
	f.fs.Var(&su, name, usage)
	if alias != "" {
		f.fs.Var(&su, alias, usage)
	}
}

// StringSlice defines a slice of string flag with specified name, alias, default value, and usage string.
// The argument p points to a string variable in which to store the value of the flag.
// Empty string for alias means no alias will be created.
func (f *FlagSet) StringSlice(p *[]string, name, alias string, value []string, usage string) {
	var ss SliceOfString = value
	*p = []string(ss)
	f.fs.Var(&ss, name, usage)
	if alias != "" {
		f.fs.Var(&ss, alias, usage)
	}
}

// Float64Slice defines a slice of float64 flag with specified name, alias, default value, and usage string.
// The argument p points to a float64 variable in which to store the value of the flag.
// Empty string for alias means no alias will be created.
func (f *FlagSet) Float64Slice(p *[]float64, name, alias string, value []float64, usage string) {
	var sf SliceOfFloat64 = value
	*p = []float64(sf)
	f.fs.Var(&sf, name, usage)
	if alias != "" {
		f.fs.Var(&sf, alias, usage)
	}
}

// DurationSlice defines a slice of time.Duration flag with specified name, alias, default value, and usage string.
// The argument p points to a time.Duration variable in which to store the value of the flag.
// The flag accepts a value acceptable to time.ParseDuration.
// Empty string for alias means no alias will be created.
func (f *FlagSet) DurationSlice(p *[]time.Duration, name, alias string, value []time.Duration, usage string) {
	var sd SliceOfDuration = value
	*p = []time.Duration(sd)
	f.fs.Var(&sd, name, usage)
	if alias != "" {
		f.fs.Var(&sd, alias, usage)
	}
}

// IntSet defines a set of int flag with specified name, alias, default value, and usage string.
// The argument p points to an int variable in which to store the value of the flag.
// Empty string for alias means no alias will be created.
func (f *FlagSet) IntSet(p *map[int]struct{}, name, alias string, value map[int]struct{}, usage string) {
	var si SetOfInt = value
	*p = map[int]struct{}(si)
	f.fs.Var(&si, name, usage)
	if alias != "" {
		f.fs.Var(&si, alias, usage)
	}
}

// Int64Set defines a set of int64 flag with specified name, alias, default value, and usage string.
// The argument p points to an int64 variable in which to store the value of the flag.
// Empty string for alias means no alias will be created.
func (f *FlagSet) Int64Set(p *map[int64]struct{}, name, alias string, value map[int64]struct{}, usage string) {
	var si SetOfInt64 = value
	*p = map[int64]struct{}(si)
	f.fs.Var(&si, name, usage)
	if alias != "" {
		f.fs.Var(&si, alias, usage)
	}
}

// UintSet defines a set of uint flag with specified name, alias, default value, and usage string.
// The argument p points to a uint variable in which to store the value of the flag.
// Empty string for alias means no alias will be created.
func (f *FlagSet) UintSet(p *map[uint]struct{}, name, alias string, value map[uint]struct{}, usage string) {
	var su SetOfUint = value
	*p = map[uint]struct{}(su)
	f.fs.Var(&su, name, usage)
	if alias != "" {
		f.fs.Var(&su, alias, usage)
	}
}

// Uint64Set defines a set of uint64 flag with specified name, alias, default value, and usage string.
// The argument p points to a uint64 variable in which to store the value of the flag.
// Empty string for alias means no alias will be created.
func (f *FlagSet) Uint64Set(p *map[uint64]struct{}, name, alias string, value map[uint64]struct{}, usage string) {
	var su SetOfUint64 = value
	*p = map[uint64]struct{}(su)
	f.fs.Var(&su, name, usage)
	if alias != "" {
		f.fs.Var(&su, alias, usage)
	}
}

// StringSet defines a set of string flag with specified name, alias, default value, and usage string.
// The argument p points to a string variable in which to store the value of the flag.
// Empty string for alias means no alias will be created.
func (f *FlagSet) StringSet(p *map[string]struct{}, name, alias string, value map[string]struct{}, usage string) {
	var ss SetOfString = value
	*p = map[string]struct{}(ss)
	f.fs.Var(&ss, name, usage)
	if alias != "" {
		f.fs.Var(&ss, alias, usage)
	}
}

// Float64Set defines a set of float64 flag with specified name, alias, default value, and usage string.
// The argument p points to a float64 variable in which to store the value of the flag.
// Empty string for alias means no alias will be created.
func (f *FlagSet) Float64Set(p *map[float64]struct{}, name, alias string, value map[float64]struct{}, usage string) {
	var sf SetOfFloat64 = value
	*p = map[float64]struct{}(sf)
	f.fs.Var(&sf, name, usage)
	if alias != "" {
		f.fs.Var(&sf, alias, usage)
	}
}

// DurationSet defines a set of time.Duration flag with specified name, alias, default value, and usage string.
// The argument p points to a time.Duration variable in which to store the value of the flag.
// The flag accepts a value acceptable to time.ParseDuration.
// Empty string for alias means no alias will be created.
func (f *FlagSet) DurationSet(p *map[time.Duration]struct{}, name, alias string, value map[time.Duration]struct{}, usage string) {
	var sd SetOfDuration = value
	*p = map[time.Duration]struct{}(sd)
	f.fs.Var(&sd, name, usage)
	if alias != "" {
		f.fs.Var(&sd, alias, usage)
	}
}
