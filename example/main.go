package main

import (
	"github.com/kvdaddy/fmtlogger"
	"os"
)

func main() {
	logger := fmtlogger.NewFmtLogger(os.Stdout)
	logger = fmtlogger.NewContext(logger).Follow("head","h1")
	logger = fmtlogger.NewContext(logger).Follow("head2", "h2")
	logger.Log("middle", "m3")
}
