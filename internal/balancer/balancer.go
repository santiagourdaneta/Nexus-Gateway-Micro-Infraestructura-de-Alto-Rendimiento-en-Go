package balancer

import "sync/atomic"

type RoundRobin struct {
	targets []string
	current uint64
}

func NewRoundRobin(targets []string) *RoundRobin {
	return &RoundRobin{targets: targets}
}

func (rr *RoundRobin) Next() string {
	if len(rr.targets) == 0 {
		return ""
	}
	idx := atomic.AddUint64(&rr.current, 1) % uint64(len(rr.targets))
	return rr.targets[idx]
}