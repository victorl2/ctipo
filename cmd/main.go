package main

import (
	"ctipo/pkg/api"
	"net/http"
	"time"
)

func main() {
	routerHandler := api.InitRouter()

	server := &http.Server{
		Addr:           ":8080",
		Handler:        routerHandler,
		ReadTimeout:    time.Duration(10 * time.Second),
		WriteTimeout:   time.Duration(10 * time.Second),
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}
