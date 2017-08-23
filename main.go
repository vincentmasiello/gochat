package main

import (
	"flag"
	"mychat/lib"
	"os"
)

func main() {
	var isHost bool

	flag.BoolVar(&isHost, "listen", false, "Listens on the specified IP address")
	flag.Parse()

	if isHost {
		// go run main.go -listen <ip>
		connIP := os.Args[2] // [1] will be -listen flag in this case
		lib.RunHost(connIP)
	} else {
		// go run main.go <ip>
		connIP := os.Args[1]
		lib.RunGuest(connIP)
	}
}
