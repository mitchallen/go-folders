/**
 * Author: Mitch Allen
 * File: beta_test.go
 */

package beta

import (
	"testing"
)

func TestBeta(t *testing.T) {
	if got := RunBeta(); got != true && got != false {
		t.Errorf("RunBeta() = %t, didn't return true or false", got)
	}
}
