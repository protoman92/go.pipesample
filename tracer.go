package pipesample

import (
	"fmt"
	"time"
)

// ILogger does logging.
type ILogger interface {
	Log(event interface{})
}

// Trace does some tracing for a function with a specified name.
func Trace(logger ILogger, funcName string) ComposableMapper {
	return func(composable Composable) Composable {
		return func(input interface{}) (interface{}, error) {
			startTime := time.Now()

			defer func() {
				elapsed := time.Now().Sub(startTime)
				logger.Log(fmt.Sprintf("%v took %v millis to run", funcName, elapsed))
			}()

			return composable(input)
		}
	}
}
