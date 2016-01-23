// Package files represents the file resource
package files

import (
	"time"

	"github.com/fragmenta/model"
	"github.com/fragmenta/model/validate"
	"github.com/fragmenta/query"

	"github.com/fragmenta/fragmenta-cms/src/lib/status"
)

// File handles saving and retreiving files from the database
type File struct {
	model.Model
	status.ModelStatus
	From     string
	Path     string
	SignedBy int64
	UserId   int64
}

// AllowedParams returns an array of allowed param keys
func AllowedParams() []string {
	return []string{"status", "from", "path", "signed_by", "user_id"}
}

// NewWithColumns creates a new file instance and fills it with data from the database cols provided
func NewWithColumns(cols map[string]interface{}) *File {

	file := New()
	file.Id = validate.Int(cols["id"])
	file.CreatedAt = validate.Time(cols["created_at"])
	file.UpdatedAt = validate.Time(cols["updated_at"])
	file.Status = validate.Int(cols["status"])
	file.From = validate.String(cols["from"])
	file.Path = validate.String(cols["path"])
	file.SignedBy = validate.Int(cols["signed_by"])
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
func Create(params map[string]string) (int64, error) {

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

	return Query().Insert(params)
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
