package setup

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

const PORT = 8080

func StartWS() {
	log.Printf("Starting server at port %d\n", PORT)
	router := gin.Default()

	for _, r := range routes {
		router.Handle(r.Method, "/user", r.Handler)
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", PORT),
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Println("Could not start listener", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	<-ctx.Done()
	log.Println("Timeout of 5 seconds.")
	log.Println("Server exiting")
}
