package useractions

import (
	"github.com/fragmenta/router"
	"github.com/fragmenta/view"

	"github.com/gophergala2016/sendto/server/src/lib/authorise"
	"github.com/gophergala2016/sendto/server/src/users"
)

// HandleCreateShow serves the create form via GET for users
func HandleCreateShow(context router.Context) error {

	// Authorise
	err := authorise.Path(context)
	if err != nil {
		return router.NotAuthorizedError(err)
	}

	// Render the template
	view := view.New(context)
	user := users.New()
	view.AddKey("user", user)

	return view.Render()
}

// HandleCreate handles the POST of the create form for users
func HandleCreate(context router.Context) error {

	// Authorise
	err := authorise.Path(context)
	if err != nil {
		return router.NotAuthorizedError(err)
	}

	// Setup context
	params, err := context.Params()
	if err != nil {
		return router.InternalError(err)
	}

	id, err := users.Create(params.Map())
	if err != nil {
		return router.InternalError(err)
	}

	// Log creation
	context.Logf("#info Created user id,%d", id)

	// Redirect to the new user
	m, err := users.Find(id)
	if err != nil {
		return router.InternalError(err)
	}

	return router.Redirect(context, m.URLIndex())
}
