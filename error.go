package pipesample

import (
	"fmt"
)

// CastError returns a cast error.
func CastError(input ...interface{}) error {
	return fmt.Errorf("Cast failure %v", input)
}
