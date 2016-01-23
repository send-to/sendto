package fileactions

import (
	"github.com/fragmenta/router"
	"github.com/fragmenta/view"

	"github.com/gophergala2016/sendto/server/src/files"
	"github.com/gophergala2016/sendto/server/src/lib/authorise"
)

// HandleCreateShow serves the create form via GET for files
func HandleCreateShow(context router.Context) error {

	// Authorise
	err := authorise.Path(context)
	if err != nil {
		return router.NotAuthorizedError(err)
	}

	// Render the template
	view := view.New(context)
	file := files.New()
	view.AddKey("file", file)
	return view.Render()
}

// HandleCreate handles the POST of the create form for files
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

	id, err := files.Create(params.Map())
	if err != nil {
		return router.InternalError(err)
	}

	// Log creation
	context.Logf("#info Created file id,%d", id)

	// Redirect to the new file
	m, err := files.Find(id)
	if err != nil {
		return router.InternalError(err)
	}

	return router.Redirect(context, m.URLIndex())
}
