// Tests for the files package
package files

import (
	"testing"
)

// Log a failure message, given msg, expected and result
func fail(t *testing.T, msg string, expected interface{}, result interface{}) {
	t.Fatalf("\n------FAILURE------\nTest failed: %s expected:%v result:%v", msg, expected, result)
}

// Test create of File
func TestCreateFile(t *testing.T) {

}

// Test update of File
func TestUpdateFile(t *testing.T) {

}
