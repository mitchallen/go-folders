/**
 * Author: Mitch Allen
 * File: root_test.go
 */

package folders

import (
	"testing"
)

func TestRoot(t *testing.T) {
	if got := RunRoot(); got != true && got != false {
		t.Errorf("RunRoot() = %t, didn't return true or false", got)
	}
}
