// Copyright (c) 2018 Dean Jackson <deanishe@deanishe.net>
// MIT Licence applies http://opensource.org/licenses/MIT

package env // import "go.deanishe.net/env"

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var (
	// System retrieves values from the system environment.
	System Env = systemEnv{}
	// Default Reader, which reads from the system environment.
	system = Reader{System}
)

// systemEnv reads values from the real environment
type systemEnv struct{}

func (env systemEnv) Lookup(key string) (string, bool) {
	return os.LookupEnv(key)
}

// Env is the data source for bindings and lookup. It is an optional
// parameter to Bind(). By specifying a custom Env, it's possible
// to populate a struct from an alternative source.
//
// The demo program in _examples/docopt implements a custom Env
// to populate a struct from docopt options via Bind().
type Env interface {
	// Lookup retrieves the value of the variable named by key.
	//
	// It follows the same semantics as os.LookupEnv(). If a variable
	// is unset, the boolean will be false. If a variable is set, the
	// boolean will be true, but the variable may still be an empty
	// string.
	Lookup(key string) (string, bool)
}

// MapEnv is a string: string mapping that implements Env.
type MapEnv map[string]string

// Lookup implements Env.
func (env MapEnv) Lookup(key string) (string, bool) {
	s, ok := env[key]
	return s, ok
}

// Reader converts values from Env into other types.
type Reader struct {
	env Env
}

// New creates a new Reader based on Env.
func New(env Env) Reader {
	return Reader{env}
}

// Get returns the value for envvar "key".
// It accepts one optional "fallback" argument. If no envvar is set,
// returns fallback or an empty string.
//
// If a variable is set, but empty, its value is used.
func Get(key string, fallback ...string) string {
	return system.Get(key, fallback...)
}

// Get returns the value for envvar "key".
// It accepts one optional "fallback" argument. If no envvar is set,
// returns fallback or an empty string.
//
// If a variable is set, but empty, its value is used.
func (r Reader) Get(key string, fallback ...string) string {
	var fb string
	if len(fallback) > 0 {
		fb = fallback[0]
	}

	s, ok := r.env.Lookup(key)
	if !ok {
		return fb
	}
	return s
}

// GetString is a synonym for Get.
func GetString(key string, fallback ...string) string {
	return system.GetString(key, fallback...)
}

// GetString is a synonym for Get.
func (r Reader) GetString(key string, fallback ...string) string {
	return r.Get(key, fallback...)
}

// GetInt returns the value for envvar "key" as an int.
// It accepts one optional "fallback" argument. If no
// envvar is set, returns fallback or 0.
//
// Values are parsed with strconv.ParseInt(). If strconv.ParseInt()
// fails, tries to parse the number with strconv.ParseFloat() and
// truncate it to an int.
func GetInt(key string, fallback ...int) int {
	return system.GetInt(key, fallback...)
}

// GetInt returns the value for envvar "key" as an int.
// It accepts one optional "fallback" argument. If no
// envvar is set, returns fallback or 0.
//
// Values are parsed with strconv.ParseInt(). If strconv.ParseInt()
// fails, tries to parse the number with strconv.ParseFloat() and
// truncate it to an int.
func (r Reader) GetInt(key string, fallback ...int) int {
	var fb int
	if len(fallback) > 0 {
		fb = fallback[0]
	}
	s, ok := r.env.Lookup(key)
	if !ok {
		return fb
	}

	i, err := parseInt(s)
	if err != nil {
		return fb
	}
	return i
}

// GetUint returns the value for envvar "key" as an int.
// It accepts one optional "fallback" argument. If no
// envvar is set, returns fallback or 0.
//
// Values are parsed with strconv.ParseUint(). If strconv.ParseUint()
// fails, tries to parse the number with strconv.ParseFloat() and
// truncate it to a uint.
func GetUint(key string, fallback ...uint) uint {
	return system.GetUint(key, fallback...)
}

