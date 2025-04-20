
go-env
======

![Build Status][github-status-icon]
[![Go Report Card][goreport-icon]][goreport-link]
[![Codacy coverage][coverage-icon]][codacy-link]
[![GitHub licence][licence-icon]][licence-link]
[![GoDoc][godoc-icon]][godoc-link]

Access environment variables from Go, and populate structs from them.

<!-- MarkdownTOC autolink=true autoanchor=true -->

- [Usage](#usage)
    - [Access environment variables](#access-environment-variables)
    - [Binding](#binding)
    - [Customisation](#customisation)
    - [Dumping](#dumping)
- [Installation](#installation)
- [Documentation](#documentation)
- [Licence](#licence)

<!-- /MarkdownTOC -->


<a id="usage"></a>
Usage
-----

Import path is `go.deanishe.net/env`.

You can directly access environment variables, or populate your structs from them using struct tags and `env.Bind()`.


<a id="access-environment-variables"></a>
### Access environment variables ###

Read `int`, `float64`, `duration` and `string` values from environment variables, with optional fallback values for unset variables.

```go
import "go.deanishe.net/env"

// Get value for key or return empty string
s := env.Get("SHELL")

// Get value for key or return a specified default
s := env.Get("DOES_NOT_EXIST", "fallback value")

// Get an int (or 0 if SOME_NUMBER is unset or empty)
i := env.GetInt("SOME_NUMBER")

// Int with a fallback
i := env.GetInt("SOME_UNSET_NUMBER", 10)
```


<a id="binding"></a>
### Binding ###

You can also populate a struct directly from the environment by appropriately tagging it and calling `env.Bind()`:

```go
// Simple configuration struct
type config struct {
    HostName string `env:"HOSTNAME"` // default would be HOST_NAME
    Port     int    // leave as default (PORT)
    SSL      bool   `env:"USE_SSL"` // default would be SSL
    Online   bool   `env:"-"`       // ignore this field
}

// Set some values in the environment for test purposes
os.Setenv("HOSTNAME", "api.example.com")
os.Setenv("PORT", "443")
os.Setenv("USE_SSL", "1")
os.Setenv("ONLINE", "1") // will be ignored

// Create a config and bind it to the environment
c := &config{}
if err := Bind(c); err != nil {
    // handle error...
}

// config struct now populated from the environment
fmt.Println(c.HostName)
fmt.Printf("%d\n", c.Port)
fmt.Printf("%v\n", c.SSL)
fmt.Printf("%v\n", c.Online)
// Output:
// api.example.com
// 443
// true
// false
```


<a id="customisation"></a>
### Customisation ###

Variables are retrieved via implementors of the `env.Env` interface (which `env.Bind()` accepts as a second, optional parameter):

```go
type Env interface {
	// Lookup retrieves the value of the variable named by key.
	//
	// It follows the same semantics as os.LookupEnv(). If a variable
	// is unset, the boolean will be false. If a variable is set, the
	// boolean will be true, but the variable may still be an empty
	// string.
	Lookup(key string) (string, bool)
}
```

So you can pass a custom `Env` implementation to `Bind()` in order to populate structs from a source other than environment variables.

See [_examples/docopt][docopt] to see how to implement a custom `Env` that populates a struct from `docopt` command-line options.


<a id="dumping"></a>
### Dumping ###

Dump a struct to a map[string]string by passing it to Dump():

```go
type options struct {
    Hostname string
    Port int
}

o := options{
    Hostname: "www.example.com",
    Port: 22,
}

vars, err := Dump(o)
if err != nil {
     // handler err
}

fmt.Println(vars["HOSTNAME"]) // -> www.example.com
fmt.Println(vars["PORT"])     // -> 22
```


<a id="installation"></a>
Installation
------------

```bash
go get go.deanishe.net/env
```


<a id="documentation"></a>
Documentation
-------------

Read the documentation on [GoDoc][godoc-link].


<a id="licence"></a>
Licence
-------

This library is released under the [MIT Licence][mit].


[mit]: ./LICENCE.txt
[docopt]: _examples/docopt/docopt_example.go

[godoc-icon]: https://godoc.org/go.deanishe.net/env?status.svg
[godoc-link]: https://godoc.org/go.deanishe.net/env
[goreport-link]: https://goreportcard.com/report/go.deanishe.net/env
[goreport-icon]: https://goreportcard.com/badge/go.deanishe.net/env
[coverage-icon]: https://img.shields.io/codacy/coverage/a0ebe54382ad43bf8604b6d6aac02400?color=brightgreen
[codacy-link]: https://www.codacy.com/app/deanishe/go-env
[azure-status-icon]: https://img.shields.io/azure-devops/build/deanishe/3b09feef-08fa-42bc-830e-57ce1de63779/2
[azure-link]: https://dev.azure.com/deanishe/go-env/_build
[licence-icon]: https://img.shields.io/github/license/deanishe/go-env
[licence-link]: https://github.com/deanishe/go-env/blob/master/LICENCE.txt
[github-status-icon]: https://img.shields.io/github/workflow/status/deanishe/go-env/Test
