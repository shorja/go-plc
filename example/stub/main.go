//go:build stub
// +build stub

package main

import (
	"fmt"

	"github.com/stellentus/go-plc/libplctag"
)

func main() {
	libplctag.SetDebug(libplctag.DebugInfo)
	libplctag.DoesIsStub()
	fmt.Println("blah")

	device, err := libplctag.NewDevice("who.cares@youtube.gov")
	panicIfError(err, "Could not create test PLC!")
	defer func() {
		err := device.Close()
		if err != nil {
			fmt.Println("Close was unsuccessful:", err.Error())
		}
	}()

	var testBoolName string = "testBool"
	var testBool bool = false

	read(device, testBoolName, &testBool)
	print(testBoolName, testBool)

	testBool = !testBool
	write(device, testBoolName, testBool)

	read(device, testBoolName, &testBool)
	print(testBoolName, testBool)

	var testInt32Name string = "testInt32"
	var testInt32 int32 = 0
	read(device, testInt32Name, &testInt32)
	print(testInt32Name, testInt32)

	testInt32 = int32(413)
	write(device, testInt32Name, testInt32)

	read(device, testInt32Name, &testInt32)
	print(testInt32Name, testInt32)

	device.GetAllTags()
	device.GetAllPrograms()
	device.Close()
}

func read(device *libplctag.Device, tagName string, value interface{}) {
	err := device.ReadTag(tagName, value)
	panicIfError(err, "Unable to read the data")
}

func write(device *libplctag.Device, tagName string, value interface{}) {
	err := device.WriteTag(tagName, value)
	panicIfError(err, "Unable to write the data")
}

func print(tagName string, value interface{}) {
	fmt.Printf("%s is %v\n", tagName, value)
}

func panicIfError(err error, reason string) {
	if err != nil {
		panic("ERROR " + err.Error() + ": " + reason)
	}
}
