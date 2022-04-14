package roomservice

import (
	"github.com/ThinkmanWang/GOThinkUtils/thinkutils/logger"
	"github.com/lonng/nano/component"
	"github.com/lonng/nano/session"
)

var (
	// All services in master server
	Services = &component.Components{}

	roomService = newRoomService()

	log *logger.LocalLogger = logger.DefaultLogger()
)

func init() {
	Services.Register(roomService)
}

func OnSessionClosed(s *session.Session) {
	roomService.OnDisconnected(s)
}