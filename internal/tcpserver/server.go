package tcpserver

import (
	"context"
	"fmt"
	"log"
	"net"
)

func StartTCPServer(ctx context.Context) {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		fmt.Printf("Error starting TCP server: %s\n", err)
		return
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			log.Printf("Error while closing the listener: %v", err)
		}
	}(listener)

	go func() {
		<-ctx.Done()
		fmt.Println("Stopping TCP server...")
		err := listener.Close()
		if err != nil {
			return
		}
	}()

	fmt.Println("TCP server running on :9000")
	for {
		conn, err := listener.Accept()
		if err != nil {
			select {
			case <-ctx.Done():
				return
			default:
				fmt.Printf("TCP connection error: %s\n", err)
			}
			continue
		}
		go handleTCPConnection(conn)
	}
}
