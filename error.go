package pipesample

import "errors"

var (
	castError error
)

func init() {
	castError = errors.New("Cast failure")
}
