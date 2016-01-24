package pageactions

import (
	"github.com/fragmenta/router"
	"github.com/fragmenta/view"

	"github.com/gophergala2016/sendto/server/src/lib/authorise"
	"github.com/gophergala2016/sendto/server/src/pages"
)

// HandleCreateShow serves the create form via GET for pages
func HandleCreateShow(context router.Context) error {

	// Authorise
	err := authorise.Path(context)
	if err != nil {
		return router.NotAuthorizedError(err)
	}

	// Render the template
	view := view.New(context)
	page := pages.New()
	view.AddKey("page", page)

	return view.Render()
}

// HandleCreate handles the POST of the create form for pages
func HandleCreate(context router.Context) error {

	// Authorise
	err := authorise.ResourceAndAuthenticity(context, nil)
	if err != nil {
		return router.NotAuthorizedError(err)
	}

	// Setup context
	params, err := context.Params()
	if err != nil {
		return router.InternalError(err)
	}

	id, err := pages.Create(params.Map())
	if err != nil {
		return router.InternalError(err)
	}

	// Log creation
	context.Logf("#info Created page id,%d", id)

	// Redirect to the new page
	m, err := pages.Find(id)
	if err != nil {
		return router.InternalError(err)
	}

	return router.Redirect(context, m.URLIndex())
}
