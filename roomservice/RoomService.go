package roomservice

import (
	"ThinkGOChat/myprotocol"
	"ThinkGOChat/thinkutils"
	"github.com/lonng/nano"
	"github.com/lonng/nano/component"
	"github.com/lonng/nano/session"
)

type RoomService struct {
	component.Base
	nextRoomId int64

	Rooms map[int64]*Room
	//Rooms []*Room
}

type Room struct {
	nano.Group
	Id int64
	Name string
	Owner *session.Session
	CreateTime int64
}

func newRoomService() *RoomService {
	return &RoomService{
		Rooms: map[int64]*Room{},
	}
}

func (this *RoomService) CreateRoom(s *session.Session, msg *myprotocol.CreateRoomReq) error {
	log.Info("create room %s", msg.Name)

	this.nextRoomId++
	rid := this.nextRoomId

	pNewRoom := &Room{
		Id: rid,
		Name: msg.Name,
		Owner: s,
		CreateTime: thinkutils.DateTime.TimestampMs(),
		Group: *nano.NewGroup(msg.Name),
	}
	pNewRoom.Add(s)

	this.Rooms[rid] = pNewRoom

	pRoonInfo := myprotocol.NewRoomInfo()
	pRoonInfo.Id = rid
	pRoonInfo.Name = msg.Name
	pRoonInfo.Members = append(pRoonInfo.Members, &myprotocol.User{
		Id: s.ID(),
	})

	return s.Response(thinkutils.AjaxResultSuccessWithData(pRoonInfo))
}

func (this *RoomService) JoinRoom(s *session.Session, msg *myprotocol.JoinRoomReq) error {
	err := this.Rooms[msg.RoomId].Add(s)
	if err != nil {
		_ = s.Response(thinkutils.AjaxResultError())
		return nil
	}

	pRoonInfo := myprotocol.NewRoomInfo()
	pRoonInfo.Id = msg.RoomId
	pRoonInfo.Name = this.Rooms[msg.RoomId].Name
	lstMember := this.Rooms[msg.RoomId].Members()
	for i := 0; i < this.Rooms[msg.RoomId].Count(); i++ {
		pRoonInfo.Members = append(pRoonInfo.Members, &myprotocol.User{
			Id: lstMember[i],
		})
	}

	return s.Response(thinkutils.AjaxResultSuccessWithData(pRoonInfo))
}

func (this *RoomService) userDisconnected(s *session.Session) {
	//if err := this.group.Leave(s); err != nil {
	//	log.Println("Remove user from group failed", s.UID(), err)
	//	return
	//}
	//log.Println("User session disconnected", s.UID())
}