package usercenter

import (
	"ThinkGOChat/thinkutils/logger"
	"github.com/lonng/nano/component"
	"github.com/lonng/nano/session"
)

var (
	// All services in master server
	Services = &component.Components{}

	// Topic service
	userService = newUserService()
	// ... other services

	log *logger.LocalLogger = logger.DefaultLogger()
)

func init() {
	Services.Register(userService)
}

func OnSessionClosed(s *session.Session) {
	userService.OnDisconnected(s)
}

