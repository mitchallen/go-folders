/**
 * Author: Mitch Allen
 * File: gamma.go
 */

package gamma

import (
	"fmt"

	"github.com/mitchallen/folders/beta"
)

func RunGamma() bool {
	fmt.Println("[gamma] RunGamma")
	beta.RunBeta()
	return true
}
