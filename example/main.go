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
	foo(logger)
	foo(logger)
	foo(logger)
}

func foo(l fmtlogger.FmtLogger) {
	logger := fmtlogger.NewContext(l).Follow("method","foo")
	logger.Log("foo", "bar")
}
