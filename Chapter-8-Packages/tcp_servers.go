package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

func server() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleServerConnection(conn)
	}
}

func handleServerConnection(conn net.Conn) {
	var msg string
	err := gob.NewDecoder(conn).Decode(&msg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Received:", msg)
	}
	conn.Close()
}

func client() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	msg := "hello world"
	fmt.Println("Sending", msg)
	err = gob.NewEncoder(conn).Encode(msg)
	if err != nil {
		fmt.Println(err)
	}
	conn.Close()
}

func main() {
	go server()
	go client()
	var input string
	fmt.Scanln(&input)
}
