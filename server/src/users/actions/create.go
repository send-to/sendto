package useractions

import (
	"fmt"

	"github.com/fragmenta/auth"
	"github.com/fragmenta/router"
	"github.com/fragmenta/view"

	"github.com/gophergala2016/sendto/server/src/lib/authorise"
	"github.com/gophergala2016/sendto/server/src/lib/status"
	"github.com/gophergala2016/sendto/server/src/users"
)

// HandleCreateShow serves the create form via GET for users
func HandleCreateShow(context router.Context) error {

	// Render the template
	view := view.New(context)
	user := users.New()
	view.AddKey("user", user)

	return view.Render()
}

// HandleCreate handles the POST of the create form for users
func HandleCreate(context router.Context) error {

	// Authorise
	err := authorise.Resource(context, nil)
	if err != nil {
		return router.NotAuthorizedError(err)
	}

	// Setup context
	params, err := context.Params()
	if err != nil {
		return router.InternalError(err)
	}

	// Default to customer role etc - admins will have to promote afterwards
	params.Set("role", fmt.Sprintf("%d", users.RoleCustomer))
	params.Set("status", fmt.Sprintf("%d", status.Published))

	id, err := users.Create(params.Map())
	if err != nil {
		return err
	}

	// Log creation
	context.Logf("#info Created user id,%d", id)

	// Redirect to the new user
	user, err := users.Find(id)
	if err != nil {
		return router.InternalError(err)
	}

	// Save the details in a secure cookie
	session, err := auth.Session(context, context.Request())
	if err != nil {
		return router.InternalError(err)
	}

	context.Logf("#info CREATE for user: %d", user.Id)
	session.Set(auth.SessionUserKey, fmt.Sprintf("%d", user.Id))
	session.Save(context)

	// Send them to their user profile page
	return router.Redirect(context, user.URLShow())
}
