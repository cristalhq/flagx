package flagx

import (
	"flag"
	"io"
	"time"
)

// FlagSet ...
type FlagSet struct {
	fs *flag.FlagSet
}

// NewFlagSet returns new FlagSet.
func NewFlagSet(name string) *FlagSet {
	return &FlagSet{fs: flag.NewFlagSet(name, flag.ContinueOnError)}
}

// AsStdlibFlagSet returns *flag.FlagSet with all flags.
func (fs *FlagSet) AsStdlibFlagSet() *flag.FlagSet {
	return fs.fs
}

func (f *FlagSet) Output() io.Writer                                  { return f.fs.Output() }
func (f *FlagSet) Name() string                                       { return f.fs.Name() }
func (f *FlagSet) ErrorHandling() flag.ErrorHandling                  { return f.fs.ErrorHandling() }
func (f *FlagSet) SetOutput(output io.Writer)                         { f.fs.SetOutput(output) }
func (f *FlagSet) VisitAll(fn func(*flag.Flag))                       { f.fs.VisitAll(fn) }
func (f *FlagSet) Visit(fn func(*flag.Flag))                          { f.fs.Visit(fn) }
func (f *FlagSet) Lookup(name string) *flag.Flag                      { return f.fs.Lookup(name) }
func (f *FlagSet) Set(name, value string) error                       { return f.fs.Set(name, value) }
func (f *FlagSet) Parse(arguments []string) error                     { return f.fs.Parse(arguments) }
func (f *FlagSet) Parsed() bool                                       { return f.fs.Parsed() }
func (f *FlagSet) Init(name string, errorHandling flag.ErrorHandling) { f.fs.Init(name, errorHandling) }
func (f *FlagSet) PrintDefaults()                                     { f.fs.PrintDefaults() }
func (f *FlagSet) NFlag() int                                         { return f.fs.NFlag() }
func (f *FlagSet) Arg(i int) string                                   { return f.fs.Arg(i) }
func (f *FlagSet) NArg() int                                          { return f.fs.NArg() }
func (f *FlagSet) Args() []string                                     { return f.Args() }


// BoolVar defines a bool flag with specified name, alias, default value, and usage string.
// The argument p points to a bool variable in which to store the value of the flag.
// Empty string for alias means no alias will be created.
func (f *FlagSet) BoolVar(p *bool, name, alias string, value bool, usage string) {
	f.fs.BoolVar(p, name, value, usage)
	if alias != "" {
		f.fs.BoolVar(p, alias, value, usage)
	}
}

// IntVar defines an int flag with specified name, alias, default value, and usage string.
// The argument p points to an int variable in which to store the value of the flag.
// Empty string for alias means no alias will be created.
func (f *FlagSet) IntVar(p *int, name, alias string, value int, usage string) {
	f.fs.IntVar(p, name, value, usage)
	if alias != "" {
		f.fs.IntVar(p, alias, value, usage)
	}
}

// Int64Var defines an int64 flag with specified name, alias, default value, and usage string.
// The argument p points to an int64 variable in which to store the value of the flag.
// Empty string for alias means no alias will be created.
func (f *FlagSet) Int64Var(p *int64, name, alias string, value int64, usage string) {
	f.fs.Int64Var(p, name, value, usage)
	if alias != "" {
		f.fs.Int64Var(p, alias, value, usage)
	}
}

// UintVar defines a uint flag with specified name, alias, default value, and usage string.
// The argument p points to a uint variable in which to store the value of the flag.
// Empty string for alias means no alias will be created.
func (f *FlagSet) UintVar(p *uint, name, alias string, value uint, usage string) {
	f.fs.UintVar(p, name, value, usage)
	if alias != "" {
		f.fs.UintVar(p, alias, value, usage)
	}
}

// Uint64Var defines a uint64 flag with specified name, alias, default value, and usage string.
// The argument p points to a uint64 variable in which to store the value of the flag.
// Empty string for alias means no alias will be created.
func (f *FlagSet) Uint64Var(p *uint64, name, alias string, value uint64, usage string) {
	f.fs.Uint64Var(p, name, value, usage)
	if alias != "" {
		f.fs.Uint64Var(p, alias, value, usage)
	}
}

// StringVar defines a string flag with specified name, alias, default value, and usage string.
// The argument p points to a string variable in which to store the value of the flag.
// Empty string for alias means no alias will be created.
func (f *FlagSet) StringVar(p *string, name, alias string, value string, usage string) {
	f.fs.StringVar(p, name, value, usage)
	if alias != "" {
		f.fs.StringVar(p, alias, value, usage)
	}
}

