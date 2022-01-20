/**
 * Author: Mitch Allen
 * File: alpha_test.go
 */

package folders

import (
	"testing"
)

func TestFlip(t *testing.T) {
	if got := RunAlpha(); got != true && got != false {
		t.Errorf("RunAlpha() = %t, didn't return true or false", got)
	}
}
