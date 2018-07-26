# config-12

Simple but powerful automatic environment variable support for easy configuration.

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
  defaults := config {
    Port:        3000,
    LogRequests: false,
  }

  // Pass it into config-12 and you're done
  settings := config12.FromEnvironment(defaults).(config)
  fmt.Println("Running on port " + settings.Port)
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
cd config-13
go test
```
