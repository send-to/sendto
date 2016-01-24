// Package users represents the user resource
package users

import (
	"time"

	"github.com/fragmenta/auth"
	"github.com/fragmenta/model"
	"github.com/fragmenta/model/validate"
	"github.com/fragmenta/query"
	"github.com/fragmenta/router"

	"github.com/gophergala2016/sendto/server/src/lib/status"
)

// Roles for users
const (
	RoleAdmin    = 100
	RoleCustomer = 1
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
	return []string{"status", "email", "key", "name", "role", "summary", "password"}
}

// AllowedParamsCustomer returns an array of params that customers can edit on their user
func AllowedParamsCustomer() []string {
	return []string{"email", "password", "key"}
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

	// Check name is unique - no duplicate names allowed
	count, err := Query().Where("name=?", params["name"]).Count()
	if err != nil {
		return 0, err
	}

	if count > 0 {
		return 0, router.InternalError(err, "User name taken", "A username with this email already exists, sorry.")
	}

	// Update date params
	params["created_at"] = query.TimeString(time.Now().UTC())
	params["updated_at"] = query.TimeString(time.Now().UTC())

	return Query().Insert(params)
}

// validateParams checks these params pass validation checks
func validateParams(params map[string]string) error {

	err := validate.Length(params["name"], 0, 100)
	if err != nil {
		return router.InternalError(err, "Name invalid length", "Your name must be between 0 and 100 characters long")
	}

	err = validate.Length(params["key"], 1000, 5000)
	if err != nil {
		return router.InternalError(err, "Key too short", "Your key must be at least 1000 characters long")
	}

	// Password may be blank
	if len(params["password"]) > 0 {
		// check length
		err := validate.Length(params["password"], 5, 100)
		if err != nil {
			return router.InternalError(err, "Password too short", "Your password must be at least 5 characters")
		}

		// HASH the password before storage at all times
		hash, err := auth.HashPassword(params["password"])
		if err != nil {
			return err
		}

		params["password"] = hash

	} else {
		// Delete password param
		delete(params, "password")
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

// FindFirst fetches the first result for this query
func FindFirst(q *query.Query) (*User, error) {

	result, err := q.FirstResult()
	if err != nil {
		return nil, err
	}
	return NewWithColumns(result), nil
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

// Anon returns true if this user is not logged in
func (m *User) Anon() bool {
	return m.Role == 0
}

// Admin returns true for admin users
func (m *User) Admin() bool {
	return m.Role == 100
}

// IsUser return true if the user given is this user
func (m *User) IsUser(u *User) bool {
	return m.Id == u.Id
}

// OwnedBy returns true if the user id passed in owns this model
func (m *User) OwnedBy(uid int64) bool {
	return uid == m.Id
}
