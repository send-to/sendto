package fileactions

import (
	"github.com/fragmenta/router"
	"github.com/fragmenta/view"

	"github.com/gophergala2016/sendto/server/src/files"
	"github.com/gophergala2016/sendto/server/src/users"
)

// HandleCreate handles the POST of the create form for files
func HandleCreate(context router.Context) error {

	// Do not perform auth on posts - we could check a shared secret here or similar but that is not secure
	// better to require siging of posts by users with their own key to confirm identity if we wanted to check submissions.

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

	// First find the username, if we have no user, reject post
	context.Logf("PARAMS:%v", params)
	// PARAMS:map[sender:[Kenny Grant] recipient:[testtest]]
	// Check for user with name recipient

	// Find the user
	user, err := users.FindName(context.Param("recipient"))
	if err != nil || user == nil {
		return router.NotFoundError(err, "User not found", "Sorry this user doesn't exist")
	}

	// link with the named recipient user
	params.SetInt("user_id", user.Id)

	context.Logf("PARAMS:%v", params)

	// Ideally perform some identity check on the sender here, and set sender id if we have a user?
	// Perhaps require sending pgp sig of data?

	// We only consider the first file
	id, err := files.Create(params.Map(), fh)
	if err != nil {
		return router.InternalError(err)
	}

	// Log creation
	context.Logf("#info Created file id,%d", id)

	/*
		_, err := files.Find(id)
		if err != nil {
			return router.InternalError(err)
		}
	*/

	// Render a 200 response
	view := view.New(context)
	view.Layout("files/views/create.json.got")
	return view.Render()
	//	return router.Redirect(context, m.URLIndex())
}