// GetUint returns the value for envvar "key" as an int.
// It accepts one optional "fallback" argument. If no
// envvar is set, returns fallback or 0.
//
// Values are parsed with strconv.ParseUint(). If strconv.ParseUint()
// fails, tries to parse the number with strconv.ParseFloat() and
// truncate it to a uint.
func (r Reader) GetUint(key string, fallback ...uint) uint {
	var fb uint
	if len(fallback) > 0 {
		fb = fallback[0]
	}
	s, ok := r.env.Lookup(key)
	if !ok {
		return fb
	}

	i, err := parseUint(s)
	if err != nil {
		return fb
	}
	return i
}

// GetFloat returns the value for envvar "key" as a float.
// It accepts one optional "fallback" argument. If no
// envvar is set, returns fallback or 0.0.
//
// Values are parsed with strconv.ParseFloat().
func GetFloat(key string, fallback ...float64) float64 {
	return system.GetFloat(key, fallback...)
}

// GetFloat returns the value for envvar "key" as a float.
// It accepts one optional "fallback" argument. If no
// envvar is set, returns fallback or 0.0.
//
// Values are parsed with strconv.ParseFloat().
func (r Reader) GetFloat(key string, fallback ...float64) float64 {
	var fb float64
	if len(fallback) > 0 {
		fb = fallback[0]
	}
	s, ok := r.env.Lookup(key)
	if !ok {
		return fb
	}

	n, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return fb
	}
	return n
}

// GetDuration returns the value for envvar "key" as a time.Duration.
// It accepts one optional "fallback" argument. If no
// envvar is set, returns fallback or 0.
//
// Values are parsed with time.ParseDuration().
func GetDuration(key string, fallback ...time.Duration) time.Duration {
	return system.GetDuration(key, fallback...)
}

// GetDuration returns the value for envvar "key" as a time.Duration.
// It accepts one optional "fallback" argument. If no
// envvar is set, returns fallback or 0.
//
// Values are parsed with time.ParseDuration().
func (r Reader) GetDuration(key string, fallback ...time.Duration) time.Duration {
	var fb time.Duration
	if len(fallback) > 0 {
		fb = fallback[0]
	}
	s, ok := r.env.Lookup(key)
	if !ok {
		return fb
	}

	d, err := time.ParseDuration(s)
	if err != nil {
		return fb
	}
	return d
}

// GetBool returns the value for envvar "key" as a boolean.
// It accepts one optional "fallback" argument. If no
// envvar is set, returns fallback or false.
//
// Values are parsed with strconv.ParseBool().
func GetBool(key string, fallback ...bool) bool {
	return system.GetBool(key, fallback...)
}

// GetBool returns the value for envvar "key" as a boolean.
// It accepts one optional "fallback" argument. If no
// envvar is set, returns fallback or false.
//
// Values are parsed with strconv.ParseBool().
func (r Reader) GetBool(key string, fallback ...bool) bool {
	var fb bool
	if len(fallback) > 0 {
		fb = fallback[0]
	}
	s, ok := r.env.Lookup(key)
	if !ok {
		return fb
	}

	b, err := strconv.ParseBool(s)
	if err != nil {
		return fb
	}
	return b
}

// parse an int, falling back to parsing it as a float
func parseInt(s string) (int, error) {
	i, err := strconv.ParseInt(s, 10, 32)
	if err == nil {
		return int(i), nil
	}

	// Try to parse as float, then convert
	n, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid int: %s", s)
	}
	return int(n), nil
}

// parse an int, falling back to parsing it as a float
func parseUint(s string) (uint, error) {
	if i, err := strconv.ParseUint(s, 10, 32); err == nil {
		return uint(i), nil
	}

	// Try to parse as float, then convert
	if f, err := strconv.ParseFloat(s, 64); err == nil {
		if f < 0 {
			return 0, fmt.Errorf("less than zero: %s", s)
		}
		return uint(f), nil
	}
	return 0, fmt.Errorf("invalid int: %s", s)
}
