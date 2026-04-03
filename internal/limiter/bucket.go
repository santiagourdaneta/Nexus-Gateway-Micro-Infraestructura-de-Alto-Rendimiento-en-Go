package limiter

import (
	"sync"
	"time"
)

type TokenBucket struct {
	capacity   int
	tokens     int
	refillRate time.Duration
	mu         sync.Mutex
}

func NewTokenBucket(capacity int, refillRate time.Duration) *TokenBucket {
	tb := &TokenBucket{
		capacity:   capacity,
		tokens:     capacity,
		refillRate: refillRate,
	}
	// Llenado automático en segundo plano
	go func() {
		for range time.Tick(tb.refillRate) {
			tb.mu.Lock()
			if tb.tokens < tb.capacity {
				tb.tokens++
			}
			tb.mu.Unlock()
		}
	}()
	return tb
}

func (tb *TokenBucket) Allow() bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()
	if tb.tokens > 0 {
		tb.tokens--
		return true
	}
	return false
}