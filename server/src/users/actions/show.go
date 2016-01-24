package useractions

import (
	"github.com/fragmenta/router"
	"github.com/fragmenta/view"

	"github.com/gophergala2016/sendto/server/src/users"
)

// NB no authorisation for access

// HandleShow displays a single user
func HandleShow(context router.Context) error {

	// Find the user
	user, err := users.Find(context.ParamInt("id"))
	if err != nil {
		return router.NotFoundError(err)
	}

	// Render the template
	view := view.New(context)
	view.AddKey("user", user)
	return view.Render()
}

// HandleShowName displays a single user by name
func HandleShowName(context router.Context) error {

	// Find the user
	user, err := users.FindName(context.Param("name"))
	if err != nil {
		return router.NotFoundError(err)
	}

	// Render the template
	view := view.New(context)
	view.AddKey("user", user)
	view.Template("users/views/show.html.got")
	return view.Render()
}
