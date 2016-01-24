package authorise

import (
	"fmt"
	"strings"

	"github.com/fragmenta/auth"
	"github.com/fragmenta/router"
	"github.com/fragmenta/server"

	"github.com/gophergala2016/sendto/server/src/users"
)

// ResourceModel defines the interface for models passed to authorise.Resource
type ResourceModel interface {
	OwnedBy(int64) bool
}

// Setup authentication and authorization keys for this app
func Setup(s *server.Server) {

	// Set up our secret keys which we take from the config
	// NB these are hex strings which we convert to bytes, for ease of presentation in secrets file
	c := s.Configuration()
	auth.HMACKey = auth.HexToBytes(c["hmac_key"])
	auth.SecretKey = auth.HexToBytes(c["secret_key"])
	auth.SessionName = "sendto"

	// Enable https cookies on production server - we don't have https, so don't do this
	if s.Production() {
		s.Log("Using secure cookies")
		auth.SecureCookies = true
	}

}

// CurrentUserFilter returns a filter function which sets the current user on the context
func CurrentUserFilter(c router.Context) error {
	u := CurrentUser(c)
	c.Set("current_user", u)
	return nil
}

// Path authorises the path for the current user
func Path(c router.Context) error {
	return Resource(c, nil)
}

// ResourceAndAuthenticity authorises the path and resource for the current user
func ResourceAndAuthenticity(c router.Context, r ResourceModel) error {

	// Check the authenticity token first
	err := AuthenticityToken(c)
	if err != nil {
		return err
	}

	// Now authorise the resource as normal
	return Resource(c, r)
}

// Resource authorises the path and resource for the current user
// if model is nil it is ignored and permission granted
func Resource(c router.Context, r ResourceModel) error {

	// Short circuit evaluation if this is a public path
	if publicPath(c.Path()) {
		return nil
	}

	user := c.Get("current_user").(*users.User)

	switch user.Role {
	case users.RoleAdmin:
		return nil
	case users.RoleCustomer:

		// RoleCustomer should have access to /files
		if c.Path() == "/files" {
			return nil
		}
		// RoleCustomer should have access to /users/x/update if they are that user
		if strings.HasPrefix(c.Path(), "/users") {
			if r != nil && r.OwnedBy(user.Id) {
				return nil
			}
		}

	}

	return fmt.Errorf("Path and Resource not authorized:%s %v", c.Path(), r)
}

// publicPath returns true if this path should always be allowed, regardless of user role
func publicPath(p string) bool {
	switch p {
	case "/":
		return true
	case "/users/create":
	case "/users/login":
		return true
	}

	return false
}
