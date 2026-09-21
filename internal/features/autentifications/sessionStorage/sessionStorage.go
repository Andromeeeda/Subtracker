package sessionstorage

import (
	"sync"

	"github.com/google/uuid"
)

type SessionStorage struct {
	mu       sync.RWMutex
	sessions map[string]string
}

func NewSessionStorage() *SessionStorage {
	return &SessionStorage{
		sessions: make(map[string]string),
	}
}

func (s *SessionStorage) Create(userId string) string {
	s.mu.Lock()
	defer s.mu.Unlock()

	sessionId := uuid.New().String()

	s.sessions[sessionId] = userId
	
	return sessionId
}

