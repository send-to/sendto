package useractions

import (
	"github.com/fragmenta/auth"
	"github.com/fragmenta/router"
)

// HandleLogout logs the current user out
func HandleLogout(context router.Context) error {

	// Clear the current session cookie
	auth.ClearSession(context)

	// Redirect to home
	return router.Redirect(context, "/")
}
