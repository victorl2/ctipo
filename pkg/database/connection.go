package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

const errMsgConnection = "The environment variable '%s' is not defined, it is required to establish a connection with the database"

var dBVariables = [...]string{"CTIPO_DB_NAME", "CTIPO_DB_HOST", "CTIPO_DB_USERNAME", "CTIPO_DB_PASSWORD"}

var dBConnection *pgxpool.Pool

func GetDBConnection() *pgxpool.Pool {
	if dBConnection != nil {
		return dBConnection
	}
	dBConnection = configToConnection(dBVariables)
	return dBConnection
}

func configToConnection(envNameList [4]string) *pgxpool.Pool {
	var dBVariables [4]string
	for i, envVariable := range envNameList {
		if variable := os.Getenv(envVariable); variable == "" {
			log.Fatalf(errMsgConnection, envNameList[i])
		} else {
			dBVariables[i] = variable
		}
	}
	return createDBConnection(dBVariables[0], dBVariables[1], dBVariables[2], dBVariables[3])
}

func createDBConnection(dbName, host, username, password string) *pgxpool.Pool {
	config, err := pgxpool.ParseConfig(fmt.Sprintf("user=%v password=%v host=%v port=5432 dbname=%v pool_max_conns=10", username, password, host, dbName))

	if err != nil {
		log.Fatal("error configuring the database: ", err)
	}

	conn, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}
	return conn
}
