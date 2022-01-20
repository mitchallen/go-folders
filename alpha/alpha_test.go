/**
 * Author: Mitch Allen
 * File: alpha_test.go
 */

package alpha

import (
	"testing"
)

func TestAlpha(t *testing.T) {
	if got := RunAlpha(); got != true && got != false {
		t.Errorf("RunAlpha() = %t, didn't return true or false", got)
	}
}
