package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	addr := net.UDPAddr{
		Port: 12345,        // Same port as used in the sender program
		IP:   net.IPv4zero, // Listen on all available IPs
	}

	// Create a UDP connection to listen for messages
	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error starting UDP listener: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Printf("Listening for broadcast messages on port %d...\n", addr.Port)

	buffer := make([]byte, 1024) // Buffer for incoming messages

	for {
		// Read from the UDP connection
		n, remoteAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading from UDP: %v\n", err)
			continue
		}

		// Print the received message and sender's address
		message := string(buffer[:n])
		fmt.Printf("Received message from %s: %s\n", remoteAddr, message)
	}
}
