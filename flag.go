package flag

import (
	f "flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	envPrefix string
)

// Usage prints a usage message documenting all defined command-line flags
// to CommandLine's output, which by default is os.Stderr.
// It is called when an error occurs while parsing flags.
// The function is a variable that may be changed to point to a custom function.
// By default it prints a simple header and calls PrintDefaults; for details about the
// format of the output and how to control it, see the documentation for PrintDefaults.
// Custom usage functions may choose to exit the program; by default exiting
// happens anyway as the command line's error handling strategy is set to
// ExitOnError.
func Usage() {
	f.Usage()
}

// SetEnvPrefix sets the prefix for environmental default values
// Must be called before using any flag definition functions (or to separate values,
// then it should be grouped)
func SetEnvPrefix(prefix string) {
	envPrefix = prefix
}

// envNameForFlagName creates an environment variable from a flag name
func envNameForFlagName(name string) string {
	envName := name
	if len(envPrefix) > 0 {
		envName = fmt.Sprintf("%s_%s", envPrefix, name)
	}
	return strings.ToUpper(strings.Replace(envName, "-", "_", -1))
}

// boolFromEnv returns true if the environment variable is set to true
// if not found returns the default value
func boolFromEnv(name string, value bool) bool {
	val, found := os.LookupEnv(envNameForFlagName(name))
	if !found {
		return value
	}
	return val == "true"
}

// Bool defines a bool flag with specified name, default value, and usage string.
// The return value is the address of a bool variable that stores the value of the flag.
func Bool(name string, value bool, usage string) *bool {
	return f.Bool(name, boolFromEnv(name, value), usage)
}

// BoolVar defines a bool flag with specified name, default value, and usage string.
// The argument p points to a bool variable in which to store the value of the flag.
func BoolVar(p *bool, name string, value bool, usage string) {
	f.BoolVar(p, name, boolFromEnv(name, value), usage)
}

// durationFromEnv returns parsed duration from environment variable. On error returning default value
func durationFromEnv(name string, value time.Duration) time.Duration {
	val, found := os.LookupEnv(envNameForFlagName(name))
	if !found {
		return value
	}
	dur, err := time.ParseDuration(val)
	if err != nil {
		return value
	}
	return dur
}

// Duration defines a time.Duration flag with specified name, default value, and usage string.
// The return value is the address of a time.Duration variable that stores the value of the flag.
// The flag accepts a value acceptable to time.ParseDuration.
func Duration(name string, value time.Duration, usage string) *time.Duration {
	return f.Duration(name, durationFromEnv(name, value), usage)
}

// DurationVar defines a time.Duration flag with specified name, default value, and usage string.
// The argument p points to a time.Duration variable in which to store the value of the flag.
// The flag accepts a value acceptable to time.ParseDuration.
func DurationVar(p *time.Duration, name string, value time.Duration, usage string) {
	f.DurationVar(p, name, durationFromEnv(name, value), usage)
}

// float64FromEnv returns parsed float64 from environment variable. On error returning default value
func float64FromEnv(name string, value float64) float64 {
	val, found := os.LookupEnv(envNameForFlagName(name))
	if !found {
		return value
	}
	fl, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return value
	}
	return fl
}

// Float64 defines a float64 flag with specified name, default value, and usage string.
// The return value is the address of a float64 variable that stores the value of the flag.
func Float64(name string, value float64, usage string) *float64 {
	return f.Float64(name, float64FromEnv(name, value), usage)
}

// Float64Var defines a float64 flag with specified name, default value, and usage string.
// The argument p points to a float64 variable in which to store the value of the flag.
func Float64Var(p *float64, name string, value float64, usage string) {
	f.Float64Var(p, name, float64FromEnv(name, value), usage)
}

// Func defines a flag with the specified name and usage string. Each time the flag is seen,
// fn is called with the value of the flag. If fn returns a non-nil error, it will be treated
// as a flag value parsing error.
func Func(name, usage string, fn func(string) error) {
	f.Func(name, usage, fn)
}

// int64FromEnv returns parsed int64 from environment variable. On error returning default value
func int64FromEnv(name string, value int64) int64 {
	val, found := os.LookupEnv(envNameForFlagName(name))
	if !found {
		return value
	}
	i, err := strconv.ParseInt(val, 0, 64)
	if err != nil {
		return value
	}
	return i
}

// Int defines an int flag with specified name, default value, and usage string.
// The return value is the address of an int variable that stores the value of the flag.
func Int(name string, value int, usage string) *int {
	return f.Int(name, int(int64FromEnv(name, int64(value))), usage)
}

// Int64 defines an int64 flag with specified name, default value, and usage string.
// The return value is the address of an int64 variable that stores the value of the flag.
func Int64(name string, value int64, usage string) *int64 {
	return f.Int64(name, int64FromEnv(name, value), usage)
}

// Int64Var defines an int64 flag with specified name, default value, and usage string.
// The argument p points to an int64 variable in which to store the value of the flag.
func Int64Var(p *int64, name string, value int64, usage string) {
	f.Int64Var(p, name, int64FromEnv(name, value), usage)
}

