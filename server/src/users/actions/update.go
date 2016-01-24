package useractions

import (
	"github.com/fragmenta/router"
	"github.com/fragmenta/view"

	"github.com/gophergala2016/sendto/server/src/lib/authorise"
	"github.com/gophergala2016/sendto/server/src/users"
)

// HandleUpdateShow renders the form to update a user
func HandleUpdateShow(context router.Context) error {

	// Find the user
	user, err := users.Find(context.ParamInt("id"))
	if err != nil {
		return router.NotFoundError(err)
	}

	// Authorise update user
	err = authorise.ResourceAndAuthenticity(context, user)
	if err != nil {
		return router.NotAuthorizedError(err)
	}

	// Render the template
	view := view.New(context)
	view.AddKey("user", user)

	return view.Render()
}

// HandleUpdate handles the POST of the form to update a user
func HandleUpdate(context router.Context) error {

	// Find the user
	user, err := users.Find(context.ParamInt("id"))
	if err != nil {
		return router.NotFoundError(err)
	}

	// Authorise update user
	err = authorise.Resource(context, user)
	if err != nil {
		return router.NotAuthorizedError(err)
	}

	// Get the params
	params, err := context.Params()
	if err != nil {
		return router.InternalError(err)
	}

	// Clean params further for customers, they may only update email, password, key
	allowedParams := params.Map()
	u := authorise.CurrentUser(context)
	if !u.Admin() {
		//	allowedParams = params.Clean(users.AllowedParamsCustomer())
	}

	err = user.Update(allowedParams)
	if err != nil {
		return router.InternalError(err)
	}

	// Redirect to user
	return router.Redirect(context, user.URLShow())
}
