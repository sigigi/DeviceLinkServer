package tcpserver

import (
	"fmt"
	"log"
	"net"
)

func handleTCPConnection(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Printf("Error while closing the net.Conn: %v", err)
		}
	}(conn)

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Printf("TCP read error: %s\n", err)
		return
	}
	fmt.Printf("Received: %s\n", string(buffer[:n]))
}
