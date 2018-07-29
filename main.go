package config12

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// FromEnvironment takes a struct and returns a new struct of the
// same type, using the initial struct's values as the defaults.
// Every struct property with a `c12` tag is then updated with any
// matching environment variable, provided it is a supported type.
// The result can be cast directly to the original struct type.
//
// EXAMPLE:
//
// // config holds settings from the environment
// type config struct {
//   Port             int    `c12:"PORT"`
//   ConnectionString string `c12:"CONNECTION_STRING"`
//   LogRequests      bool   `c12:"LOG_REQUESTS"`
//   SiteName         string
// }
// defaults := config {
//   Port:        3000,
//   LogRequests: false,
// }
// settings := config12.FromEnvironment(defaults).(config)
func FromEnvironment(c interface{}) (interface{}, error) {
	return apply(os.LookupEnv, c)
}

func apply(f func(key string) (string, bool), c interface{}) (interface{}, error) {
	// Basic sanity check
	defaults := reflect.ValueOf(c)
	kind := defaults.Kind()
	if kind != reflect.Struct {
		// Developer error - impossible to recover from
		return nil, fmt.Errorf("(FromEnvironment) expected struct, got %s", kind)
	}

	// Handle every field in turn
	structType := reflect.TypeOf(c)
	fieldCount := structType.NumField()
	result := reflect.New(structType)
	for fieldIdx := 0; fieldIdx < fieldCount; fieldIdx++ {

		// Skip unexported fields (privates)
		if result.Elem().Field(fieldIdx).CanSet() {

			// Apply the passed-in default value
			fieldDefault := defaults.Field(fieldIdx)
			result.Elem().Field(fieldIdx).Set(fieldDefault)

			// Check for an environment variable tag
			field := structType.Field(fieldIdx)
			if envName, ok := field.Tag.Lookup("c12"); ok {

				// Has one, so see if the variable is set
				envValue, found := f(envName)
				if found && len(strings.TrimSpace(envValue)) > 0 {

					// If so, update this field (by supported type)
					switch field.Type.String() {
					case "string":
						result.Elem().Field(fieldIdx).SetString(envValue)
					case "int":
						intEnvValue, err := strconv.Atoi(envValue)
						if err != nil {
							return nil, fmt.Errorf("expected int for %s, got %s", envName, envValue)
						}
						result.Elem().Field(fieldIdx).SetInt(int64(intEnvValue))
					case "bool":
						isTrue := strings.ToLower(envValue) == "true"
						result.Elem().Field(fieldIdx).SetBool(isTrue)
					}
				}
			}
		}
	}

	// Result can be cast directly to the incoming type by the caller
	return result.Elem().Interface(), nil
}
