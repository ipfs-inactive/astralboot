// logging creator
package main

import (
	"github.com/op/go-logging"
	"os"
)

var logger = logging.MustGetLogger("example")

var format = logging.MustStringFormatter(
	"%{color}%{time:0102 15:04:05.000} %{shortfunc:15s} > %{level:.7s} %{id:03x}%{color:reset} %{message}",
)

//LogSetup : set up the logging for information output
func LogSetup(level int) {

	backend1 := logging.NewLogBackend(os.Stderr, "", 0)
	backend1Formatter := logging.NewBackendFormatter(backend1, format)
	backend1Leveled := logging.AddModuleLevel(backend1Formatter)
	switch level {
	case 0:
		backend1Leveled.SetLevel(logging.NOTICE, "example")
	case 1:
		backend1Leveled.SetLevel(logging.INFO, "example")
	case 2:
		backend1Leveled.SetLevel(logging.DEBUG, "example")
	}
	logging.SetBackend(backend1Leveled)
}
