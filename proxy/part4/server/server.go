package server

import (
	"bufio"
	"net"
	"sync"
	"time"
)

type Backend struct {
	net.Conn
	Reader *bufio.Reader
	Writer *bufio.Writer
}

var backendQueue chan *Backend
var requestBytes map[string]int64
var requestLock sync.Mutex

func init(){
	requestBytes = make(map[string]int64)
	backendQueue = make(chan *Backend, 10)
}

func GetBackend() (*Backend, error){
	select {
	case be := <-backendQueue:
		return be, nil
	case <-time.After(100 * time.Millisecond):
		be, err := net.Dial("tcp", "127.0.0.1:8081")
		if err != nil{
			return nil, err
		}
		return &Backend{
			Conn: be,
			Reader: bufio.NewReader(be),
			Writer: bufio.NewWriter(be),
		},nil
	}
}

func QueueBackend(be *Backend){
	select {
	case backendQueue <- be:
		// Backend re-enqueued safely, move on
	case <-time.After(1 * time.Second):
		be.Close()
	}
}