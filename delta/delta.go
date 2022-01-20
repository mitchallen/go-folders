/**
 * Author: Mitch Allen
 * File: delta.go
 */

package delta

import (
	"fmt"

	"github.com/mitchallen/folders"
)

func RunDelta() bool {
	fmt.Println("[delta] RunDelta")
	fmt.Println("... calling parent ...")
	folders.RunRoot()
	return true
}
