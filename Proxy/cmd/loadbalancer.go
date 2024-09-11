package main

import (
	"sync"
)

type LoadBalancer struct {
	Current int
	Mutex   sync.Mutex
	Servers []*Server
}

func (lb *LoadBalancer) getNextServer(servers []*Server) *Server {
	lb.Mutex.Lock()
	defer lb.Mutex.Unlock()

	if len(servers) == 0 {
		return nil
	}

	start := lb.Current
	for i := 0; i < len(servers); i++ {
		idx := start % len(servers)
		nextServer := servers[idx]
		lb.Current++

		nextServer.Mutex.Lock()
		isHealthy := nextServer.IsHealthy
		nextServer.Mutex.Unlock()

		if isHealthy {
			return nextServer
		}

		start++
	}

	return nil
}
