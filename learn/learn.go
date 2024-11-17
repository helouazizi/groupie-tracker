package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	// Dial to a server
	conn, err := net.Dial("tcp", "zone01oujda.ma:8080")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()
	fmt.Println(conn.RemoteAddr())

	// Set a write deadline
	conn.SetWriteDeadline(time.Now().Add(5 * time.Second))

	// Write an HTTP request
	request := "GET / HTTPS/1.1\r\nHost: zone01oujda.ma/\r\n\r\n"
	_, err = conn.Write([]byte(request))
	if err != nil {
		fmt.Println("Error writing to connection:", err)
		return
	}

	// Set a read deadline
	conn.SetReadDeadline(time.Now().Add(5 * time.Second))

	// Read the response
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading from connection:", err)
		return
	}
	fmt.Println("Response:", string(buf[:n]))
}
