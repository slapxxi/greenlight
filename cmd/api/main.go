package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/slapxxi/greenlight/internal/data"
)

const version = "0.0.1"

type config struct {
	port int
	env  string
	db   struct {
		// dsn is a connection string to the database
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}
}

type application struct {
	config config
	logger *log.Logger
	models data.Models
}

// NewServer returns a new HTTP server configured with the application
func (a *application) NewServer() *http.Server {
	return &http.Server{
		Addr:         fmt.Sprintf(":%d", a.config.port),
		Handler:      a.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
}

func main() {
	var cfg config
	cfg.parseFlags()

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
		models: data.NewModels(db),
	}

	server := app.NewServer()
	l.Printf("Starting %s server on port %s", cfg.env, server.Addr)
	err = server.ListenAndServe()
	l.Fatal(err)
}
