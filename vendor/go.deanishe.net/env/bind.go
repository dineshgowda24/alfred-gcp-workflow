// Copyright (c) 2018 Dean Jackson <deanishe@deanishe.net>
// MIT Licence applies http://opensource.org/licenses/MIT

package env

import (
	"encoding"
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Errors returned by Dump and Bind if they are called with inappropriate values. Bind() requires a pointer to a struct,
// while Dump requires either a struct or a pointer to a struct.
var (
	ErrNotStruct    = errors.New("not a struct")
	ErrNotStructPtr = errors.New("not a pointer to a struct")
)

// function that can parse a string into a type's native values.
type parseFunc func(s string) (interface{}, error)

// return function from kindParsers/typeParsers appropriate for fieldType.
func getParseFunc(fieldType reflect.Type) (fun parseFunc, ok bool) {
	if fun, ok = typeParsers[fieldType]; ok {
		return
	}
	fun, ok = kindParsers[fieldType.Kind()]
	return
}

// Functions to parse strings into type-appropriate values.
var (
	kindParsers = map[reflect.Kind]parseFunc{
		reflect.Bool: func(s string) (interface{}, error) {
			return strconv.ParseBool(s)
		},
		reflect.String: func(s string) (interface{}, error) {
			return s, nil
		},
		reflect.Int: func(s string) (interface{}, error) {
			i, err := strconv.ParseInt(s, 10, 32)
			return int(i), err
		},
		reflect.Int8: func(s string) (interface{}, error) {
			i, err := strconv.ParseInt(s, 10, 8)
			return int8(i), err
		},
		reflect.Int16: func(s string) (interface{}, error) {
			i, err := strconv.ParseInt(s, 10, 16)
			return int16(i), err
		},
		reflect.Int32: func(s string) (interface{}, error) {
			i, err := strconv.ParseInt(s, 10, 32)
			return int32(i), err
		},
		reflect.Int64: func(s string) (interface{}, error) {
			return strconv.ParseInt(s, 10, 64)
		},
		reflect.Uint: func(s string) (interface{}, error) {
			i, err := strconv.ParseUint(s, 10, 32)
			return uint(i), err
		},
		reflect.Uint8: func(s string) (interface{}, error) {
			i, err := strconv.ParseUint(s, 10, 8)
			return uint8(i), err
		},
		reflect.Uint16: func(s string) (interface{}, error) {
			i, err := strconv.ParseUint(s, 10, 16)
			return uint16(i), err
		},
		reflect.Uint32: func(s string) (interface{}, error) {
			i, err := strconv.ParseUint(s, 10, 32)
			return uint32(i), err
		},
		reflect.Uint64: func(s string) (interface{}, error) {
			return strconv.ParseUint(s, 10, 64)
		},
		reflect.Float32: func(s string) (interface{}, error) {
			n, err := strconv.ParseFloat(s, 32)
			return float32(n), err
		},
		reflect.Float64: func(s string) (interface{}, error) {
			return strconv.ParseFloat(s, 64)
		},
	}
	typeParsers = map[reflect.Type]parseFunc{
		reflect.TypeOf(url.URL{}): func(s string) (interface{}, error) {
			u, err := url.Parse(s)
			if err != nil {
				return nil, fmt.Errorf("invalid URL %q: %w", s, err)
			}
			return *u, nil
		},
		reflect.TypeOf(time.Nanosecond): func(s string) (interface{}, error) {
			d, err := time.ParseDuration(s)
			if err != nil {
				return nil, fmt.Errorf("invalid duration %q: %w", s, err)
			}
			return d, nil
		},
	}
)

// ErrUnsupported is returned by Bind if a field of an unsupported type is tagged for binding.
// Unsupported fields that are not tagged are ignored.
type ErrUnsupported string

// implements error.Error.
func (err ErrUnsupported) Error() string {
	return "unsupported type: " + string(err)
}

// Bind populates the fields of a struct from environment variables.
//
// Variables are mapped to fields using `env:"..."` tags, and the
// struct is populated by passing it to Bind(). Unset or empty
// environment variables are ignored.
//
// Untagged fields have a default environment variable assigned to
// them. See VarName() for details of how names are generated.
//
// Bind accepts an optional Env argument. If provided, values will
// be looked up via that Env instead of the program's environment.
func Bind(v interface{}, env ...Env) error {
	var e Env
	if len(env) > 0 {
		e = env[0]
	} else {
		e = &systemEnv{}
	}

	return bind(v, e)
}

// populate struct v from Env.
func bind(v interface{}, env Env) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr {
		return ErrNotStructPtr
	}
	rv = rv.Elem()
	if rv.Kind() != reflect.Struct {
		return ErrNotStructPtr
	}
	return populate(rv, env)
}

