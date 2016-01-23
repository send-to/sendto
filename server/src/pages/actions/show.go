package pageactions

import (
	"github.com/fragmenta/router"
	"github.com/fragmenta/view"

	"github.com/gophergala2016/sendto/server/src/lib/authorise"
	"github.com/gophergala2016/sendto/server/src/pages"
)

// HandleShow displays a single page
func HandleShow(context router.Context) error {

	// Find the page
	page, err := pages.Find(context.ParamInt("id"))
	if err != nil {
		return router.InternalError(err)
	}

	// Authorise access
	err = authorise.Resource(context, page)
	if err != nil {
		return router.NotAuthorizedError(err)
	}

	// Render the template
	view := view.New(context)
	view.AddKey("page", page)
	return view.Render()
}
