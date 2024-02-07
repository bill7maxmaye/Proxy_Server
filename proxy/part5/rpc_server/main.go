package main



import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"part5/rpc_server/server"
)

func main(){
	rpc.Register(&rpcserver.RpcServer{})
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":8079")
	if err != nil{
		log.Fatalf("Failed to listen: %s", err)
	}
	go http.Serve(l, nil)
}