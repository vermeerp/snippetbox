package main

import (
	"github.com/alexedwards/scs"
	"github.com/vermeerp/snippetbox/pkg/models"
)

// App struct to hold the application-wide dependencies and configuration
// settings for our web application.
type App struct {
	Addr      string // Add an Addr field
	Database  *models.Database
	HTMLDir   string
	Sessions  *scs.Manager
	StaticDir string
	TLSCert   string // Add a TLSCert field
	TLSKey    string // Add a TLSKey field
}
