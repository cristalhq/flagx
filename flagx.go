package flagx

import (
	"encoding"
	"flag"
	"fmt"
	"io"
	"reflect"
	"strings"
	"time"
)

// FlagSet represents a set of flags.
// Flag names must be unique within a FlagSet.
// An attempt to define a flag whose name is already in use will cause a panic.
type FlagSet struct {
	fs      *flag.FlagSet
	aliases map[string]string // a mapping from a flag's name to its alias, empty value means no alias is defined.
}

// NewFlagSet returns new FlagSet.
func NewFlagSet(name string, output io.Writer) *FlagSet {
	fs := flag.NewFlagSet(name, flag.ContinueOnError)
	fs.SetOutput(output)
	return &FlagSet{
		fs:      fs,
		aliases: make(map[string]string),
	}
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
func (f *FlagSet) VisitAll(fn func(*flag.Flag))   { f.fs.VisitAll(fn) }
func (f *FlagSet) Visit(fn func(*flag.Flag))      { f.fs.Visit(fn) }
func (f *FlagSet) Lookup(name string) *flag.Flag  { return f.fs.Lookup(name) }
func (f *FlagSet) Set(name, value string) error   { return f.fs.Set(name, value) }

// Var defines a flag with the specified name and usage string. The type and
// value of the flag are represented by the first argument, of type Value, which
// typically holds a user-defined implementation of Value. For instance, the
// caller could create a flag that turns a comma-separated string into a slice
// of strings by giving the slice the methods of Value; in particular, Set would
// decompose the comma-separated string into the slice.
// Empty string for alias means no alias will be created.
func (f *FlagSet) Var(value flag.Value, name, alias, usage string) {
	f.aliases[name] = alias
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
	f.aliases[name] = alias
	f.fs.Func(name, usage, fn)
	if alias != "" {
		f.fs.Func(alias, usage, fn)
	}
}

// Text defines a flag with the specified name and usage string.
// The argument p must be a pointer to a variable that will hold the value
// of the flag, and p must implement encoding.TextUnmarshaler.
// If the flag is used, the flag value will be passed to p's UnmarshalText method.
// The type of the default value must be the same as the type of p.
// Empty string for alias means no alias will be created.
func (f *FlagSet) Text(p encoding.TextUnmarshaler, name, alias string, value encoding.TextMarshaler, usage string) {
	// TODO(cristaloleg): for Go 1.19 this can be f.fs.TextVar(...)
	f.Var(newTextValue(value, p), name, alias, usage)
}

// Bool defines a bool flag with specified name, alias, default value, and usage string.
// The argument p points to a bool variable in which to store the value of the flag.
// Empty string for alias means no alias will be created.
func (f *FlagSet) Bool(p *bool, name, alias string, value bool, usage string) {
	f.aliases[name] = alias
	f.fs.BoolVar(p, name, value, usage)
	if alias != "" {
		f.fs.BoolVar(p, alias, value, usage)
	}
}

// Int defines an int flag with specified name, alias, default value, and usage string.
// The argument p points to an int variable in which to store the value of the flag.
// Empty string for alias means no alias will be created.
func (f *FlagSet) Int(p *int, name, alias string, value int, usage string) {
	f.aliases[name] = alias
	f.fs.IntVar(p, name, value, usage)
	if alias != "" {
		f.fs.IntVar(p, alias, value, usage)
	}
}

// Int64 defines an int64 flag with specified name, alias, default value, and usage string.
// The argument p points to an int64 variable in which to store the value of the flag.
// Empty string for alias means no alias will be created.
func (f *FlagSet) Int64(p *int64, name, alias string, value int64, usage string) {
	f.aliases[name] = alias
	f.fs.Int64Var(p, name, value, usage)
	if alias != "" {
		f.fs.Int64Var(p, alias, value, usage)
	}
}

// Uint defines a uint flag with specified name, alias, default value, and usage string.
// The argument p points to a uint variable in which to store the value of the flag.
// Empty string for alias means no alias will be created.
func (f *FlagSet) Uint(p *uint, name, alias string, value uint, usage string) {
	f.aliases[name] = alias
	f.fs.UintVar(p, name, value, usage)
	if alias != "" {
		f.fs.UintVar(p, alias, value, usage)
	}
}

// Uint64 defines a uint64 flag with specified name, alias, default value, and usage string.
// The argument p points to a uint64 variable in which to store the value of the flag.
// Empty string for alias means no alias will be created.
func (f *FlagSet) Uint64(p *uint64, name, alias string, value uint64, usage string) {
	f.aliases[name] = alias
	f.fs.Uint64Var(p, name, value, usage)
	if alias != "" {
		f.fs.Uint64Var(p, alias, value, usage)
	}
}

// String defines a string flag with specified name, alias, default value, and usage string.
// The argument p points to a string variable in which to store the value of the flag.
// Empty string for alias means no alias will be created.
func (f *FlagSet) String(p *string, name, alias, value, usage string) {
	f.aliases[name] = alias
	f.fs.StringVar(p, name, value, usage)
	if alias != "" {
		f.fs.StringVar(p, alias, value, usage)
	}
}

// Float64 defines a float64 flag with specified name, alias, default value, and usage string.
// The argument p points to a float64 variable in which to store the value of the flag.
// Empty string for alias means no alias will be created.
func (f *FlagSet) Float64(p *float64, name, alias string, value float64, usage string) {
	f.aliases[name] = alias
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
	f.aliases[name] = alias
	f.fs.DurationVar(p, name, value, usage)
	if alias != "" {
		f.fs.DurationVar(p, alias, value, usage)
	}
}

// BoolSlice defines a slice of bool flag with specified name, alias, default value, separator, and usage string.
// The argument p points to a bool variable in which to store the value of the flag.
// Empty string for alias means no alias will be created.
func (f *FlagSet) BoolSlice(p *[]bool, name, alias string, value []bool, sep, usage string) {
	*p = value
	s := boolSlice{sep: sep, value: p}
	f.Var(s, name, alias, usage)
}

// IntSlice defines a slice of int flag with specified name, alias, default value, separator, and usage string.
// The argument p points to an int variable in which to store the value of the flag.
// Empty string for alias means no alias will be created.
func (f *FlagSet) IntSlice(p *[]int, name, alias string, value []int, sep, usage string) {
	*p = value
	s := intSlice{sep: sep, value: p}
	f.Var(s, name, alias, usage)
}

// Int64Slice defines a slice of int64 flag with specified name, alias, default value, separator, and usage string.
// The argument p points to an int64 variable in which to store the value of the flag.
// Empty string for alias means no alias will be created.
func (f *FlagSet) Int64Slice(p *[]int64, name, alias string, value []int64, sep, usage string) {
	*p = value
	s := int64Slice{sep: sep, value: p}
	f.Var(s, name, alias, usage)
}

// UintSlice defines a slice of uint flag with specified name, alias, default value, separator, and usage string.
// The argument p points to a uint variable in which to store the value of the flag.
// Empty string for alias means no alias will be created.
func (f *FlagSet) UintSlice(p *[]uint, name, alias string, value []uint, sep, usage string) {
	*p = value
	s := uintSlice{sep: sep, value: p}
	f.Var(s, name, alias, usage)
}

// Uint64Slice defines a slice of uint64 flag with specified name, alias, default value, separator, and usage string.
// The argument p points to a uint64 variable in which to store the value of the flag.
// Empty string for alias means no alias will be created.
func (f *FlagSet) Uint64Slice(p *[]uint64, name, alias string, value []uint64, sep, usage string) {
	*p = value
	s := uint64Slice{sep: sep, value: p}
	f.Var(s, name, alias, usage)
}

// StringSlice defines a slice of string flag with specified name, alias, default value, separator, and usage string.
// The argument p points to a string variable in which to store the value of the flag.
// Empty string for alias means no alias will be created.
func (f *FlagSet) StringSlice(p *[]string, name, alias string, value []string, sep, usage string) {
	*p = value
	s := stringSlice{sep: sep, value: p}
	f.Var(s, name, alias, usage)
}

// Float64Slice defines a slice of float64 flag with specified name, alias, default value, separator, and usage string.
// The argument p points to a float64 variable in which to store the value of the flag.
// Empty string for alias means no alias will be created.
func (f *FlagSet) Float64Slice(p *[]float64, name, alias string, value []float64, sep, usage string) {
	*p = value
	s := float64Slice{sep: sep, value: p}
	f.Var(s, name, alias, usage)
}

// DurationSlice defines a slice of time.Duration flag with specified name, alias, default value, separator, and usage string.
// The argument p points to a time.Duration variable in which to store the value of the flag.
// The flag accepts a value acceptable to time.ParseDuration.
// Empty string for alias means no alias will be created.
func (f *FlagSet) DurationSlice(p *[]time.Duration, name, alias string, value []time.Duration, sep, usage string) {
	*p = value
	s := durationSlice{sep: sep, value: p}
	f.Var(s, name, alias, usage)
}

// IntSet defines a set of int flag with specified name, alias, default value, and usage string.
// The argument p points to an int variable in which to store the value of the flag.
// Empty string for alias means no alias will be created.
func (f *FlagSet) IntSet(p *map[int]struct{}, name, alias string, value map[int]struct{}, usage string) {
	f.aliases[name] = alias
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
	f.aliases[name] = alias
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
	f.aliases[name] = alias
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
	f.aliases[name] = alias
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
	f.aliases[name] = alias
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
	f.aliases[name] = alias
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
	f.aliases[name] = alias
	var sd SetOfDuration = value
	*p = map[time.Duration]struct{}(sd)
	f.fs.Var(&sd, name, usage)
	if alias != "" {
		f.fs.Var(&sd, alias, usage)
	}
}

// PrintDefaults prints, to standard error unless configured otherwise, the
// default values of all defined command-line flags in the set.
func (f *FlagSet) PrintDefaults() {
	// NOTE(junk1tm): copy-pasted from flag.PrintDefaults with a few modifications to support aliases.

	var isZeroValueErrs []error
	f.VisitAll(func(fl *flag.Flag) {
		if _, ok := f.aliases[fl.Name]; !ok {
			// The flag is an alias, do not print it separately.
			return
		}
		var b strings.Builder
		fmt.Fprintf(&b, "  -%s", fl.Name) // Two spaces before -; see next two comments.
		if alias := f.aliases[fl.Name]; alias != "" {
			fmt.Fprintf(&b, " (-%s)", alias)
		}
		name, usage := flag.UnquoteUsage(fl)
		if len(name) > 0 {
			b.WriteString(" ")
			b.WriteString(name)
		}
		// Boolean flags of one ASCII letter are so common we
		// treat them specially, putting their usage on the same line.
		if b.Len() <= 4 { // space, space, '-', 'x'.
			b.WriteString("\t")
		} else {
			// Four spaces before the tab triggers good alignment
			// for both 4- and 8-space tab stops.
			b.WriteString("\n    \t")
		}
		b.WriteString(strings.ReplaceAll(usage, "\n", "\n    \t"))

		// Print the default value only if it differs to the zero value
		// for this flag type.
		if isZero, err := isZeroValue(fl, fl.DefValue); err != nil {
			isZeroValueErrs = append(isZeroValueErrs, err)
		} else if !isZero {
			// HACK(junk1tm): flag.stringValue is unexported, so we have to compare the type's name.
			if fmt.Sprintf("%T", fl.Value) == "*flag.stringValue" {
				// put quotes on the value
				fmt.Fprintf(&b, " (default %q)", fl.DefValue)
			} else {
				fmt.Fprintf(&b, " (default %v)", fl.DefValue)
			}
		}
		fmt.Fprint(f.fs.Output(), b.String(), "\n")
	})
	// If calling String on any zero flag.Values triggered a panic, print
	// the messages after the full set of defaults so that the programmer
	// knows to fix the panic.
	if errs := isZeroValueErrs; len(errs) > 0 {
		fmt.Fprintln(f.fs.Output())
		for _, err := range errs {
			fmt.Fprintln(f.fs.Output(), err)
		}
	}
}

func isZeroValue(fl *flag.Flag, value string) (ok bool, err error) {
	// NOTE(junk1tm): copy-pasted from flag.isZeroValue as a part of flag.PrintDefaults.

	// Build a zero value of the flag's Value type, and see if the
	// result of calling its String method equals the value passed in.
	// This works unless the Value type is itself an interface type.
	typ := reflect.TypeOf(fl.Value)
	var z reflect.Value
	if typ.Kind() == reflect.Pointer {
		z = reflect.New(typ.Elem())
	} else {
		z = reflect.Zero(typ)
	}
	// Catch panics calling the String method, which shouldn't prevent the
	// usage message from being printed, but that we should report to the
	// user so that they know to fix their code.
	defer func() {
		if e := recover(); e != nil {
			if typ.Kind() == reflect.Pointer {
				typ = typ.Elem()
			}
			err = fmt.Errorf("panic calling String method on zero %v for flag %s: %v", typ, fl.Name, e)
		}
	}()
	return value == z.Interface().(flag.Value).String(), nil
}
