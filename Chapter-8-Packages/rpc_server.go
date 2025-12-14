package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type Server struct{}

func (this *Server) Negate(i int64, reply *int64) error {
	*reply = -i
	return nil
}

func server_rpc() {
	rpc.Register(new(Server))
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(conn)
	}
}

func client_rpc() {
	conn, err := rpc.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	var result int64
	err = conn.Call("Server.Negate", int64(99), &result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("server negate result:", result)
	}
}

func main() {
	go server_rpc()
	go client_rpc()
	var input string
	fmt.Scanln(&input)
}
