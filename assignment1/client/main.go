package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	c, err := net.Dial("tcp", "127.0.0.1:2000")
	if err != nil {
		fmt.Println(err)
		return
	}

	// The entire for loop that is used to read user input will only terminate when you send the STOP command to the TCP server.
	for {

		// bufio.NewReader(os.Stdin) and ReadString() is used to read user input. Any user input is sent to the TCP server over the network using Fprintf().
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("send the message>> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(c, text+"\n")

		// bufio reader and the bufio.NewReader(c).ReadString('\n') statement read the TCP server’s response. The error variable is ignored here for simplicity.
		message, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Print("->: " + message)
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting...")
			return
		}
	}
}
