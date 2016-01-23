package useractions

import (
	"github.com/fragmenta/router"
	"github.com/fragmenta/view"

	"github.com/gophergala2016/sendto/server/src/lib/authorise"
	"github.com/gophergala2016/sendto/server/src/users"
)

// HandleShow displays a single user
func HandleShow(context router.Context) error {

	// Find the user
	user, err := users.Find(context.ParamInt("id"))
	if err != nil {
		return router.InternalError(err)
	}

	// Authorise access
	err = authorise.Resource(context, user)
	if err != nil {
		return router.NotAuthorizedError(err)
	}

	// Render the template
	view := view.New(context)
	view.AddKey("user", user)
	return view.Render()
}
