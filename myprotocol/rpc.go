package myprotocol

type NewUserRequest struct {
	Nickname string `json:"nickname"`
	GateUid  int64  `json:"gateUid"`
}

type JoinWorldRequest struct {
	Nickname  string `json:"nickname"`
	GateUid   int64  `json:"gateUid"`
	MasterUid int64  `json:"masterUid"`
}

type MasterStats struct {
	Uid int64 `json:"uid"`
}