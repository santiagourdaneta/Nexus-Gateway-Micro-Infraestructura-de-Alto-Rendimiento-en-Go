package scanner

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func ScanLocalPorts(host string, startPort, endPort int) []int {
	var openPorts []int
	var mu sync.Mutex
	var wg sync.WaitGroup

	for port := startPort; port <= endPort; port++ {
		wg.Add(1)
		go func(p int) {
			defer wg.Done()
			address := fmt.Sprintf("%s:%d", host, p)
			conn, err := net.DialTimeout("tcp", address, 500*time.Millisecond)
			if err == nil {
				conn.Close()
				mu.Lock()
				openPorts = append(openPorts, p)
				mu.Unlock()
			}
		}(port)
	}
	wg.Wait()
	return openPorts
}