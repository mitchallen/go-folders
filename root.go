/**
 * Author: Mitch Allen
 * File: root.go
 */

package folders

import (
	"fmt"

	"github.com/mitchallen/folders/alpha"
	"github.com/mitchallen/folders/beta"
	"github.com/mitchallen/folders/gamma"
)

func RunRoot() bool {
	fmt.Println("[root] RunRoot")
	fmt.Println("... calling children ...")
	alpha.RunAlpha()
	beta.RunBeta()
	gamma.RunGamma()
	return true
}
