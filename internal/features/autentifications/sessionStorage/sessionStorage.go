package sessionstorage

import (
	"fmt"
	"sync"

	common_errors "github.com/Andromeeeda/Subtracker/internal/core/errors/common"
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

func (s *SessionStorage) Get(sessionId string) (string, bool) {
	session, ok := s.sessions[sessionId]
	if !ok {
		return "", false
	}

	return session, true
}

func (s *SessionStorage) DeleteSession(userID string) error {

	for sessionID, uid := range s.sessions {
		if uid == userID {
			delete(s.sessions, sessionID)
			return nil
		}
	}

	return fmt.Errorf("session not found! %w", common_errors.ErrNotFound)
}
