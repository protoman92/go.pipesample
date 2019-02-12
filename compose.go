package pipesample

// Composable represents a function that takes input to produce output/error.
type Composable = func(interface{}) (interface{}, error)

// Compose pipes a sequence of functions.
func Compose(sources []Composable) Composable {
	return func(input interface{}) (interface{}, error) {
		input1 := input

		for _, source := range sources {
			output, err := source(input1)

			if err != nil {
				return nil, err
			}

			input1 = output
		}

		return input1, nil
	}
}
