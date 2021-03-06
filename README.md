# config-12

Simple but powerful automatic environment variable support for easy configuration.

* All your **environment variable** configuration entries are made available as a **single Go structure**
* Provide your own environment variable name for each field in your structure
* Supports **strings**, **integers**, and **boolean flags**
* Validates there is *actual content* within string variables
* Validates a *whole number* is provided for integer variables
* Accepts *case-insensitive* 'true' for boolean flags
* Fully tested, extremely simple, and no further dependencies

## Installation

```sh
go install github.com/kcartlidge/config-12
```

## Usage

```golang
package main

// Import the package

import (
  "fmt"

  config12 "github.com/kcartlidge/config-12"
)

// Declare a struct for your settings

type config struct {
  Port             int    `c12:"PORT"`
  ConnectionString string `c12:"CONNECTION_STRING"`
  LogRequests      bool   `c12:"LOG_REQUESTS"`
  SiteName         string
}

func main() {
  // Create an instance with your defaults
  conf := config {
    Port:        3000,
    LogRequests: false,
  }

  // Pass it into config-12 and you're done
  s, err := config12.FromEnvironment(conf)
  if err != nil {
    log.Fatalln(err.Error())
  }
  conf = s.(config)
  fmt.Println("Running on port " + conf.Port)
}
```

## How it works

We start with the defaults, then update each field of the struct that has a `c12` tag with any matching (valid) environment variable of the name given.

_Note that only public fields are considered. Unexported fields are neither looked for, nor given any defaults._

Environment variables are considered valid when:

- a string has actual content (not just whitespace)
- an int is convertable to a whole number
- a bool is 'true' (case-insensitive); all other values are false

Other standard types are not currently supported, but will be arriving in due course. Nested or complex types won't be. There are other packages out there for more comprehensive stuff; _config-12_ is aiming for simplicity and ease of use.

## Running the tests

Navigate into the package folder and run via Go:

```sh
cd config-12
go test
```
