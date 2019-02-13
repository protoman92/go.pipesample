package pipesample

// Composable represents a function that takes input to produce output/error.
type Composable = func(interface{}) (interface{}, error)

// ComposableMapper represents a Composable converter.
type ComposableMapper = func(Composable) Composable

// Compose pipes a sequence of functions.
func Compose(sources ...Composable) Composable {
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

// ComposeMapper pipes a sequence of mappers
func ComposeMapper(mappers ...ComposableMapper) ComposableMapper {
	return func(composable Composable) Composable {
		composable1 := composable

		for _, mapper := range mappers {
			composable1 = mapper(composable1)
		}

		return composable1
	}
}
