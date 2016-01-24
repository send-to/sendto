package app

import (
	"github.com/fragmenta/router"
	"github.com/gophergala2016/sendto/server/src/files/actions"
	"github.com/gophergala2016/sendto/server/src/pages/actions"
	"github.com/gophergala2016/sendto/server/src/users/actions"
)

// Define routes for this app
func setupRoutes(r *router.Router) {

	r.Add("/users", useractions.HandleIndex)
	r.Add("/users/create", useractions.HandleCreateShow)
	r.Add("/users/create", useractions.HandleCreate).Post()
	r.Add("/users/login", useractions.HandleLoginShow)
	r.Add("/users/login", useractions.HandleLogin).Post()
	r.Add("/users/logout", useractions.HandleLogout).Post()
	r.Add("/users/{id:[0-9]+}/update", useractions.HandleUpdateShow)
	r.Add("/users/{id:[0-9]+}/update", useractions.HandleUpdate).Post()
	r.Add("/users/{id:[0-9]+}/destroy", useractions.HandleDestroy).Post()
	r.Add("/users/{name:.+}/key", useractions.HandleShowKey)
	r.Add("/users/{id:[0-9]+}", useractions.HandleShow)
	r.Add("/users/{name:.+}", useractions.HandleShowName)

	r.Add("/pages", pageactions.HandleIndex)
	r.Add("/pages/create", pageactions.HandleCreateShow)
	r.Add("/pages/create", pageactions.HandleCreate).Post()
	r.Add("/pages/{id:[0-9]+}/update", pageactions.HandleUpdateShow)
	r.Add("/pages/{id:[0-9]+}/update", pageactions.HandleUpdate).Post()
	r.Add("/pages/{id:[0-9]+}/destroy", pageactions.HandleDestroy).Post()
	r.Add("/pages/{id:[0-9]+}", pageactions.HandleShow)

	r.Add("/files", fileactions.HandleIndex)
	r.Add("/files/create", fileactions.HandleCreate).Post()
	r.Add("/files/{id:[0-9]+}/update", fileactions.HandleUpdateShow)
	r.Add("/files/{id:[0-9]+}/update", fileactions.HandleUpdate).Post()
	r.Add("/files/{id:[0-9]+}/destroy", fileactions.HandleDestroy).Post()
	r.Add("/files/{id:[0-9]+}/download", fileactions.HandleDownload)

	// Set the default file handler
	r.FileHandler = fileHandler
	r.ErrorHandler = errHandler

	r.Add("/favicon.ico", fileHandler)

	// Add the custom page route
	r.Add("/{path:[a-z0-9]?}", pageactions.HandleShowPath)

}
