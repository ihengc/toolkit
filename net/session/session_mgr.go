package session

import "sync"

/*
* @author: Heng ChenChi
* @date: 2022/12/22 0022 10:24
* @version: 1.0
* @description:
**/

var (
	once       sync.Once
	sessionMgr *SessionMgr
)

type SessionMgr struct {
	mu       sync.Mutex
	sessions map[uint32]SessionInterface
}

func NewSessionMgr() *SessionMgr {
	once.Do(func() {
		if sessionMgr == nil {
			sessionMgr = &SessionMgr{
				sessions: make(map[uint32]SessionInterface),
			}
		}
	})
	return sessionMgr
}

func (s *SessionMgr) Add(session SessionInterface) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.sessions[session.GetSessionID()] = session
}

func (s *SessionMgr) Remove(session SessionInterface) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.sessions[session.GetSessionID()]; ok {
		delete(s.sessions, session.GetSessionID())
		return true
	}
	return false
}

func (s *SessionMgr) Count() int {
	s.mu.Lock()
	defer s.mu.Lock()
	return len(s.sessions)
}

func (s *SessionMgr) Push(sessionID uint32, message []byte) {

}

func (s *SessionMgr) PushAll(message []byte) {

}
