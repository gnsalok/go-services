package main

import (
	"auth/data"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const webPort = "80"

var counts int64

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	log.Println("Starting auth service")

	// Connect to database

	conn := connectToDb()

	if conn == nil {
		log.Panic("Can't connect to Postgres!")
	}

	// Setup config
	app := Config{
		DB:     conn,
		Models: data.New(conn),
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()

	if err != nil {
		log.Panic(err)
	}

}

// open connection to database
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil

}

func connectToDb() *sql.DB {

	dsn := os.Getenv("DSN")

	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("Postgress is not ready yet...")
			counts++
		} else {
			log.Println("Connected to Postgres.")
			return connection
		}

		// try for 10 times means 20 seconds
		if counts > 10 {
			log.Println(err)
			return nil
		}

		log.Println("Backing off for two seconds")
		time.Sleep(2 * time.Second)
		continue
	}
}
