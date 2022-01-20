/**
 * Author: Mitch Allen
 * File: gamma_test.go
 */

package gamma

import (
	"testing"
)

func TestGamma(t *testing.T) {
	if got := RunGamma(); got != true && got != false {
		t.Errorf("RunGamma() = %t, didn't return true or false", got)
	}
}
