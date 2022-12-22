package session

/*
* @author: Heng ChenChi
* @date: 2022/12/22 0022 10:21
* @version: 1.0
* @description:
**/

type SessionInterface interface {
	GetSessionID() uint32
	Send(message []byte)
	Close() error
}

type SessionMgrInterface interface {
	Add(session SessionInterface)
	Remove(session SessionInterface) bool
	Count() int

	Push(sessionID uint32, message []byte)
	PushAll(message []byte)

	KickOut(sessionID uint32) bool
}
