package idgen

import (
	"fmt"
	"sync"
	"time"
)

type Snowflake struct {
	mu           sync.Mutex
	lastEpoch    int64
	workerID     int64
	sequence     int64
}

func NewSnowflake(workerID int64) *Snowflake {
	return &Snowflake{workerID: workerID}
}

func (s *Snowflake) NextID() string {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now().UnixMilli()
	if now == s.lastEpoch {
		s.sequence++
	} else {
		s.sequence = 0
	}
	s.lastEpoch = now

	// Formato simple: Tiempo + Worker + Secuencia
	return fmt.Sprintf("%d-%d-%d", now, s.workerID, s.sequence)
}