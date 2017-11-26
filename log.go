package wunderground

import (
	"log"
	"os"
)

var logger = log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile)
