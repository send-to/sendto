package pageactions

import (
	"github.com/fragmenta/router"

	"github.com/gophergala2016/sendto/server/src/lib/authorise"
	"github.com/gophergala2016/sendto/server/src/pages"
)

// HandleDestroy handles a DESTROY request for pages
func HandleDestroy(context router.Context) error {

	// Find the page
	page, err := pages.Find(context.ParamInt("id"))
	if err != nil {
		return router.NotFoundError(err)
	}

	// Authorise destroy page
	err = authorise.Resource(context, page)
	if err != nil {
		return router.NotAuthorizedError(err)
	}

	// Destroy the page
	page.Destroy()

	// Redirect to pages root
	return router.Redirect(context, page.URLIndex())
}