// IntVar defines an int flag with specified name, default value, and usage string.
// The argument p points to an int variable in which to store the value of the flag.
func IntVar(p *int, name string, value int, usage string) {
	f.IntVar(p, name, int(int64FromEnv(name, int64(value))), usage)
}

// NArg is the number of arguments remaining after flags have been processed.
func NArg() int {
	return f.NArg()
}

// NFlag returns the number of command-line flags that have been set.
func NFlag() int {
	return f.NFlag()
}

// Parse parses the command-line flags from os.Args[1:]. Must be called after all flags are defined and before flags are accessed by the program.
func Parse() {
	f.Parse()
}

// Parsed reports whether the command-line flags have been parsed.
func Parsed() bool {
	return f.Parsed()
}

// PrintDefaults prints, to standard error unless configured otherwise,
// a usage message showing the default settings of all defined
// command-line flags.
// For an integer valued flag x, the default output has the form
//
//	-x int
//		usage-message-for-x (default 7)
//
// The usage message will appear on a separate line for anything but
// a bool flag with a one-byte name. For bool flags, the type is
// omitted and if the flag name is one byte the usage message appears
// on the same line. The parenthetical default is omitted if the
// default is the zero value for the type. The listed type, here int,
// can be changed by placing a back-quoted name in the flag's usage
// string; the first such item in the message is taken to be a parameter
// name to show in the message and the back quotes are stripped from
// the message when displayed. For instance, given
//
//	flag.String("I", "", "search `directory` for include files")
//
// the output will be
//
//	-I directory
//		search directory for include files.
//
// To change the destination for flag messages, call CommandLine.SetOutput.
func PrintDefaults() {
	f.PrintDefaults()
}

// Set sets the value of the named command-line flag.
func Set(name, value string) error {
	return f.Set(name, value)
}

// stringFromEnv returns environment variable. On error returning default value
func stringFromEnv(name string, value string) string {
	val, found := os.LookupEnv(envNameForFlagName(name))
	if !found {
		return value
	}
	return val
}

// String defines a string flag with specified name, default value, and usage string.
//The return value is the address of a string variable that stores the value of the flag.
func String(name string, value string, usage string) *string {
	return f.String(name, stringFromEnv(name, value), usage)
}

// StringVar defines a string flag with specified name, default value, and usage string.
// The argument p points to a string variable in which to store the value of the flag.
func StringVar(p *string, name string, value string, usage string) {
	f.StringVar(p, name, stringFromEnv(name, value), usage)
}

// uint64FromEnv returns parsed int64 from environment variable. On error returning default value
func uint64FromEnv(name string, value uint64) uint64 {
	val, found := os.LookupEnv(envNameForFlagName(name))
	if !found {
		return value
	}
	i, err := strconv.ParseUint(val, 0, 64)
	if err != nil {
		return value
	}
	return i
}

// Uint defines a uint flag with specified name, default value, and usage string.
// The return value is the address of a uint variable that stores the value of the flag.
func Uint(name string, value uint, usage string) *uint {
	return f.Uint(name, uint(uint64FromEnv(name, uint64(value))), usage)
}

// Uint64 defines a uint64 flag with specified name, default value, and usage string.
// The return value is the address of a uint64 variable that stores the value of the flag.
func Uint64(name string, value uint64, usage string) *uint64 {
	return f.Uint64(name, uint64FromEnv(name, value), usage)
}

// Uint64Var defines a uint64 flag with specified name, default value, and usage string.
//The argument p points to a uint64 variable in which to store the value of the flag.
func Uint64Var(p *uint64, name string, value uint64, usage string) {
	f.Uint64Var(p, name, uint64FromEnv(name, value), usage)
}

// UintVar defines a uint flag with specified name, default value, and usage string.
// The argument p points to a uint variable in which to store the value of the flag.
func UintVar(p *uint, name string, value uint, usage string) {
	f.Uint(name, uint(uint64FromEnv(name, uint64(value))), usage)
}

// UnquoteUsage extracts a back-quoted name from the usage string for a flag and returns it and the un-quoted usage.
// Given "a `name` to show" it returns ("name", "a name to show"). If there are no back quotes, the name is an
// educated guess of the type of the flag's value, or the empty string if the flag is boolean.
func UnquoteUsage(usageString string) (name string, usage string) {
	return f.UnquoteUsage(&f.Flag{Usage: usageString})
}

// Var defines a flag with the specified name and usage string. The type and value of the flag are represented
// by the first argument, of type Value, which typically holds a user-defined implementation of Value. For instance,
// the caller could create a flag that turns a comma-separated string into a slice of strings by giving the slice
// the methods of Value; in particular, Set would decompose the comma-separated string into the slice.
func Var(value f.Value, name string, usage string) {
	f.Var(value, name, usage)
}

// Visit visits the command-line flags in lexicographical order, calling fn for each. It visits only those
// flags that have been set.
func Visit(fn func(*f.Flag)) {
	f.Visit(fn)
}

// VisitAll visits the command-line flags in lexicographical order, calling fn for each. It visits all flags,
// even those not set.
func VisitAll(fn func(*f.Flag)) {
	f.VisitAll(fn)
}
