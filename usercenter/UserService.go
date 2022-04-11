package usercenter

import (
	"ThinkGOChat/myprotocol"
	"ThinkGOChat/thinkutils"
	"github.com/lonng/nano/component"
	"github.com/lonng/nano/session"
	"github.com/pingcap/errors"
)

type User struct {
	session  *session.Session `json:"-"`
	name string `json:"name"`
	gateId   int64 `json:"-"`
	masterId int64 `json:"-"`
}

type UserService struct {
	component.Base
	nextUid int64
	users   map[int64]*User
}

func newUserService() *UserService {
	return &UserService{
		users: map[int64]*User{},
	}
}

type ExistsMembersResponse struct {
	Members []*User `json:"members"`
}

func (this *UserService) NewUser(s *session.Session, msg *myprotocol.NewUserRequest) error {
	log.Info(thinkutils.JSONUtils.ToJson(msg))

	this.nextUid++
	uid := this.nextUid
	if err := s.Bind(uid); err != nil {
		return errors.Trace(err)
	}

	var members []*User
	for _, u := range this.users {
		members = append(members, u)
	}
	err := s.Push("onMembers", &ExistsMembersResponse{Members: members})
	if err != nil {
		return errors.Trace(err)
	}

	user := &User{
		session:  s,
		name: msg.Nickname,
		gateId:   msg.GateUid,
		masterId: uid,
	}
	this.users[uid] = user

	chat := &myprotocol.JoinWorldRequest{
		Nickname:  msg.Nickname,
		GateUid:   msg.GateUid,
		MasterUid: uid,
	}
	return s.RPC("WorldService.JoinRoom", chat)
}

func (this *UserService) userDisconnected(s *session.Session) {
	uid := s.UID()
	delete(this.users, uid)
	log.Info("User session disconnected %d", s.UID())
}