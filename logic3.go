package pipesample

// ILogic3Dependency serves as dependency for logic 3.
type ILogic3Dependency interface {
	TransformLogic3(input bool) (string, error)
}

// Logic3 performs logic 3.
func Logic3(dependency ILogic3Dependency) func(
	inputCh <-chan bool,
) (<-chan string, <-chan error) {
	return func(inputCh <-chan bool) (<-chan string, <-chan error) {
		outputCh := make(chan string)
		errCh := make(chan error)
		input := <-inputCh

		go func() {
			output, err := dependency.TransformLogic3(input)

			if err != nil {
				errCh <- err
			}

			outputCh <- output
		}()

		return outputCh, errCh
	}
}
