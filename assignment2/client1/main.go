package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

//Queue represents a queue that holds a slice
type Queue struct {
	items []string
}

//Enqueue adds the item in the Queue
func (q *Queue) Enqueue(text string) {
	q.items = append(q.items, text)

}

func main() {

	myQueue := Queue{}

	c, err := net.Dial("tcp", "127.0.0.1:9090")
	if err != nil {
		fmt.Println(err)
		return
	}

	reader := bufio.NewReader(os.Stdin)
	for {

		fmt.Print("message sent>>")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(c, text+"\n")
		myQueue.Enqueue(text)
		fmt.Println(myQueue)

		//message, _ := bufio.NewReader(c).ReadString('\n')
		//fmt.Print("message recieved>>" + message)
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("client exiting...")
			return
		}

	}
}
