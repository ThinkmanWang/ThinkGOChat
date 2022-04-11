package main

import (
	"ThinkGOChat/thinkutils"
	"ThinkGOChat/thinkutils/logger"
)

var (
	log *logger.LocalLogger = logger.DefaultLogger()
)

func main() {
	thinkutils.FileUtils.ReadLine("example.txt", func(nLine uint32, szLine string) {
		log.Info("%d %s", nLine, szLine)
	})

	thinkutils.FileUtils.Copy("example.txt", "test123.txt")
}
