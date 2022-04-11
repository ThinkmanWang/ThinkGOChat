package worldservice

import (
	"github.com/lonng/nano/component"
	"github.com/lonng/nano/session"
)

var (
	// All services in master server
	Services = &component.Components{}

	worldService = newWorldService()
)

func init() {
	Services.Register(worldService)
}

func OnSessionClosed(s *session.Session) {
	worldService.userDisconnected(s)
}
