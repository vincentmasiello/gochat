package lib

import (
    "fmt"
    "log"
    "net"
    "bufio"
)
const port = "8080"
// RunHost takes an ip as argument
// and listens for connections on that ip
func RunGuest(ip string) {
	fmt.Println("look, I'm using fmt!")
}
// RunGuest takes destination ip as argument
// and connects to that ip
func RunHost(ip string) {
	ipAndPort := ip + ":" + port
	listener, listenErr := net.Listen("tcp", ipAndPort)
	if listenErr != nil {
		log.Fatal("Error: ", listenErr)
	}

	conn, acceptErr := listener.Accept()

	if acceptErr != nil {
		log.Fatal("Error: ", acceptErr)
	}
	// read the message from the connection
	reader := bufio.NewReader(conn)
	message, readErr := reader.ReadString('\n') // accepts a delimiter for end of message (copy loud and clear, OVER)
	if readErr != nil {
		log.Fatal("Error: ", readErr)
	}

	fmt.Println("Message received: " + message)
}
