/**
 * Author: Mitch Allen
 * File: delta_test.go
 */

package delta

import (
	"testing"
)

func TestDelta(t *testing.T) {
	if got := RunDelta(); got != true && got != false {
		t.Errorf("RunDelta() = %t, didn't return true or false", got)
	}
}
