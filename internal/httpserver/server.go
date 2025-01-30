package httpserver

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func StartHTTPServer(ctx context.Context) {
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		<-ctx.Done()
		fmt.Println("Stopping HTTP server...")
		err := server.Shutdown(context.Background())
		if err != nil {
			return
		}
	}()

	fmt.Println("HTTP server running on :8080")
	if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("HTTP server error: %s\n", err)
	}
}
