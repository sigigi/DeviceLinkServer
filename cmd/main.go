package main

import (
	"context"
	"fmt"
	"github.com/sigigi/DeviceLinkServer/internal/httpserver"
	"github.com/sigigi/DeviceLinkServer/internal/tcpserver"
	"github.com/sigigi/DeviceLinkServer/internal/udpserver"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	// Graceful Shutdown을 위해 OS 신호 수신
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
		<-sigChan
		fmt.Println("Shutting down servers...")
		cancel()
	}()

	// HTTP 서버 실행
	wg.Add(1)
	go func() {
		defer wg.Done()
		httpserver.StartHTTPServer(ctx)
	}()

	// TCP 서버 실행
	wg.Add(1)
	go func() {
		defer wg.Done()
		tcpserver.StartTCPServer(ctx)
	}()

	// UDP 서버 실행
	wg.Add(1)
	go func() {
		defer wg.Done()
		udpserver.StartUDPServer(ctx)

	}()

	wg.Wait()
	fmt.Println("All servers stopped.")
}
