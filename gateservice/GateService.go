package gateservice

import (
	"ThinkGOChat/myprotocol"
	"github.com/ThinkmanWang/GOThinkUtils/thinkutils"
	"github.com/lonng/nano/component"
	"github.com/lonng/nano/session"
	"github.com/pingcap/errors"
)

type GateService struct {
	component.Base
}

func newGateService() *GateService {
	return &GateService{}
}

type (
	LoginRequest struct {
		Nickname string `json:"nickname"`
	}
)

var (
	// All services in master server
	Services = &component.Components{}
	gateService = newGateService()
)

func init() {
	Services.Register(gateService)
}

func OnSessionClosed(s *session.Session) {
	gateService.OnDisconnected(s)
}

func (this *GateService) Login(s *session.Session, msg *LoginRequest) error {
	log.Info(thinkutils.JSONUtils.ToJson(msg))

	szOpenId := thinkutils.UUIDUtils.New()
	request := &myprotocol.NewUserRequest{
		Nickname: msg.Nickname,
		OpenId: szOpenId,
	}

	//if err := s.RPC("UserService.OnConnected", request); err != nil {
	//	return errors.Trace(err)
	//}

	go func() {
		if err := s.RPC("WorldService.OnConnected", request); err != nil {
			return
		}

		if err := s.RPC("RoomService.OnConnected", request); err != nil {
			return
		}
	}()

	return s.Response(thinkutils.AjaxResultSuccess())
}

func (this *GateService) BindChatServer(s *session.Session, msg []byte) error {
	return errors.Errorf("not implement")
}

func (this *GateService) OnConnected(s *session.Session, msg *myprotocol.NewUserRequest) error {
	return nil
}

func (this *GateService) OnDisconnected(s *session.Session) {

}