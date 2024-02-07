package rpcserver

import (
	"sync"
)

type Empty struct{}
type Stats struct {
	RequestBytes map[string]int64
}

var requestBytes map[string]int64
var requestLock sync.Mutex

type RpcServer struct{}

// Exported method suitable for RPC
func (r *RpcServer) GetStats(args *Empty, reply *Stats) error {
	requestLock.Lock()
	defer requestLock.Unlock()

	reply.RequestBytes = make(map[string]int64)
	for k, v := range requestBytes {
		reply.RequestBytes[k] = v
	}

	return nil
}

