package wunderground

import (
	"flag"
	"log"
	"os"
)

var logger = log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile)
var verbosity = flag.Int("verbosity", 0, "Log more info at higher verbosity levels")

func logDebug(args ...interface{}) {
	if *verbosity < 2 {
		return
	}

	p := logger.Prefix()
	defer logger.SetPrefix(p)
	logger.SetPrefix("debug: ")
	logger.Print(args...)
}

func logDebugf(format string, args ...interface{}) {
	if *verbosity < 2 {
		return
	}

	p := logger.Prefix()
	defer logger.SetPrefix(p)
	logger.SetPrefix("debug: ")
	logger.Printf(format, args...)
}
