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

func (f *FlagSet) BoolVar(p *bool, name, alias string, value bool, usage string) {
	f.fs.BoolVar(p, name, value, usage)
	if alias != "" {
		f.fs.BoolVar(p, alias, value, usage)
	}
}
func (f *FlagSet) IntVar(p *int, name, alias string, value int, usage string) {
	f.fs.IntVar(p, name, value, usage)
	if alias != "" {
		f.fs.IntVar(p, alias, value, usage)
	}
}
func (f *FlagSet) Int64Var(p *int64, name, alias string, value int64, usage string) {
	f.fs.Int64Var(p, name, value, usage)
	if alias != "" {
		f.fs.Int64Var(p, alias, value, usage)
	}
}
func (f *FlagSet) UintVar(p *uint, name, alias string, value uint, usage string) {
	f.fs.UintVar(p, name, value, usage)
	if alias != "" {
		f.fs.UintVar(p, alias, value, usage)
	}
}
func (f *FlagSet) Uint64Var(p *uint64, name, alias string, value uint64, usage string) {
	f.fs.Uint64Var(p, name, value, usage)
	if alias != "" {
		f.fs.Uint64Var(p, alias, value, usage)
	}
}
func (f *FlagSet) StringVar(p *string, name, alias string, value string, usage string) {
	f.fs.StringVar(p, name, value, usage)
	if alias != "" {
		f.fs.StringVar(p, alias, value, usage)
	}
}
func (f *FlagSet) Float64Var(p *float64, name, alias string, value float64, usage string) {
	f.fs.Float64Var(p, name, value, usage)
	if alias != "" {
		f.fs.Float64Var(p, alias, value, usage)
	}
}
func (f *FlagSet) DurationVar(p *time.Duration, name, alias string, value time.Duration, usage string) {
	f.fs.DurationVar(p, name, value, usage)
	if alias != "" {
		f.fs.DurationVar(p, alias, value, usage)
	}
}

func (f *FlagSet) Bool(name, alias string, value bool, usage string) *bool {
	p := new(bool)
	f.BoolVar(p, name, alias, value, usage)
	return p
}
func (f *FlagSet) Int(name, alias string, value int, usage string) *int {
	p := new(int)
	f.IntVar(p, name, alias, value, usage)
	return p
}
func (f *FlagSet) Int64(name, alias string, value int64, usage string) *int64 {
	p := new(int64)
	f.Int64Var(p, name, alias, value, usage)
	return p
}
func (f *FlagSet) Uint(name, alias string, value uint, usage string) *uint {
	p := new(uint)
	f.UintVar(p, name, alias, value, usage)
	return p
}
func (f *FlagSet) Uint64(name, alias string, value uint64, usage string) *uint64 {
	p := new(uint64)
	f.Uint64Var(p, name, alias, value, usage)
	return p
}
func (f *FlagSet) String(name, alias string, value string, usage string) *string {
	p := new(string)
	f.StringVar(p, name, alias, value, usage)
	return p
}
func (f *FlagSet) Float64(name, alias string, value float64, usage string) *float64 {
	p := new(float64)
	f.Float64Var(p, name, alias, value, usage)
	return p
}
func (f *FlagSet) Duration(name, alias string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	f.DurationVar(p, name, alias, value, usage)
	return p
}

func (f *FlagSet) Var(value flag.Value, name, alias string, usage string) {
	f.fs.Var(value, name, usage)
	if alias != "" {
		f.fs.Var(value, alias, usage)
	}
}
func (f *FlagSet) Func(name, alias, usage string, fn func(string) error) {
	f.fs.Func(name, usage, fn)
	if alias != "" {
		f.fs.Func(name, alias, fn)
	}
}