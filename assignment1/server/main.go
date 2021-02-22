package main

import (
	"bufio"
	"fmt"
	"net"

	// "os"
	"strings"
	"time"
)

//This file creates the main package, which declares the main() function. The function will use the
// imported packages to create a TCP server.
func main() {

	l, err := net.Listen("tcp", ":2000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	//It is only after a successful call to Accept() that the TCP server can begin to interact with TCP clients.

	c, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	for {

		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		if strings.TrimSpace(string(netData)) == "STOP" {
			fmt.Println("Exiting TCP server!")
			return
		}

		fmt.Print("recieved the message >> ", string(netData))
		t := time.Now()
		myTime := t.Format(time.RFC3339) + "\n"
		c.Write([]byte(myTime))
	}
}
