package useractions

import (
	"github.com/fragmenta/router"

	"github.com/gophergala2016/sendto/server/src/lib/authorise"
	"github.com/gophergala2016/sendto/server/src/users"
)

// HandleDestroy handles a DESTROY request for users
func HandleDestroy(context router.Context) error {

	// Find the user
	user, err := users.Find(context.ParamInt("id"))
	if err != nil {
		return router.NotFoundError(err)
	}

	// Authorise destroy user
	err = authorise.Resource(context, user)
	if err != nil {
		return router.NotAuthorizedError(err)
	}

	// Destroy the user
	user.Destroy()

	// Redirect to users root
	return router.Redirect(context, user.URLIndex())
}
