package useractions

import (
	"io"

	"github.com/fragmenta/router"

	"github.com/gophergala2016/sendto/server/src/users"
)

// HandleShowKey displays a single user's key
func HandleShowKey(context router.Context) error {

	// Find the user
	user, err := users.FindName(context.Param("name"))
	if err != nil {
		return router.InternalError(err)
	}

	// Render the key directly to the httpwriter as text
	context.Writer().Header().Set("Content-Type", "text/plain; charset=utf-8")
	_, err = io.WriteString(context.Writer(), user.Key)
	return err
}
