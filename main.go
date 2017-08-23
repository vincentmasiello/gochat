package main

import (
    "fmt"
    "flag"
    "net"
    "log"
    "os"
    "bufio"
)

func main() {
    var isHost bool

    flag.BoolVar(&isHost, "listen", false, "Listens on the specified IP address")
    flag.Parse()

    if isHost {
        // go run main.go -listen <ip>
        connIP := os.Args[2] // [1] will be -listen flag in this case
        runHost(connIP)
    } else {
        // go run main.go <ip>
        connIP := os.Args[1]
        runGuest(connIP)
    }

}

//temp const
const port = "8080"

func runGuest(ip string) {
    fmt.Println("look, I'm using fmt!")
}

func runHost(ip string) {
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
