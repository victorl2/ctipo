package main

import (
	"ctipo/pkg/api"
	"ctipo/pkg/database"
	"net/http"
	"os"
	"time"
)

func main() {
	os.Setenv("CTIPO_DB_NAME", "ctipo")
	os.Setenv("CTIPO_DB_USERNAME", "postgres")
	os.Setenv("CTIPO_DB_PASSWORD", "postgres")
	os.Setenv("CTIPO_DB_HOST", "ctipo.cjjaygkvsrbq.us-east-1.rds.amazonaws.com")

	if database.GetDBConnection().Stat().TotalConns() < 1 {
		panic("Coudln`t stablish a connection with the kate database")
	}

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
