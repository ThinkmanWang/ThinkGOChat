package myprotocol

import "github.com/lonng/nano/session"

type NewUserBroadcast struct {
	Content string `json:"content"`
}

type NewUserRequest struct {
	Nickname string `json:"nickname"`
	OpenId string `json:"openId"`
}

type BaseService interface {
	OnConnected(s *session.Session, msg *NewUserRequest) error
	OnDisconnected(s *session.Session)
}

type JoinWorldRequest struct {
	Nickname  string `json:"nickname"`
	GateUid   int64  `json:"gateUid"`
	MasterUid int64  `json:"masterUid"`
}

type MasterStats struct {
	Uid int64 `json:"uid"`
}

type CreateRoomReq struct {
	Name string `json:"name"`
}

type User struct {
	OpenId string `json:"openId"`
	Name string `json:"name"`
}

type RoomInfo struct {
	Id int64 `json:"id"`
	OwnerId string `json:"ownerId"`
	Name string `json:"name"`
	Members []*User `json:"members"`
}

type JoinRoomReq struct {
	RoomId int64 `json:"roomId"`
}

type EmptyReq struct {
}

func NewRoomInfo() *RoomInfo {
	pRoom := &RoomInfo{
		Members: []*User{},
	}

	return pRoom
}

type RoomMessage struct {
	OpenId string `json:"openId"`
	RoomId int64 `json:"roomId"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

const (
	CLIENT_EVENT_ON_NEW_USER = "onNewUser"
	CLIENT_EVENT_ON_JOIN_ROOM = "onJoinRoom"
	CLIENT_EVENT_ON_CREATE_ROOM = "onCreateRoom"
	CLIENT_EVENT_ON_ROOM_UPDATE = "onRoomUpdate"
	CLIENT_EVENT_ON_CHAT_MESSAGE = "onChatMessage"
)