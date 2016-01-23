package app

import (
	"github.com/fragmenta/router"
	"github.com/gophergala2016/sendto/server/src/files/actions"
)

// Define routes for this app
func setupRoutes(r *router.Router) {

	r.Add("/files", fileactions.HandleIndex)
	r.Add("/files/create", fileactions.HandleCreateShow)
	r.Add("/files/create", fileactions.HandleCreate).Post()
	r.Add("/files/{id:[0-9]+}/update", fileactions.HandleUpdateShow)
	r.Add("/files/{id:[0-9]+}/update", fileactions.HandleUpdate).Post()
	r.Add("/files/{id:[0-9]+}/destroy", fileactions.HandleDestroy).Post()
	r.Add("/files/{id:[0-9]+}", fileactions.HandleShow)

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
