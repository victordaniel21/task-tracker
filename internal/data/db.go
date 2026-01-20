package data

import (
	"database/sql"
	"log"
	"time"

	// the "_" means we import the package for its "side effects"
	// (registering the driver), but we don't refer to the name "pq" in this code
	_ "github.com/lib/pq"
)

func OpenDB(dsn string) (*sql.DB, error) {
	// 1. open the pool (this doesn't actually connect yet, just validates argument)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// 2. Test the connection with a Ping, we set a context with a 5-second timeout.
	// if the DB doesn't respond in 5s we fail.
	if err = db.Ping(); err != nil {
		return nil, err
	}

	// 3. Configure the connection Pool
	// SetMaxOpenConns: Max number of open connections to the database.
	db.SetMaxOpenConns(25)
	// SetMaxIdleConns: Max number of connections in the idle pool.
	db.SetMaxIdleConns(25)
	// SetConnMaxLifetime: how long a connection can be reused
	db.SetConnMaxLifetime(5 * time.Minute)

	log.Print("✅ Database connection pool established")
	return db, nil

}
