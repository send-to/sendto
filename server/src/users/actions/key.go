package useractions

import (
	"github.com/fragmenta/router"
	"github.com/fragmenta/view"

	"github.com/gophergala2016/sendto/server/src/users"
)

// HandleShowKey displays a single user's key
func HandleShowKey(context router.Context) error {

	// Find the user
	user, err := users.FindName(context.Param("name"))
	if err != nil {
		return router.InternalError(err)
	}

	// Render the template
	view := view.New(context)
	view.AddKey("user", user)
	view.Layout("") // render plain text, no layout
	return view.Render()
}
