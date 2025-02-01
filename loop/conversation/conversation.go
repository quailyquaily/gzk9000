// Package ephemeralconversations provides an in-memory store for ephemeral
// conversations. Any conversation that has been idle (no new calls/messages)
// for 30 minutes is automatically removed from memory.
package conversation

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/lyricat/goutils/ai"
)

// conversationTimeout defines the duration after which a conversation is
// considered stale and should be removed.
const conversationTimeout = 30 * time.Minute

// Conversation holds the details of a conversation between two users.
type Conversation struct {
	UserA      uint64
	UserB      uint64
	Messages   []*ai.GeneralChatCompletionMessage
	lastActive time.Time
}

// Manager is responsible for creating, retrieving, and cleaning up
// conversations in memory.
type Manager struct {
	mu            sync.Mutex
	conversations map[string]*Conversation
}

// NewManager initializes and returns a new Manager instance.
// It starts a background goroutine to periodically clean up stale conversations.
func NewManager() *Manager {
	m := &Manager{
		conversations: make(map[string]*Conversation),
	}
	go m.cleanupRoutine()
	return m
}

// GetConversationsByUserIDs retrieves the conversation for the specified user IDs.
// If the conversation does not exist, an error is returned.
func (m *Manager) GetConversationsByUserIDs(ctx context.Context, userIDA, userIDB uint64) (*Conversation, error) {
	key := conversationKey(userIDA, userIDB)

	m.mu.Lock()
	defer m.mu.Unlock()

	conv, exists := m.conversations[key]
	if !exists {
		return nil, errors.New("conversation not found")
	}

	// Optional: refresh lastActive on access
	conv.lastActive = time.Now()

	return conv, nil
}

// CreateConversation creates a new conversation for the specified user IDs
// if one does not already exist. If it already exists, the existing
// conversation is returned (with lastActive refreshed).
func (m *Manager) CreateConversation(ctx context.Context, userIDA, userIDB uint64) (*Conversation, error) {
	key := conversationKey(userIDA, userIDB)

	m.mu.Lock()
	defer m.mu.Unlock()

	// If conversation already exists, just refresh its activity timestamp
	if conv, exists := m.conversations[key]; exists {
		conv.lastActive = time.Now()
		return conv, nil
	}

	// Otherwise, create a new conversation
	conv := &Conversation{
		UserA:      userIDA,
		UserB:      userIDB,
		Messages:   []*ai.GeneralChatCompletionMessage{},
		lastActive: time.Now(),
	}

	m.conversations[key] = conv
	return conv, nil
}

// cleanupRoutine runs periodically (once per minute) to remove any
// conversation that has been inactive longer than conversationTimeout.
func (m *Manager) cleanupRoutine() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		now := time.Now()

		m.mu.Lock()
		for key, conv := range m.conversations {
			if now.Sub(conv.lastActive) > conversationTimeout {
				delete(m.conversations, key)
			}
		}
		m.mu.Unlock()
	}
}

// conversationKey generates a unique key for two user IDs so that
// a conversation for (2,1) is the same as (1,2).
func conversationKey(userA, userB uint64) string {
	if userA > userB {
		userA, userB = userB, userA
	}
	return fmt.Sprintf("%d_%d", userA, userB)
}
