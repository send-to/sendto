package fileactions

import (
	"net/http"

	"github.com/fragmenta/router"

	"github.com/gophergala2016/sendto/server/src/files"
	"github.com/gophergala2016/sendto/server/src/lib/authorise"
)

// HandleDownload sends a single file
func HandleDownload(context router.Context) error {

	// Find the file
	file, err := files.Find(context.ParamInt("id"))
	if err != nil {
		return router.InternalError(err)
	}

	// Authorise access to this file - only the file owners can access their own files
	err = authorise.Resource(context, file)
	if err != nil {
		return router.NotAuthorizedError(err)
	}

	// If we are permitted, send the file to the user
	http.ServeFile(context, context.Request(), file.Path)
	return nil
}
