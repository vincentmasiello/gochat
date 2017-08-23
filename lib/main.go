package lib

import (
    "fmt"
    "log"
    "net"
    "bufio"
    "os"
)

const port = "8080"

// RunHost takes an ip as argument
// and listens for connections on that ip
func RunGuest(ip string) {
    ipAndPort := ip + ":" + port
    conn, dialErr := net.Dial("tcp", ipAndPort)
    if dialErr != nil {
        log.Fatal("Error: ", dialErr)
    }

    for {
        handleGuest(conn)
    }
}

// RunGuest takes destination ip as argument
// and connects to that ip
func RunHost(ip string) {
	ipAndPort := ip + ":" + port
	listener, listenErr := net.Listen("tcp", ipAndPort)
	if listenErr != nil {
		log.Fatal("Error: ", listenErr)
	}
    fmt.Println("Listening on ", ipAndPort)

	conn, acceptErr := listener.Accept()

	if acceptErr != nil {
		log.Fatal("Error: ", acceptErr)
	}
    fmt.Println("New Connection Accepted")

    for {
        handleHost(conn)
    }
}

func handleHost(conn net.Conn) {
	// read the message from the connection
	reader := bufio.NewReader(conn)
	message, readErr := reader.ReadString('\n') // accepts a delimiter for end of message (copy loud and clear, OVER)
	if readErr != nil {
		log.Fatal("Error: ", readErr)
	}

	fmt.Println("Message received: " + message)
    fmt.Print("SendMessage: ")

    replyReader := bufio.NewReader(os.Stdin)
    replyMessage, replyErr := replyReader.ReadString('\n')
    if replyErr != nil {
        log.Fatal("Error: ", replyMessage)
    }

    fmt.Fprint(conn, replyMessage)
}

func handleGuest(conn net.Conn) {
    
    fmt.Print("Send Message: ")

    reader := bufio.NewReader(os.Stdin)
    message, readErr := reader.ReadString('\n')
    if readErr != nil {
        log.Fatal("Error: ", readErr)
    }

    fmt.Fprint(conn, message)

    replyReader := bufio.NewReader(conn)
    replyMessage, replyErr := replyReader.ReadString('\n')
    if replyErr != nil {
        log.Fatal("Error: ", replyErr)
    }
    fmt.Println("Message received: ", replyMessage)
}
