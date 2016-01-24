package fileactions

import (
	"github.com/fragmenta/router"
	"github.com/fragmenta/view"

	"github.com/gophergala2016/sendto/server/src/files"
	"github.com/gophergala2016/sendto/server/src/lib/authorise"
)

// HandleIndex displays a list of files
func HandleIndex(context router.Context) error {

	// Authorise
	err := authorise.Path(context)
	if err != nil {
		return router.NotAuthorizedError(err)
	}

	// Find the current user and check status
	u := authorise.CurrentUser(context)
	if u.Anon() {
		//	return router.NotAuthorizedError(err)
	}

	// For admins, show all files, order by date desc
	q := files.Query().Order("updated_at desc")

	// otherwise show just the logged in user's files
	if !u.Admin() {
		// Find the files for this user, unless
		q = files.Where("user_id=?", u.Id)
	}

	// Fetch the files
	results, err := files.FindAll(q)
	if err != nil {
		return router.InternalError(err)
	}

	// Render the template
	view := view.New(context)
	view.AddKey("files", results)
	return view.Render()

}
