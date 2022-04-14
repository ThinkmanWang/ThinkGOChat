package worldservice

import (
	"ThinkGOChat/myprotocol"
	"fmt"
	"github.com/lonng/nano"
	"github.com/lonng/nano/component"
	"github.com/lonng/nano/session"
)

type User struct {
	Session *session.Session
}

type WorldService struct {
	component.Base
	group *nano.Group

	Users map[string]*User
}

func newWorldService() *WorldService {
	return &WorldService{
		group: nano.NewGroup("all-users"),
		Users: map[string]*User{},
	}
}

type SendMessage struct {
	OpenId string `json:"openId"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

func (this *WorldService) SendMessage(s *session.Session, msg *SendMessage) error {
	log.Info("%p %s call SendMessage", s, s.String("openId"))
	msg.OpenId = s.String("openId")
	return this.group.Broadcast(myprotocol.CLIENT_EVENT_ON_CHAT_MESSAGE, msg)
}

func (this *WorldService) OnConnected(s *session.Session, msg *myprotocol.NewUserRequest) error {
	log.Info("%p %s JoinRoom", s, msg.OpenId)
	s.Set("openId", msg.OpenId)
	this.Users[msg.OpenId] = &User{
		Session: s,
	}

	broadcast := &myprotocol.NewUserBroadcast{
		Content: fmt.Sprintf("User user join: %v", msg.Nickname),
	}
	if err := this.group.Broadcast(myprotocol.CLIENT_EVENT_ON_NEW_USER, broadcast); err != nil {
	}

	if err := this.group.Add(s); err != nil {
	}

	log.Info("All User Count: %d", len(this.Users))

	return nil
}

func (this *WorldService) OnDisconnected(s *session.Session) {
	delete(this.Users, s.String("openId"))
	log.Info("All User Count: %d", len(this.Users))
}
