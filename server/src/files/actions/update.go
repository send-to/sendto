package fileactions

import (
	"github.com/fragmenta/router"
	"github.com/fragmenta/view"

	"github.com/gophergala2016/sendto/server/src/files"
	"github.com/gophergala2016/sendto/server/src/lib/authorise"
)

// HandleUpdateShow renders the form to update a file
func HandleUpdateShow(context router.Context) error {

	// Find the file
	file, err := files.Find(context.ParamInt("id"))
	if err != nil {
		return router.NotFoundError(err)
	}

	// Authorise update file
	err = authorise.ResourceAndAuthenticity(context, file)
	if err != nil {
		return router.NotAuthorizedError(err)
	}

	// Render the template
	view := view.New(context)
	view.AddKey("file", file)
	return view.Render()
}

// HandleUpdateShow handles the POST of the form to update a file
func HandleUpdate(context router.Context) error {

	// Find the file
	file, err := files.Find(context.ParamInt("id"))
	if err != nil {
		return router.NotFoundError(err)
	}

	// Authorise update file
	err = authorise.Resource(context, file)
	if err != nil {
		return router.NotAuthorizedError(err)
	}

	// Update the file from params
	params, err := context.Params()
	if err != nil {
		return router.InternalError(err)
	}

	// Find the to user, by querying users on username or email
	// Set the user id if found, else return 404 error, user not found

	// TODO: Make *sure* this only accepts the params we want
	err = file.Update(params.Map())
	if err != nil {
		return router.InternalError(err)
	}

	// Redirect to file
	return router.Redirect(context, file.URLShow())
}
