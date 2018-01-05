package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/mbonell/rpc-server/server"
)

func main() {
	var service server.RemoteMethods
	s := new(server.Service)
	service = s
	rpc.Register(service)
	rpc.HandleHTTP()

	l, err := net.Listen("tcp", "localhost:8767")
	if err != nil {
		log.Fatal("Listen error:", err)
	}

	log.Println("Server started!")
	log.Println(l.Addr())
	err = http.Serve(l, nil)
	if err != nil {
		log.Fatal("Serve error:", err)
	}
}
