package main

import "sync"

type PubSub[T any] struct {
	subscribers []chan T
	mu          sync.RWMutex
	isClosed    bool
}

func NewPubSub[T any]() *PubSub[T] {
	return &PubSub[T]{
		mu: sync.RWMutex{},
	}
}

func (ps *PubSub[T]) Subscribe() <-chan T {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	if ps.isClosed {
		return nil
	}

	sub := make(chan T)
	ps.subscribers = append(ps.subscribers, sub)

	return sub
}

func (ps *PubSub[T]) Publish(value T) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	if ps.isClosed {
		return
	}

	for _, sub := range ps.subscribers {
		sub <- value
	}
}

func (ps *PubSub[T]) Close() {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	if ps.isClosed {
		return
	}

	for _, sub := range ps.subscribers {
		close(sub)
	}

	ps.isClosed = true
}
