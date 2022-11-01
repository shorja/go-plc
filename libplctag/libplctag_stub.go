//go:build stub
// +build stub

package libplctag

import "fmt"

/*
#cgo LDFLAGS: -lplctagstub
#include <libplctag.h>
*/
import "C"

func init() {

	fmt.Println("Hello from go-plc stub!")

}

const StubActive = true
