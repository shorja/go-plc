package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/stellentus/go-plc/example"
)

var (
	plcAddr  = flag.String("plc-address", "192.168.1.176", "Hostname or IP address of the PLC")
	httpAddr = flag.String("http", ":8784", "Port for http server to listen to")
)

var knownTags = map[string]interface{}{
	"DUMMY_AQUA_DATA_0": uint16(0),
	"DUMMY_AQUA_DATA_1": uint16(0),
	"DUMMY_AQUA_DATA_2": uint16(0),
	"DUMMY_AQUA_DATA_3": uint16(0),
	"DUMMY_AQUA_DATA_4": uint16(0),
	"DUMMY_AQUA_DATA_5": uint16(0),
	"DUMMY_AQUA_DATA_6": uint16(0),
	"DUMMY_AQUA_DATA_7": uint16(0),
	"DUMMY_AQUA_DATA_8": uint16(0),
	"DUMMY_AQUA_DATA_9": uint16(0),
}

func main() {
	flag.Parse()

	device, err := example.NewDevice(example.Config{
		PrintReadDebug:   true,
		PrintWriteDebug:  true,
		DebugFunc:        fmt.Printf,
		DeviceConnection: map[string]string{"gateway": *plcAddr},
	})
	panicIfError(err, "Could not create test PLC!")
	defer func() {
		err := device.Close()
		if err != nil {
			panic("ERROR: Close was unsuccessful:" + err.Error())
		}
	}()

	http.Handle("/tags/raw", RawTagsHandler{device, knownTags})
	fmt.Printf("Making PLC '%s' available at '%s'\n", *plcAddr, *httpAddr)
	err = http.ListenAndServe(*httpAddr, nil)
	panicIfError(err, "Could not start http server!")
}

func panicIfError(err error, reason string) {
	if err != nil {
		panic("ERROR " + err.Error() + ": " + reason)
	}
}
