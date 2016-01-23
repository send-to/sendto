package fileactions

import (
	"github.com/fragmenta/router"
	"github.com/fragmenta/view"

	"github.com/gophergala2016/sendto/server/src/files"
	"github.com/gophergala2016/sendto/server/src/lib/authorise"
)

// HandleShow displays a single file
func HandleShow(context router.Context) error {

	// Find the file
	file, err := files.Find(context.ParamInt("id"))
	if err != nil {
		return router.InternalError(err)
	}

	// Authorise access
	err = authorise.Resource(context, file)
	if err != nil {
		return router.NotAuthorizedError(err)
	}

	// Render the template
	view := view.New(context)
	view.AddKey("file", file)
	return view.Render()
}
