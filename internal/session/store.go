package session

import (
	"sync"
)

type Session struct {
	ClientID    string
	RedirectURI string
	State       string
	Nonce       string
	Verified    bool
	SubjectDID  string
}

var store = struct {
	sync.RWMutex
	sessions map[string]*Session
}{
	sessions: make(map[string]*Session),
}

func Save(id string, s *Session) {
	store.Lock()
	defer store.Unlock()
	store.sessions[id] = s
}

func Get(id string) (*Session, bool) {
	store.RLock()
	defer store.RUnlock()
	s, found := store.sessions[id]
	return s, found
}

func MarkVerified(id, subjectDID string) {
	store.Lock()
	defer store.Unlock()
	if s, ok := store.sessions[id]; ok {
		s.Verified = true
		s.SubjectDID = subjectDID
	}
}
