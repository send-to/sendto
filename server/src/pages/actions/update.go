package pageactions

import (
	"github.com/fragmenta/router"
	"github.com/fragmenta/view"

	"github.com/gophergala2016/sendto/server/src/lib/authorise"
	"github.com/gophergala2016/sendto/server/src/pages"
)

// HandleUpdateShow renders the form to update a page
func HandleUpdateShow(context router.Context) error {

	// Find the page
	page, err := pages.Find(context.ParamInt("id"))
	if err != nil {
		return router.NotFoundError(err)
	}

	// Authorise update page
	err = authorise.Resource(context, page)
	if err != nil {
		return router.NotAuthorizedError(err)
	}

	// Render the template
	view := view.New(context)
	view.AddKey("page", page)

	return view.Render()
}

// HandleUpdate handles the POST of the form to update a page
func HandleUpdate(context router.Context) error {

	// Find the page
	page, err := pages.Find(context.ParamInt("id"))
	if err != nil {
		return router.NotFoundError(err)
	}

	// Authorise update page
	err = authorise.ResourceAndAuthenticity(context, page)
	if err != nil {
		return router.NotAuthorizedError(err)
	}

	// Update the page from params
	params, err := context.Params()
	if err != nil {
		return router.InternalError(err)
	}
	err = page.Update(params.Map())
	if err != nil {
		return router.InternalError(err)
	}

	// Redirect to page
	return router.Redirect(context, page.URLShow())
}
