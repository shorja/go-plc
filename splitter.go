package plc

import (
	"fmt"
	"reflect"
)

// SplitReader splits reads of structs and arrays into separate reads of their components.
type SplitReader struct {
	Reader
}

var _ = Reader(SplitReader{}) // Compiler makes sure this type is a Reader

// NewSplitReader returns a SplitReader.
func NewSplitReader(rd Reader) SplitReader {
	return SplitReader{rd}
}

func (r SplitReader) ReadTag(name string, value interface{}) error {
	v := reflect.ValueOf(value)
	if v.Kind() != reflect.Ptr {
		return fmt.Errorf("ReadTag expects a pointer type but got %v", v.Kind())
	}

	err := error(nil)
	switch v.Elem().Kind() {
	case reflect.Struct:
		str := v.Elem()
		for i := 0; i < str.NumField(); i++ {
			// Generate the name of the struct's field and recurse
			// TODO use .Tag when it exists
			fieldName := name + "." + str.Type().Field(i).Name
			if !str.Field(i).CanAddr() {
				err = fmt.Errorf("Cannot address %s", fieldName)
				break
			}
			fieldPointer := str.Field(i).Addr().Interface()

			err = r.ReadTag(fieldName, fieldPointer)
			if err != nil {
				break
			}
		}
	default:
		// Just try with the underlying type
		err = r.Reader.ReadTag(name, value)
	}

	return err
}
