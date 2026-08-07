package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

const version = "0.0.1"

type config struct {
	port int
	env  string
	db   struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}
}

type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&cfg.db.dsn, "db", "postgres://greenlight:password@localhost/greenlight?sslmode=disable", "Postgres DSN")
	flag.IntVar(&cfg.db.maxOpenConns, "db.max-open-conns", 25, "Maximum number of open database connections")
	flag.IntVar(&cfg.db.maxIdleConns, "db.max-idle-conns", 25, "Maximum number of idle database connections")
	flag.StringVar(&cfg.db.maxIdleTime, "db.max-idle-time", "15m", "Maximum amount of time a database connection may be idle")
	flag.Parse()

	l := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := openDB(cfg)
	if err != nil {
		l.Fatal(err)
	}
	defer db.Close()
	l.Println("Connected to database")

	app := &application{
		config: cfg,
		logger: l,
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	l.Printf("Starting %s server on port %s", cfg.env, server.Addr)
	err = server.ListenAndServe()
	l.Fatal(err)
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.db.maxOpenConns)
	db.SetMaxIdleConns(cfg.db.maxIdleConns)

	duration, err := time.ParseDuration(cfg.db.maxIdleTime)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