// set Value rv from Env.
func populate(rv reflect.Value, env Env) error {
	rvType := rv.Type()

	for i := 0; i < rvType.NumField(); i++ {
		fieldVal := rv.Field(i)
		if !fieldVal.CanSet() {
			continue
		}

		// pointer fieldVal
		if fieldVal.Kind() == reflect.Ptr && !fieldVal.IsNil() {
			if err := bind(fieldVal.Interface(), env); err != nil {
				return err
			}
			continue
		}

		// embedded struct
		if fieldVal.Kind() == reflect.Struct && fieldVal.CanAddr() && fieldVal.Type().Name() == "" {
			if err := bind(fieldVal.Addr().Interface(), env); err != nil {
				return err
			}
			continue
		}

		field := rvType.Field(i)
		key := getFieldKey(field)
		if key == "-" {
			continue
		}
		value, _ := env.Lookup(key)

		if value == "" {
			if fieldVal.Kind() == reflect.Struct {
				if err := populate(fieldVal, env); err != nil {
					return err
				}
			}
			continue
		}
		if err := setField(fieldVal, field, value); err != nil {
			return err
		}
	}

	return nil
}

func getFieldKey(field reflect.StructField) string {
	key := field.Tag.Get("env")
	if key == "" {
		key = VarName(field.Name)
	}
	return key
}

// populate Value rv with value parsed from string.
func setField(rv reflect.Value, field reflect.StructField, value string) error {
	if rv.Kind() == reflect.Slice {
		return setSlice(rv, field, value)
	}

	// ensure pointer values are non-nil
	if rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			rv.Set(reflect.New(rv.Type().Elem()))
		}
	}

	if tm := asTextUnmarshaller(rv); tm != nil {
		return tm.UnmarshalText([]byte(value))
	}

	fieldType := field.Type
	if fieldType.Kind() == reflect.Ptr {
		fieldType = fieldType.Elem()
		rv = rv.Elem()
	}

	if parseFn, ok := getParseFunc(fieldType); ok {
		val, err := parseFn(value)
		if err != nil {
			return err
		}
		rv.Set(reflect.ValueOf(val).Convert(fieldType))
		return nil
	}

	return ErrUnsupported(fieldType.String())
}

// populate a slice with multiple values parsed from string.
func setSlice(rv reflect.Value, field reflect.StructField, value string) error {
	parts := strings.Split(value, ",")

	fieldType := field.Type.Elem()
	if fieldType.Kind() == reflect.Ptr {
		fieldType = fieldType.Elem()
	}

	if _, ok := reflect.New(fieldType).Interface().(encoding.TextUnmarshaler); ok {
		return unmarshalSlice(rv, parts)
	}

	parseFn, ok := getParseFunc(fieldType)
	if !ok {
		return ErrUnsupported(fieldType.String())
	}

	values := reflect.MakeSlice(field.Type, 0, len(parts))
	for _, s := range parts {
		v, err := parseFn(s)
		if err != nil {
			return err
		}
		val := reflect.ValueOf(v).Convert(fieldType)
		if field.Type.Elem().Kind() == reflect.Ptr {
			val = reflect.New(fieldType)
			val.Elem().Set(reflect.ValueOf(v).Convert(fieldType))
		}
		values = reflect.Append(values, val)
	}

	rv.Set(values)
	return nil
}

func asTextUnmarshaller(rv reflect.Value) encoding.TextUnmarshaler {
	if tm, ok := rv.Interface().(encoding.TextUnmarshaler); ok {
		return tm
	}
	return nil
}

func unmarshalSlice(rv reflect.Value, parts []string) error {
	itemType := rv.Type().Elem()
	values := reflect.MakeSlice(reflect.SliceOf(itemType), len(parts), len(parts))
	for i, s := range parts {
		sv := values.Index(i)
		kind := sv.Kind()
		if kind == reflect.Ptr {
			sv = reflect.New(itemType.Elem())
		}
		tm := sv.Interface().(encoding.TextUnmarshaler)
		if err := tm.UnmarshalText([]byte(s)); err != nil {
			return err
		}
		if kind == reflect.Ptr {
			values.Index(i).Set(sv)
		}
	}

	rv.Set(values)
	return nil
}

// VarName generates an environment variable name from a field name.
// This is documented to show how the automatic names are generated.
func VarName(name string) string {
	if !isCamelCase(name) {
		return strings.ToUpper(name)
	}
	return splitCamelCase(name)
}

func isCamelCase(s string) bool {
	if ok, _ := regexp.MatchString(".*[a-z]+[0-9_]*[A-Z]+.*", s); ok {
		return true
	}
	if ok, _ := regexp.MatchString("[A-Z][A-Z][A-Z]+[0-9_]*[a-z]+.*", s); ok {
		return true
	}
	return false
}

func splitCamelCase(name string) string {
	var (
		i     int
		re    *regexp.Regexp
		rest  string
		words []string
	)

	rest = name

	// Start with 3 or more capital letters.
	re = regexp.MustCompile("[A-Z]+([A-Z])[0-9]*[a-z]")
	for {
		if idx := re.FindStringSubmatchIndex(rest); idx != nil {
			i = idx[2]
			s := rest[:i]
			rest = rest[i:]
			words = append(words, s)
		} else {
			break
		}
	}

	re = regexp.MustCompile("[a-z][0-9_]*([A-Z])")
	for {
		if idx := re.FindStringSubmatchIndex(rest); idx != nil {
			i = idx[2]
			s := rest[:i]
			rest = rest[i:]
			words = append(words, s)
		} else {
			break
		}
	}

	if rest != "" {
		words = append(words, rest)
	}

	if len(words) > 0 {
		s := strings.ToUpper(strings.Join(words, "_"))
		return s
	}

	return ""
}
