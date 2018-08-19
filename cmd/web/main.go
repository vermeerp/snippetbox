package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/vermeerp/snippetbox/pkg/models"
)

func main() {

	// Define command-line flags for the network address and location of the static
	// files directory.
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "sb:u4UHCQQs#Agoqgi@/snippetbox?parseTime=true", "MySQL DSN")
	htmlDir := flag.String("html-dir", "./ui/html", "Path to HTML templates")
	staticDir := flag.String("static-dir", "./ui/static", "Path to static assets")
	flag.Parse()

	db := connect(*dsn)
	defer db.Close()

	// Initialize a new instance of App containing the dependencies.
	app := &App{
		Database:  &models.Database{db},
		HTMLDir:   *htmlDir,
		StaticDir: *staticDir,
	}

	log.Printf("Starting server on %s\n", *addr)
	err := http.ListenAndServe(*addr, app.Routes())
	log.Fatal(err)
}

// The connect() function wraps sql.Open() and returns a sql.DB connection pool
// for a given DSN.
func connect(dsn string) *sql.DB {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}
