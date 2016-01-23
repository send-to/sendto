package app

import (
	"github.com/fragmenta/router"
)

// Define routes for this app
func setupRoutes(r *router.Router) {

	// Set the default file handler
	r.FileHandler = fileHandler
	r.ErrorHandler = errHandler

	// Add a files route to handle static images under files
	// - nginx deals with this in production - perhaps only do this in dev?
	r.Add("/files/{path:.*}", fileHandler)
	r.Add("/favicon.ico", fileHandler)

	// Add the home page route
	r.Add("/", homeHandler)

}
