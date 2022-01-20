/**
 * Author: Mitch Allen
 * File: alpha.go
 */

package folders

import (
	"fmt"

	"github.com/mitchallen/folders/beta"
	"github.com/mitchallen/folders/gamma"
)

func RunAlpha() bool {
	fmt.Println("[alpha] RunAlpha")
	beta.RunBeta()
	gamma.RunGamma()

	return true
}
