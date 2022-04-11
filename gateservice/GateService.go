package gateservice

import (
	"ThinkGOChat/thinkutils"
	"github.com/lonng/nano/component"
	"github.com/lonng/nano/examples/cluster/protocol"
	"github.com/lonng/nano/session"
	"github.com/pingcap/errors"
)

type GateService struct {
	component.Base
	nextGateUid int64
}

func newGateService() *GateService {
	return &GateService{}
}

type (
	LoginRequest struct {
		Nickname string `json:"nickname"`
	}
	LoginResponse struct {
		Code int `json:"code"`
	}
)

func (this *GateService) Login(s *session.Session, msg *LoginRequest) error {
	log.Info(thinkutils.JSONUtils.ToJson(msg))

	this.nextGateUid++
	uid := this.nextGateUid
	request := &protocol.NewUserRequest{
		Nickname: msg.Nickname,
		GateUid:  uid,
	}
	if err := s.RPC("UserService.NewUser", request); err != nil {
		return errors.Trace(err)
	}
	return s.Response(&LoginResponse{})
}

func (this *GateService) BindChatServer(s *session.Session, msg []byte) error {
	return errors.Errorf("not implement")
}