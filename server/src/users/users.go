// Package users represents the user resource
package users

import (
	"time"

	"github.com/fragmenta/model"
	"github.com/fragmenta/model/validate"
	"github.com/fragmenta/query"

	"github.com/gophergala2016/sendto/server/src/lib/status"
)

// User handles saving and retreiving users from the database
type User struct {
	model.Model
	status.ModelStatus
	Email    string
	Key      string
	Name     string
	Password string
	Role     int64
	Summary  string
}

// AllowedParams returns an array of allowed param keys
func AllowedParams() []string {
	return []string{"status", "email", "key", "name", "password", "role", "summary"}
}

// NewWithColumns creates a new user instance and fills it with data from the database cols provided
func NewWithColumns(cols map[string]interface{}) *User {

	user := New()
	user.Id = validate.Int(cols["id"])
	user.CreatedAt = validate.Time(cols["created_at"])
	user.UpdatedAt = validate.Time(cols["updated_at"])
	user.Status = validate.Int(cols["status"])
	user.Email = validate.String(cols["email"])
	user.Key = validate.String(cols["key"])
	user.Name = validate.String(cols["name"])
	user.Password = validate.String(cols["password"])
	user.Role = validate.Int(cols["role"])
	user.Summary = validate.String(cols["summary"])

	return user
}

// New creates and initialises a new user instance
func New() *User {
	user := &User{}
	user.Model.Init()
	user.Status = status.Draft
	user.TableName = "users"
	return user
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
func Find(id int64) (*User, error) {
	result, err := Query().Where("id=?", id).FirstResult()
	if err != nil {
		return nil, err
	}
	return NewWithColumns(result), nil
}

// FindAll returns all results for this query
func FindAll(q *query.Query) ([]*User, error) {

	// Fetch query.Results from query
	results, err := q.Results()
	if err != nil {
		return nil, err
	}

	// Return an array of users constructed from the results
	var users []*User
	for _, cols := range results {
		p := NewWithColumns(cols)
		users = append(users, p)
	}

	return users, nil
}

// Query returns a new query for users
func Query() *query.Query {
	p := New()
	return query.New(p.TableName, p.KeyName)
}

// Published returns a query for all users with status >= published
func Published() *query.Query {
	return Query().Where("status>=?", status.Published)
}

// Where returns a Where query for users with the arguments supplied
func Where(format string, args ...interface{}) *query.Query {
	return Query().Where(format, args...)
}

// Update sets the record in the database from params
func (m *User) Update(params map[string]string) error {

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
func (m *User) Destroy() error {
	return Query().Where("id=?", m.Id).Delete()
}
