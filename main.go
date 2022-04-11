package main

import (
	"ThinkGOChat/thinkutils/logger"
	"runtime"
)

var (
	log *logger.LocalLogger = logger.DefaultLogger()
)

func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())
    log.Info("Hello World")
	
}
