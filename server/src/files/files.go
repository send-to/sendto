// Package files represents the file resource
package files

import (
	"fmt"
	"mime/multipart"
	"path"
	"time"

	"github.com/fragmenta/model"
	"github.com/fragmenta/model/file"
	"github.com/fragmenta/model/validate"
	"github.com/fragmenta/query"

	"github.com/gophergala2016/sendto/server/src/lib/status"
)

// File handles saving and retreiving files from the database
type File struct {
	model.Model
	status.ModelStatus
	Path     string
	Sender   string
	SenderId int64
	Status   int64
	UserId   int64
}

// AllowedParams returns an array of allowed param keys
func AllowedParams() []string {
	return []string{"status", "path", "sender", "sender_id", "status", "user_id"}
}

// NewWithColumns creates a new file instance and fills it with data from the database cols provided
func NewWithColumns(cols map[string]interface{}) *File {

	file := New()
	file.Id = validate.Int(cols["id"])
	file.CreatedAt = validate.Time(cols["created_at"])
	file.UpdatedAt = validate.Time(cols["updated_at"])
	file.Status = validate.Int(cols["status"])
	file.Path = validate.String(cols["path"])
	file.Sender = validate.String(cols["sender"])
	file.SenderId = validate.Int(cols["sender_id"])
	file.Status = validate.Int(cols["status"])
	file.UserId = validate.Int(cols["user_id"])

	return file
}

// New creates and initialises a new file instance
func New() *File {
	file := &File{}
	file.Model.Init()
	file.Status = status.Draft
	file.TableName = "files"
	return file
}

// Create inserts a new record in the database using params, and returns the newly created id
func Create(params map[string]string, fh *multipart.FileHeader) (int64, error) {

	// Remove params not in AllowedParams
	params = model.CleanParams(params, AllowedParams())

	// Check params for invalid values
	err := validateParams(params)
	if err != nil {
		return 0, err
	}

	// Update date params
	params["created_at"] = query.TimeString(time.Now().UTC())
	params["updated_at"] = query.TimeString(time.Now().UTC())

	// Create the file
	id, err := Query().Insert(params)

	if fh != nil && id != 0 {

		// Now retrieve and save the file representation
		f, err := Find(id)
		if err != nil {
			return id, err
		}

		// Save files to disk using the passed in file data (if any).
		err = f.saveFile(fh)
		if err != nil {
			return id, err
		}
	}

	return id, err
}

// SaveFile saves files to disk
func (m *File) saveFile(fh *multipart.FileHeader) error {

	// Retreive the form image data by opening the referenced tmp file.
	f, err := fh.Open()
	if err != nil {
		return err
	}

	// If we have no path, set it to a default value /files/documents/id/name
	if len(m.Path) == 0 {
		err = m.NewFilePath(fh)
		if err != nil {
			return err
		}
	}

	// Make sure our path exists first
	err = file.CreatePathTo(m.Path)
	if err != nil {
		return err
	}

	// Write out our file to disk
	return file.Save(f, m.Path)

}

// NewFilePath sets a default path for this file.
func (m *File) NewFilePath(fh *multipart.FileHeader) error {
	m.Path = fmt.Sprintf("files/%d/%s", m.Id, file.SanitizeName(fh.Filename))

	// Perhaps tidy this to happen before create?
	return m.Update(map[string]string{"path": m.Path})
}

// validateParams checks these params pass validation checks
func validateParams(params map[string]string) error {

	// Now check params are as we expect
	err := validate.Length(params["id"], 0, -1)
	if err != nil {
		return err
	}
	err = validate.Length(params["name"], 0, 255)
	if err != nil {
		return err
	}

	return err
}

// Find returns a single record by id in params
func Find(id int64) (*File, error) {
	result, err := Query().Where("id=?", id).FirstResult()
	if err != nil {
		return nil, err
	}
	return NewWithColumns(result), nil
}

// FindAll returns all results for this query
func FindAll(q *query.Query) ([]*File, error) {

	// Fetch query.Results from query
	results, err := q.Results()
	if err != nil {
		return nil, err
	}

	// Return an array of files constructed from the results
	var files []*File
	for _, cols := range results {
		p := NewWithColumns(cols)
		files = append(files, p)
	}

	return files, nil
}

// Query returns a new query for files
func Query() *query.Query {
	p := New()
	return query.New(p.TableName, p.KeyName)
}

// Published returns a query for all files with status >= published
func Published() *query.Query {
	return Query().Where("status>=?", status.Published)
}

// Where returns a Where query for files with the arguments supplied
func Where(format string, args ...interface{}) *query.Query {
	return Query().Where(format, args...)
}

// Update sets the record in the database from params
func (m *File) Update(params map[string]string) error {

	// Remove params not in AllowedParams
	params = model.CleanParams(params, AllowedParams())

	// Check params for invalid values
	err := validateParams(params)
	if err != nil {
		return err
	}

	// Update date params
	params["updated_at"] = query.TimeString(time.Now().UTC())

	return Query().Where("id=?", m.Id).Update(params)
}

// Destroy removes the record from the database
func (m *File) Destroy() error {
	return Query().Where("id=?", m.Id).Delete()
}

// Name returns the file path basename
func (m *File) Name() string {
	return path.Base(m.Path)
}
