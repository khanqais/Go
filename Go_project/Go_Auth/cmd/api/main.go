package main

import (
	"context"
	"go-auth/internal/app"
	"go-auth/internal/httpserver"
	"log"
	"net/http"
	"time"
)

func main() {

	ctx := context.Background()
	a, err := app.New(ctx)
	if err != nil {
		log.Fatal("start up error")
	}
	defer func() {
		if err := a.Close(ctx); err != nil {
			log.Printf("shutdown warning: %v ", err)
		}
	}()
	router := httpserver.NewRouter()
	//standard go type that run a http server
	srv := &http.Server{
		Addr:              ":5000",
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}
	log.Printf("Api is running on %s", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Printf("Server Closed")
			return
		}
		log.Fatal("Server Error")
	}
}
