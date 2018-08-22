package main

import "github.com/vermeerp/snippetbox/pkg/models"

// App struct to hold the application-wide dependencies and configuration
// settings for our web application.
type App struct {
	Database  *models.Database
	HTMLDir   string
	StaticDir string
}
