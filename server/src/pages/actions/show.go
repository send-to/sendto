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

	return render(context, page)
}

// HandleShowPath serves requests to a custom page url
func HandleShowPath(context router.Context) error {

	// Setup context for template
	path := context.Path()

	q := pages.Query().Where("url=?", path).Limit(1)
	pages, err := pages.FindAll(q)
	if err != nil || len(pages) == 0 {
		return router.NotFoundError(err)
	}

	// Get the first of pages to render
	page := pages[0]

	// If not published, check authorisation
	if !page.IsPublished() {
		// Authorise
		err = authorise.Resource(context, page)
		if err != nil {
			return router.NotAuthorizedError(err)
		}
	}

	return render(context, page)
}

// render the given page, with metadata set
func render(context router.Context, page *pages.Page) error {
	// Render the template
	view := view.New(context)
	view.AddKey("page", page)
	view.AddKey("page", page)
	view.AddKey("meta_title", page.Name)
	view.AddKey("meta_desc", page.Summary)
	view.AddKey("meta_keywords", page.Keywords)
	view.Template("pages/views/show.html.got")
	return view.Render()
}
