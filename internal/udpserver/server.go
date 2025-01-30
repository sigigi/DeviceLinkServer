package udpserver

import (
	"context"
	"fmt"
	"net"
)

func StartUDPServer(ctx context.Context) {
	addr, err := net.ResolveUDPAddr("udp", ":9001")
	if err != nil {
		fmt.Printf("Error resolving UDP address: %s\n", err)
		return
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Printf("Error starting UDP server: %s\n", err)
		return
	}
	defer conn.Close()

	go func() {
		<-ctx.Done()
		fmt.Println("Stopping UDP server...")
		conn.Close()
	}()

	fmt.Println("UDP server running on :9001")
	buffer := make([]byte, 1024)
	for {
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			select {
			case <-ctx.Done():
				return
			default:
				fmt.Printf("UDP read error: %s\n", err)
			}
			continue
		}
		fmt.Printf("Received from %s: %s\n", addr, string(buffer[:n]))
	}
}
