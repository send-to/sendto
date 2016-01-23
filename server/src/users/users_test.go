// Tests for the users package
package users

import (
	"testing"
)

// Log a failure message, given msg, expected and result
func fail(t *testing.T, msg string, expected interface{}, result interface{}) {
	t.Fatalf("\n------FAILURE------\nTest failed: %s expected:%v result:%v", msg, expected, result)
}

// Test create of User
func TestCreateUser(t *testing.T) {

}

// Test update of User
func TestUpdateUser(t *testing.T) {

}
