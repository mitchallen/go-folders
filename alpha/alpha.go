/**
 * Author: Mitch Allen
 * File: alpha.go
 */

package alpha

import (
	"fmt"

	"github.com/mitchallen/folders/beta"
	"github.com/mitchallen/folders/gamma"
)

func RunAlpha() bool {
	fmt.Println("[alpha] RunAlpha")
	fmt.Println("... calling siblings ...")
	beta.RunBeta()
	gamma.RunGamma()
	return true
}
