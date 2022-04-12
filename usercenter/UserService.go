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
	Name string `json:"name"`
	OpenId string `json:"openId"`
}

type UserService struct {
	component.Base
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

	log.Info("%d", s.ID())
	log.Info(thinkutils.JSONUtils.ToJson(msg))

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
		Name: msg.Nickname,
		OpenId: msg.OpenId,
	}
	s.Set("openId", msg.OpenId)

	this.users[msg.OpenId] = user

	log.Info("User Count %d", len(this.users))

	return nil
}

func (this *UserService) OnDisconnected(s *session.Session) {
	szOpenId := s.String("openId")
	log.Info("%s disconnected", szOpenId)

	delete(this.users, szOpenId)
}