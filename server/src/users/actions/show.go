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

	// NB no authorisation for access

	// Render the template
	view := view.New(context)
	view.AddKey("user", user)
	return view.Render()
}

// HandleShowName displays a single user by name
func HandleShowName(context router.Context) error {

	// Fetch the user by username
	q := users.Where("name=?", context.Param("name")).Limit(1)
	results, err := users.FindAll(q)
	if err != nil {
		return router.InternalError(err)
	}

	if len(results) != 1 {
		return router.NotFoundError(err, "User not found", "Sorry, this user could not be found")
	}

	// Get the first result
	user := results[0]

	// Authorise access
	err = authorise.Resource(context, user)
	if err != nil {
		return router.NotAuthorizedError(err)
	}

	// Render the template
	view := view.New(context)
	view.AddKey("user", user)
	view.Template("users/views/show.html.got")
	return view.Render()
}
