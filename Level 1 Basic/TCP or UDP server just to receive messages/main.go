package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	// 1. ask the os to listen for incoming connection on port 8080

	//2. Handle potential error (e.g if port 8080 is already in use)
	if err != nil {
		fmt.Println("Error starting server", err)
		os.Exit(1)
	}

	defer listener.Close()

	fmt.Println("TCP Server is listening on port 8080...")

	//3. the infinite loop
	for {
		//4.wait for the client to connect
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error while accepitng the connection:", err)
			continue //Skip the rest of this loop and try to acceting again
		}

		fmt.Println("A new client connected!")

		//5. Create a buket to hold the incoming message ( upto 1024 bytes)
		buffer := make([]byte, 1024)

		//6. Read the data from the client into out bucket
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading data:", err)
			conn.Close()
			continue
		}

		// 7. COnvert the raw bytes into a string and print it
		fmt.Printf("Message received: %s\n", string(buffer[:n]))

		//8. Close the connection with this specific client
		conn.Close()
	}
}
