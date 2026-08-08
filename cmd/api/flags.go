package main

import "flag"

func (c *config) parseFlags() {
	flag.IntVar(&c.port, "port", 4000, "API server port")
	flag.StringVar(&c.env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&c.db.dsn, "db", "postgres://greenlight:password@localhost/greenlight?sslmode=disable", "Postgres DSN")
	flag.IntVar(&c.db.maxOpenConns, "db.max-open-conns", 25, "Maximum number of open database connections")
	flag.IntVar(&c.db.maxIdleConns, "db.max-idle-conns", 25, "Maximum number of idle database connections")
	flag.StringVar(&c.db.maxIdleTime, "db.max-idle-time", "15m", "Maximum amount of time a database connection may be idle")
	flag.Parse()
}
