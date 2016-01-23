// Tests for the pages actions package
package pageactions

import (
	"testing"
)

// Log a failure message, given msg, expected and result
func fail(t *testing.T, msg string, expected interface{}, result interface{}) {
	t.Fatalf("\n------FAILURE------\nTest failed: %s expected:%v result:%v", msg, expected, result)
}

// Test create of Page
func TestCreatePage(t *testing.T) {

}

// Test update of Page
func TestUpdatePage(t *testing.T) {

}
