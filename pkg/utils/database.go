package utils

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

//CreateDBConnection generates a connection with the postgres database
func CreateDBConnection() *pgxpool.Pool {
	//postgres://{username}:{password}@{hostname}:{port}/toby?pool_max_conns=10
	config, err := pgxpool.ParseConfig("postgres://postgres:postgres@localhost:5432/toby?pool_max_conns=10")
	if err != nil {
		log.Fatal("error configuring the database: ", err)
	}

	conn, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}
	return conn
}
