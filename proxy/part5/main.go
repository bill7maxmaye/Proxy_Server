package main

import (
	"log"
	"net/rpc"
	rpcserver "part5/rpc_server/server"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:8079")
	if err != nil{
		log.Fatalf("Failed to dial: %s", err)
	}

	var reply rpcserver.Stats
	err = client.Call("rpcserver.GetStats", &rpcserver.Empty{}, &reply)
	if err != nil{
		log.Fatalf("Failed to GetStats: %s",err)
	}
}




//The RPC (Remote Procedure Call) client is connecting to an RPC server that exposes methods for retrieving server statistics.

//The rpc.DialHTTP function establishes a connection to the RPC server using the TCP protocol at the specified address (127.0.0.1:8079).

//The client then calls the GetStats method on the server using client.Call. It sends an empty request (rpcserver.Empty{}) as an argument and expects the server to return server statistics in the form of rpcserver.Stats. The response is stored in the reply variable.

//If there's an error during the RPC call, the program logs a fatal error message.

//This code represents a client that connects to an RPC server, invokes a remote method (GetStats), and retrieves server statistics. The specifics of the RPC server methods and data types (rpcserver.Stats and rpcserver.Empty) are defined in the part5/rpc_server/server package. The client is responsible for handling errors and responding to the server's output.