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
	AllUser *nano.Group
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
		AllUser: nano.NewGroup("all-users"),
	}
}

func (this *RoomService) OnConnected(s *session.Session, msg *myprotocol.NewUserRequest) error {
	s.Set("openId", msg.OpenId)
	if err := this.AllUser.Add(s); err != nil {
	}

	log.Info("%p %s OnConnected", s, msg.OpenId)
	return nil
}

func (this *RoomService) OnDisconnected(s *session.Session) {

}

func (this *RoomService) CreateRoom(s *session.Session, msg *myprotocol.CreateRoomReq) error {
	log.Info("%s call CreateRoom", s.String("openId"))

	this.nextRoomId++
	rid := this.nextRoomId

	pNewRoom := &Room{
		Id: rid,
		Name: msg.Name,
		Owner: s,
		CreateTime: thinkutils.DateTime.TimestampMs(),
		Group: *nano.NewGroup(msg.Name),
	}
	if err := pNewRoom.Add(s); err != nil {
		return err
	}

	this.Rooms[rid] = pNewRoom

	pRoonInfo := myprotocol.NewRoomInfo()
	pRoonInfo.Id = rid
	pRoonInfo.OwnerId = s.String("openId")
	pRoonInfo.Name = msg.Name
	pRoonInfo.Members = append(pRoonInfo.Members, &myprotocol.User{
		OpenId: s.String("openId"),
	})

	_ = s.Response(thinkutils.AjaxResultSuccessWithData(pRoonInfo))

	_ = this.AllUser.Broadcast("onCreateRoom", pRoonInfo)

	return nil
}

func (this *RoomService) JoinRoom(s *session.Session, msg *myprotocol.JoinRoomReq) error {
	log.Info("%s join room %d", s.String("openId"), msg.RoomId)
	err := this.Rooms[msg.RoomId].Add(s)
	if err != nil {
		_ = s.Response(thinkutils.AjaxResultError())
		return nil
	}

	pRoonInfo := this.createRoomInfo(this.Rooms[msg.RoomId])

	return s.Response(thinkutils.AjaxResultSuccessWithData(pRoonInfo))
}

func (this *RoomService) createRoomInfo(pRoom *Room) *myprotocol.RoomInfo {
	pRoonInfo := myprotocol.NewRoomInfo()
	pRoonInfo.Id = pRoom.Id
	pRoonInfo.Name = pRoom.Name
	pRoonInfo.OwnerId = pRoom.Owner.String("openId")

	lstMember := pRoom.Members()
	for i := 0; i < pRoom.Count(); i++ {
		if session, err := pRoom.Member(lstMember[i]); err == nil {
			pRoonInfo.Members = append(pRoonInfo.Members, &myprotocol.User{
				OpenId: session.String("openId"),
			})
		}
	}

	return pRoonInfo
}

func (this *RoomService) RoomList(s *session.Session, msg *myprotocol.EmptyReq) error {
	lstRoom := make([]*myprotocol.RoomInfo, 0)

	for _, item := range this.Rooms {
		pRoonInfo := this.createRoomInfo(item)
		lstRoom = append(lstRoom, pRoonInfo)
	}

	return s.Response(thinkutils.AjaxResultSuccessWithData(lstRoom))
}

func (this *RoomService) SendMessage(s *session.Session, msg *myprotocol.RoomMessage) error {
	log.Info("%p %s call SendMessage", s, s.String("openId"))
	msg.OpenId = s.String("openId")
	return this.Rooms[msg.RoomId].Broadcast("onMessage", msg)
}