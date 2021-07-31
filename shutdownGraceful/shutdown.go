package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome to Test Shutdown Server")
	})

	srv := http.Server{
		Addr:    ":18080",
		Handler: router,
	}

	go func() {
		// Start Server
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server start error: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	// Wait for close
	<-quit

	log.Println("Shutdown servering...")
	// 5s 内存处理结果，否则还是会退出
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// net/http 提供优雅停止
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("http shutdown error:", err)
	}
	log.Println("Server exit.")
}
