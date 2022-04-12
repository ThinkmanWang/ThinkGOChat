package usercenter

import (
	"ThinkGOChat/myprotocol"
	"github.com/lonng/nano/component"
	"github.com/lonng/nano/session"
)

type User struct {
	session  *session.Session `json:"-"`
	Name string `json:"name"`
	OpenId string `json:"openId"`
}

type UserService struct {
	component.Base
	nextUid int64
	users   map[string]*User
}

func newUserService() *UserService {
	return &UserService{
		users: map[string]*User{},
	}
}

type ExistsMembersResponse struct {
	Members []*User `json:"members"`
}

func (this *UserService) OnConnected(s *session.Session, msg *myprotocol.NewUserRequest) error {
	return nil
}

func (this *UserService) OnDisconnected(s *session.Session) {

}

//func (this *UserService) NewUser(s *session.Session, msg *myprotocol.NewUserRequest) error {
//	log.Info("%d", s.ID())
//	log.Info(thinkutils.JSONUtils.ToJson(msg))
//
//	this.nextUid++
//	uid := this.nextUid
//	if err := s.Bind(uid); err != nil {
//		return errors.Trace(err)
//	}
//
//	var members []*User
//	for _, u := range this.users {
//		members = append(members, u)
//	}
//	err := s.Push("onMembers", &ExistsMembersResponse{Members: members})
//	if err != nil {
//		return errors.Trace(err)
//	}
//
//	user := &User{
//		session:  s,
//		Name: msg.Nickname,
//		GateId:   msg.GateUid,
//		Id: uid,
//	}
//	this.users[uid] = user
//
//	chat := &myprotocol.JoinWorldRequest{
//		Nickname:  msg.Nickname,
//		GateUid:   msg.GateUid,
//		MasterUid: uid,
//	}
//	return s.RPC("WorldService.JoinRoom", chat)
//}

//func (this *UserService) userDisconnected(s *session.Session) {
//	uid := s.UID()
//	delete(this.users, uid)
//	log.Info("User session disconnected %d", s.UID())
//}