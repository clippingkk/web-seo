package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/clippingkk/web-seo/routes"
	"github.com/sirupsen/logrus"
)

func main() {
	port := "4789"
	indexPageTemplate := os.Getenv("INDEX_TEM")

	h := routes.InitRoutes(indexPageTemplate)

	server := &http.Server{
		Addr:    "0.0.0.0:" + port,
		Handler: h,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("listen: %s\n", err)
		}
	}()
	logrus.Println("server running on :" + port)
	logrus.Println(fmt.Sprintf("GraphiQL: http://localhost:%s/api/v2/graphql", port))

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logrus.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logrus.Fatal("Server forced to shutdown:", err)
	}
	logrus.Println("Server exiting")

}
