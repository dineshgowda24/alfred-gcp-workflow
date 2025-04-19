// Copyright (c) 2018 Dean Jackson <deanishe@deanishe.net>
// MIT Licence applies http://opensource.org/licenses/MIT

/*
Package env maps environment variables to struct fields and vice versa.

It is heavily based on github.com/caarlos0/env, but has different semantics,
and also allows the dumping of a struct to environment variables, not just
populating a struct from environment variables.


Reading variables

Read environment variables with the Get* functions:

	// String
	s := env.Get("HOME")
	// String with fallback
	s = env.Get("NON_EXISTENT_VAR", "default") // -> default

	// Int
	i := env.GetInt("SOME_COUNT")
	// Int with fallback
	i = env.GetInt("NON_EXISTENT_VAR", 10) // -> 10

	// Duration (e.g. "1m", "2h30m", "5.3s")
	d := env.GetDuration("SOME_TIME")
	// Duration with fallback
	d = env.GetDuration("NON_EXISTENT_VAR", time.Minute * 120) // -> 2h0m


Populating structs

Populate a struct from the environment by passing it to Bind():

	type options struct {
		Hostname string
		Port     int
	}

	o := &options{}
	if err := env.Bind(o); err != nil {
		// handle error...
	}

	fmt.Println(o.Hostname) // -> value of HOSTNAME environment variable
	fmt.Println(o.Port)     // -> value of PORT environment variable

Use tags to specify a variable name or ignore a field:

	type options {
		HostName string `env:"HOSTNAME"` // default would be HOST_NAME
		Port     int
		Online   bool `env:"-"` // ignored
	}


Dumping structs

Dump a struct to a map[string]string by passing it to Dump():

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


Tags

Add `env:"..."` tags to your struct fields to bind them to specific
environment variables or ignore them. `env:"-"` tells Bind() to
ignore the field:

	type options {
		UserName      string
		LastUpdatedAt string `env:"-"` // Not loaded from environment
	}

Add `env:"VARNAME"` to bind a field to the variable VARNAME:

	type options {
		UserName string `env:"USERNAME"`   // default = USER_NAME
		APIKey   string `env:"APP_SECRET"` // default = API_KEY
	}


Customisation

Variables are retrieved via implementors of the Env interface, which
Bind() accepts as a second, optional parameter.

So you can pass a custom Env implementation to Bind() to populate
structs from a source other than environment variables.

See _examples/docopt to see a custom Env implementation used to
populate a struct from docopt command-line options.

You can also customise the map keys used when dumping a struct by passing
VarNameFunc to Dump().


Licence

This library is released under the MIT Licence.

*/
package env
