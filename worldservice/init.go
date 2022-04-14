package worldservice

import (
	"github.com/ThinkmanWang/GOThinkUtils/thinkutils/logger"
	"github.com/lonng/nano/component"
	"github.com/lonng/nano/session"
)

var (
	// All services in master server
	Services = &component.Components{}

	worldService = newWorldService()
	log *logger.LocalLogger = logger.DefaultLogger()
)

func init() {
	Services.Register(worldService)
}

func OnSessionClosed(s *session.Session) {
	worldService.OnDisconnected(s)
}
