package helper

import (
	"bytes"
	"fmt"
	"reflect"
	"unsafe"
)

func printContextInternals(ctx interface{}, inner bool) *bytes.Buffer {
	buf := &bytes.Buffer{}

	contextValues := reflect.ValueOf(ctx).Elem()
	contextKeys := reflect.TypeOf(ctx).Elem()

	if !inner {
		fmt.Printf("\nFields for %s.%s\n", contextKeys.PkgPath(), contextKeys.Name())
	}

	if contextKeys.Kind() == reflect.Struct {
		for i := 0; i < contextValues.NumField(); i++ {
			reflectValue := contextValues.Field(i)
			reflectValue = reflect.NewAt(reflectValue.Type(), unsafe.Pointer(reflectValue.UnsafeAddr())).Elem()

			reflectField := contextKeys.Field(i)

			if reflectField.Name == "Context" {
				printContextInternals(reflectValue.Interface(), true)
			} else {
				fmt.Fprintf(buf, "field name: %+v\n", reflectField.Name)
				fmt.Fprintf(buf, "value: %+v\n", reflectValue.Interface())
			}
		}
	} else {
		fmt.Fprintf(buf, "context is empty (int)\n")
	}

	return buf
}
