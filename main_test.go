package config12

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// structure for tests
type configType struct {
	StringWithDefault  string `c12:"MISSING"`
	StringNoDefault    string `c12:"STRING"`
	IntNoDefault       int    `c12:"INT"`
	InvalidInt         int    `c12:"INVALID_INT"`
	BoolTrueNoDefault  bool   `c12:"BOOL_TRUE"`
	BoolFalseNoDefault bool   `c12:"BOOL_FALSE"`
}

// instance for tests
var settings = configType{
	StringWithDefault: "default",
}

func Test_Apply_InputNotStruct_ReturnsError(t *testing.T) {
	_, err := apply(notCalled, "this is not a struct")

	assert.Error(t, err)
}

func Test_Apply_StructWithNoEnvVars_UsesDefaults(t *testing.T) {
	r, err := apply(notFound, settings)

	res := r.(configType)

	assert.NoError(t, err)
	assert.Equal(t, settings.StringWithDefault, res.StringWithDefault)
	assert.Equal(t, "", res.StringNoDefault)
	assert.Equal(t, 0, res.IntNoDefault)
	assert.Equal(t, 0, res.InvalidInt)
	assert.Equal(t, false, res.BoolTrueNoDefault)
	assert.Equal(t, false, res.BoolFalseNoDefault)
}

func Test_Apply_StructWithEnvVars_OverridesDefaults(t *testing.T) {
	r, err := apply(found, settings)
	res := r.(configType)

	assert.NoError(t, err)

	assert.Equal(t, settings.StringWithDefault, res.StringWithDefault)
	assert.Equal(t, "found", res.StringNoDefault)
	assert.Equal(t, 123, res.IntNoDefault)
	assert.Equal(t, 0, res.InvalidInt)
	assert.Equal(t, true, res.BoolTrueNoDefault)
	assert.Equal(t, false, res.BoolFalseNoDefault)
}

// Stub functions to simulate os.LookupEnv

func notCalled(key string) (string, bool) {
	return "", false
}

func notFound(key string) (string, bool) {
	return "", false
}

func found(key string) (string, bool) {
	switch key {
	case "STRING":
		return "found", true
	case "INT":
		return "123", true
	case "INVALID_INT":
		return "not a number", true
	case "BOOL_TRUE":
		return "true", true
	case "BOOL_FALSE":
		return "whatever", true
	}
	return "", false
}
