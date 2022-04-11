package gateservice

import (
	"ThinkGOChat/thinkutils/logger"
	"github.com/lonng/nano/component"
)

var (
	// All services in master server
	Services = &component.Components{}

	gateService = newGateService()

	log *logger.LocalLogger = logger.DefaultLogger()
)

func init() {
	Services.Register(gateService)
}