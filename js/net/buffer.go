package net

import (
	"sync"
)

type message []byte

// messageBuffer is an unbounded channel.
type messageBuffer struct {
	c       chan message
	mu      sync.Mutex // Protects following.
	backlog []message
}

func newMessageBuffer() *messageBuffer {
	return &messageBuffer{
		c: make(chan message, 1),
	}
}

// store the message in ch, if empty, otherwise in queue.
func (b *messageBuffer) store(m message) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if len(b.backlog) == 0 {
		select {
		case b.c <- m:
			return
		default:
		}
	}
	b.backlog = append(b.backlog, m)
}

// load moves a message from the queue into ch.
func (b *messageBuffer) load() {
	b.mu.Lock()
	defer b.mu.Unlock()
	if len(b.backlog) > 0 {
		select {
		case b.c <- b.backlog[0]:
			b.backlog[0] = nil // Remove reference from underlying array.
			b.backlog = b.backlog[1:]
		default:
		}
	}
}

// clear removes all messages from buffer.
func (b *messageBuffer) clear() {
	b.mu.Lock()
	b.backlog = nil
	b.mu.Unlock()

	select {
	case <-b.c:
	default:
	}
}

func (b *messageBuffer) get() <-chan message {
	return b.c
}
