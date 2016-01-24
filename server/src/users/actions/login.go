package useractions

import (
	"fmt"

	"github.com/fragmenta/auth"
	"github.com/fragmenta/router"
	"github.com/fragmenta/view"

	"github.com/gophergala2016/sendto/server/src/lib/authorise"
	"github.com/gophergala2016/sendto/server/src/users"
)

// HandleLoginShow handles GET /users/login
func HandleLoginShow(context router.Context) error {

	// Check we have no current user
	u := authorise.CurrentUser(context)
	if !u.Anon() {
		return router.Redirect(context, fmt.Sprintf("/users/%d", u.Id))
	}

	// Render the template
	view := view.New(context)
	view.AddKey("error", context.Param("error"))
	return view.Render()
}

// HandleLogin handles a post to /users/login
func HandleLogin(context router.Context) error {

	params, err := context.Params()
	if err != nil {
		return router.NotFoundError(err)
	}

	// Check users against their username - we could also check against the email later?
	name := params.Get("name")
	q := users.Where("name=?", name)
	user, err := users.FindFirst(q)
	if err != nil {
		context.Logf("#error Login failed for user : %s %s", name, err)
		return router.Redirect(context, "/users/login?error=failed_name")
	}

	err = auth.CheckPassword(params.Get("password"), user.Password)
	if err != nil {
		context.Logf("#error Login failed for user : %s %s", name, err)
		return router.Redirect(context, "/users/login?error=failed_password")
	}

	// Save the details in a secure cookie
	session, err := auth.Session(context, context.Request())
	if err != nil {
		return router.InternalError(err)
	}

	context.Logf("#info LOGIN for user: %d", user.Id)
	session.Set(auth.SessionUserKey, fmt.Sprintf("%d", user.Id))
	session.Save(context)

	// Send them to their user profile page
	return router.Redirect(context, user.URLShow())

}
