/* Package goenvconfig provides immutability for configuration automatically loaded from environment variables. */
package goenvconfig

import (
	"errors"
	"os"
	"reflect"
	"strconv"
	"unsafe"
)

const (
	envKey        = "env"
	envDefaultKey = "default"
)

/* GoEnvParser represents an object capable of parsing environment variables into a struct, given specific tags. */
type GoEnvParser interface {
	/* Parse accepts a struct pointer and populates private properties according to "env" and "default" tag keys. */
	Parse(object interface{}) error
}

type goEnvParser struct{}

/* NewGoEnvParser returns a new GoEnvParser. */
func NewGoEnvParser() GoEnvParser {
	return &goEnvParser{}
}

func (*goEnvParser) Parse(object interface{}) error {
	if reflect.TypeOf(object).Kind() != reflect.Ptr {
		return errors.New("objects passed to env.Parse() must be of kind pointer")
	}

	addressableCopy := createAddressableCopy(object)

	for i := 0; i < addressableCopy.NumField(); i++ {
		fieldRef := addressableCopy.Field(i)
		fieldRef = reflect.NewAt(fieldRef.Type(), unsafe.Pointer(fieldRef.UnsafeAddr())).Elem()

		newValue := getValueForTag(reflect.TypeOf(object).Elem(), i)

		switch fieldRef.Type().Kind() {
		case reflect.Int:
			if newInt, err := strconv.ParseInt(newValue, 10, 32); err == nil {
				fieldRef.SetInt(newInt)
			}
		case reflect.String:
			fieldRef.SetString(newValue)
		}
	}

	object = addressableCopy.Interface()

	return nil
}

func createAddressableCopy(object interface{}) reflect.Value {
	originalValue := reflect.ValueOf(object)
	objectCopy := reflect.New(originalValue.Type()).Elem()
	objectCopy.Set(originalValue)

	if originalValue.Type().Kind() == reflect.Ptr {
		objectCopy = objectCopy.Elem()
	}

	return objectCopy
}

func getValueForTag(addressableCopy reflect.Type, fieldNum int) string {
	if envKey, ok := addressableCopy.Field(fieldNum).Tag.Lookup(envKey); ok {
		if envVar, exists := os.LookupEnv(envKey); exists {
			return envVar
		}
	}

	if envDefaultValue, ok := addressableCopy.Field(fieldNum).Tag.Lookup(envDefaultKey); ok {
		return envDefaultValue
	}

	return ""
}
