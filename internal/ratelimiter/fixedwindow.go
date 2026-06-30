package ratelimiter

import (
	"sync"
	"time"
)

type windowState struct {
	count       int
	windowStart time.Time
}

type FixedWindowLimiter struct {
	mu         sync.Mutex
	limit      int
	windowSize time.Duration
	clients    map[string]*windowState
}

func NewFixedWindowLimiter(limit int, windowSize time.Duration) *FixedWindowLimiter {
	return &FixedWindowLimiter{
		limit:      limit,
		windowSize: windowSize,
		clients:    make(map[string]*windowState),
	}
}

func (l *FixedWindowLimiter) Allow(clientID string) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()
	state, exists := l.clients[clientID]

	if !exists || now.Sub(state.windowStart) >= l.windowSize {
		l.clients[clientID] = &windowState{
			count:       1,
			windowStart: now,
		}
		return true
	}

	if state.count < l.limit {
		state.count++
		return true
	}

	return false
}
