package models

import (
	"time"
)

// Snippet type to hold the information about an individual snippet.
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// Snippets type, which is a slice for holding multiple Snippet objects.
type Snippets []*Snippet
