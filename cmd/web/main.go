package main

import (
	"database/sql"
	"flag"
	"log"
	"time"

	"github.com/alexedwards/scs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/vermeerp/snippetbox/pkg/models"
)

func main() {

	// Define command-line flags for the network address and location of the static
	// files directory.
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "sb:u4UHCQQs#Agoqgi@/snippetbox?parseTime=true", "MySQL DSN")
	htmlDir := flag.String("html-dir", "./ui/html", "Path to HTML templates")
	secret := flag.String("secret", "s6Nd%+pPbnzHbS*+9Pk8qGWhTzbpa@ge", "Secret key")
	staticDir := flag.String("static-dir", "./ui/static", "Path to static assets")
	tlsCert := flag.String("tls-cert", "./tls/cert.pem", "Path to TLS certificate")
	tlsKey := flag.String("tls-key", "./tls/key.pem", "Path to TLS key")

	flag.Parse()

	db := connect(*dsn)
	defer db.Close()

	sessionManager := scs.NewCookieManager(*secret)
	sessionManager.Lifetime(12 * time.Hour)
	sessionManager.Persist(true)

	// Initialize a new instance of App containing the dependencies.
	app := &App{
		Addr:      *addr,
		Database:  &models.Database{DB: db},
		HTMLDir:   *htmlDir,
		Sessions:  sessionManager,
		StaticDir: *staticDir,
		TLSCert:   *tlsCert,
		TLSKey:    *tlsKey,
	}

	app.RunServer()

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
