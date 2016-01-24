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

	// Do not perform auth on posts - we could check a shared secret here or similar but that is not secure
	// better to require siging of posts by users with their own key to confirm identity if we wanted to check submissions.

	/*
		// Authorise
		err := authorise.Path(context)
		if err != nil {
			return router.NotAuthorizedError(err)
		}
	*/

	// Parse multipart first - must fix this to do it automatically
	fileParams, err := context.ParamFiles("file")
	if err != nil || len(fileParams) < 1 {
		return router.InternalError(err, "Invalid file", "Sorry, the file upload failed.")
	}
	fh := fileParams[0]

	// Now extract other params
	params, err := context.Params()
	if err != nil {
		return router.InternalError(err)
	}

	// We only consider the first file
	id, err := files.Create(params.Map(), fh)
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
