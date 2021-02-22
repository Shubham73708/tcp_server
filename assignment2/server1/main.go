package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	fmt.Println("server listening on 9090")
	handler, err := net.Listen("tcp", "localhost:9090")
	if err != nil {
		log.Fatalln(err)
	}
	defer handler.Close()

	// handling for incoming connections
	for {
		conn, err := handler.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("new client connection")

		// handle connections in another gorutine
		go handleConnection(conn)
	}
}

// handling for messages from connection
func handleConnection(conn net.Conn) {
	for {
		buffer := make([]byte, 1400)
		dataSize, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("connection closed")
			return
		}

		// the actual message
		data := buffer[:dataSize]
		fmt.Println("message received:", string(data))
		// echoing the message back out
		_, err = conn.Write(data)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("Message sent:", string(data))
	}
}
