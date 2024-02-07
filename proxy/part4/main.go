package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"net/http"
	"part4/server"
	"strconv"
	"sync"
)

func main(){
	ln, err := net.Listen("tcp", ":8080")
	if err != nil{
		log.Fatalf("Failed to Listen: %s", err)
	}
	for{
		if conn, err := ln.Accept(); err == nil{
			go handleConnection(conn)
		}
	}
}

var requestBytes map[string] int64
var requestLock sync.Mutex

func init(){
	requestBytes = make(map[string]int64)
}

func updateStatus(req *http.Request, resp *http.Response) int64{
	requestLock.Lock()
	defer requestLock.Unlock()
	bytes := requestBytes[req.URL.Path] + resp.ContentLength
	requestBytes[req.URL.Path] = bytes
	return bytes
}

func handleConnection(conn net.Conn){
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		req, err := http.ReadRequest(reader)
		if err != nil{
			if err != io.EOF{
				log.Printf("Failed to read request: %s", err)
			}
			return
		}
		be, err := server.GetBackend()
		if err != nil{
			return
		}
		if err := req.Write(be.Writer); err == nil{
			be.Writer.Flush()
			if resp, err := http.ReadResponse(be.Reader, req); err == nil{
				bytes := updateStatus(req,resp)
				resp.Header.Set("X-Bytes", strconv.FormatInt(bytes,10))
				if err := resp.Write(conn); err == nil{
					log.Printf("%s: %d", req.URL.Path, resp.StatusCode)
				}
				
			}
		}
		go server.QueueBackend(be)
		
	}

}

//This version enhances the proxy server by tracking the total bytes transferred for each URL path and including this information in the response header.

//Additionally, it updates the status by adding the content length of the response to the total bytes transferred for the specific URL path.

//in simple terms it means it keeps tracks of the total amount of data transferred for each unique URL path and includes this information in the response header. 

//This version integrates a backend server management system, allowing the proxy server to retrieve and queue backend connections.

//This enhanced proxy server not only forwards requests and responses but also manages backend servers efficiently. It ensures that backend servers are obtained and released appropriately, allowing for better resource utilization and handling of varying loads.

 
//It provides a 

//GetBackend function--- to get a connection from the pool  and 

//QueueBackend function--- to return a connection to the pool. 

//The handleConnection function reads the client's HTTP request from the connection, gets a backend server connection from the server package, forwards the client's HTTP request to the backend server, reads the backend server's HTTP response, updates the status, and sends the response back to the client. The handleConnection function also logs the requested URL path and the response status code. 


//This enhanced proxy server not only forwards requests and responses but also manages backend servers efficiently. It ensures that backend servers are obtained and released appropriately, allowing for better resource utilization and handling of varying loads.


// The server package also initializes the connection pool with a single backend server at port 8081. The server package is used to manage the backend server connections in the handleConnection function.

