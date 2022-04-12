package worldservice

import (
	"ThinkGOChat/myprotocol"
	"ThinkGOChat/thinkutils"
	"fmt"
	"github.com/lonng/nano"
	"github.com/lonng/nano/component"
	"github.com/lonng/nano/session"
)

type WorldService struct {
	component.Base
	group *nano.Group
}

func newWorldService() *WorldService {
	return &WorldService{
		group: nano.NewGroup("all-users"),
	}
}

type SendMessage struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

func (this *WorldService) SendMessage(s *session.Session, msg *SendMessage) error {
	// Send an RPC to master server to stats
	//if err := s.RPC("TopicService.Stats", &protocol.MasterStats{Uid: s.UID()}); err != nil {
	//	return errors.Trace(err)
	//}

	// Sync message to all members in this room
	return this.group.Broadcast("onMessage", msg)
}

func (this *WorldService) userDisconnected(s *session.Session) {
	if err := this.group.Leave(s); err != nil {
		log.Info("Remove user from group failed", s.UID(), err)
		return
	}
	log.Info("User session disconnected", s.UID())
}

func (this *WorldService) NewUser(s *session.Session, msg *myprotocol.NewUserRequest) error {
	log.Info("%d", s.ID())
	log.Info(thinkutils.JSONUtils.ToJson(msg))

	return nil
}

func (this *WorldService) OnConnected(s *session.Session, msg *myprotocol.NewUserRequest) error {
	log.Info("JoinRoom uid: %d", s.ID())
	s.Set("openid", msg.OpenId)

	broadcast := &myprotocol.NewUserBroadcast{
		Content: fmt.Sprintf("User user join: %v", msg.Nickname),
	}
	if err := this.group.Broadcast("onNewUser", broadcast); err != nil {
	}

	if err := this.group.Add(s); err != nil {

	}

	return nil
}

func (this *WorldService) OnDisconnected(s *session.Session) {

}
