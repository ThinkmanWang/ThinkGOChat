package worldservice

import (
	"fmt"
	"github.com/lonng/nano"
	"github.com/lonng/nano/component"
	"github.com/lonng/nano/examples/cluster/protocol"
	"github.com/lonng/nano/session"
	"github.com/pingcap/errors"
	"log"
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

func (this *WorldService) JoinRoom(s *session.Session, msg *protocol.JoinRoomRequest) error {
	log.Println("JoinRoom")
	if err := s.Bind(msg.MasterUid); err != nil {
		return errors.Trace(err)
	}

	broadcast := &protocol.NewUserBroadcast{
		Content: fmt.Sprintf("User user join: %v", msg.Nickname),
	}
	if err := this.group.Broadcast("onNewUser", broadcast); err != nil {
		return errors.Trace(err)
	}
	return this.group.Add(s)
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
		log.Println("Remove user from group failed", s.UID(), err)
		return
	}
	log.Println("User session disconnected", s.UID())
}