// Float64Var defines a float64 flag with specified name, alias, default value, and usage string.
// The argument p points to a float64 variable in which to store the value of the flag.
// Empty string for alias means no alias will be created.
func (f *FlagSet) Float64Var(p *float64, name, alias string, value float64, usage string) {
	f.fs.Float64Var(p, name, value, usage)
	if alias != "" {
		f.fs.Float64Var(p, alias, value, usage)
	}
}

// DurationVar defines a time.Duration flag with specified name, alias, default value, and usage string.
// The argument p points to a time.Duration variable in which to store the value of the flag.
// The flag accepts a value acceptable to time.ParseDuration.
// Empty string for alias means no alias will be created.
func (f *FlagSet) DurationVar(p *time.Duration, name, alias string, value time.Duration, usage string) {
	f.fs.DurationVar(p, name, value, usage)
	if alias != "" {
		f.fs.DurationVar(p, alias, value, usage)
	}
}

// Bool defines a bool flag with specified name, alias, default value, and usage string.
// The return value is the address of a bool variable that stores the value of the flag.
// Empty string for alias means no alias will be created.
func (f *FlagSet) Bool(name, alias string, value bool, usage string) *bool {
	p := new(bool)
	f.BoolVar(p, name, alias, value, usage)
	return p
}

// Int defines an int flag with specified name, alias, default value, and usage string.
// The return value is the address of an int variable that stores the value of the flag.
// Empty string for alias means no alias will be created.
func (f *FlagSet) Int(name, alias string, value int, usage string) *int {
	p := new(int)
	f.IntVar(p, name, alias, value, usage)
	return p
}

// Int64 defines an int64 flag with specified name, alias, default value, and usage string.
// The return value is the address of an int64 variable that stores the value of the flag.
// Empty string for alias means no alias will be created.
func (f *FlagSet) Int64(name, alias string, value int64, usage string) *int64 {
	p := new(int64)
	f.Int64Var(p, name, alias, value, usage)
	return p
}

// Uint defines a uint flag with specified name, alias, default value, and usage string.
// The return value is the address of a uint variable that stores the value of the flag.
// Empty string for alias means no alias will be created.
func (f *FlagSet) Uint(name, alias string, value uint, usage string) *uint {
	p := new(uint)
	f.UintVar(p, name, alias, value, usage)
	return p
}

// Uint64Var defines a uint64 flag with specified name, alias, default value, and usage string.
// The argument p points to a uint64 variable in which to store the value of the flag.
// Empty string for alias means no alias will be created.
func (f *FlagSet) Uint64(name, alias string, value uint64, usage string) *uint64 {
	p := new(uint64)
	f.Uint64Var(p, name, alias, value, usage)
	return p
}

// String defines a string flag with specified name, alias, default value, and usage string.
// The return value is the address of a string variable that stores the value of the flag.
// Empty string for alias means no alias will be created.
func (f *FlagSet) String(name, alias string, value string, usage string) *string {
	p := new(string)
	f.StringVar(p, name, alias, value, usage)
	return p
}

// Float64 defines a float64 flag with specified name, alias, default value, and usage string.
// The return value is the address of a float64 variable that stores the value of the flag.
// Empty string for alias means no alias will be created.
func (f *FlagSet) Float64(name, alias string, value float64, usage string) *float64 {
	p := new(float64)
	f.Float64Var(p, name, alias, value, usage)
	return p
}

// Duration defines a time.Duration flag with specified name, alias, default value, and usage string.
// The return value is the address of a time.Duration variable that stores the value of the flag.
// The flag accepts a value acceptable to time.ParseDuration.
// Empty string for alias means no alias will be created.
func (f *FlagSet) Duration(name, alias string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	f.DurationVar(p, name, alias, value, usage)
	return p
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

// Var defines a flag with the specified name and usage string. The type and
// value of the flag are represented by the first argument, of type Value, which
// typically holds a user-defined implementation of Value. For instance, the
// caller could create a flag that turns a comma-separated string into a slice
// of strings by giving the slice the methods of Value; in particular, Set would
// decompose the comma-separated string into the slice.
// Empty string for alias means no alias will be created.
func (f *FlagSet) Var(value flag.Value, name, alias string, usage string) {
	f.fs.Var(value, name, usage)
	if alias != "" {
		f.fs.Var(value, alias, usage)
	}
}
