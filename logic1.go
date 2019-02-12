package pipesample

// Logic1Dependency serves as dependency for logic 1.
type Logic1Dependency interface {
	TransformLogic1(input string) (int, error)
}

// Logic1 performs logic 1.
func Logic1(dependency Logic1Dependency) func(
	inputCh <-chan string,
) (<-chan int, <-chan error) {
	return func(inputCh <-chan string) (<-chan int, <-chan error) {
		outputCh := make(chan int)
		errCh := make(chan error)
		input := <-inputCh

		go func() {
			output, err := dependency.TransformLogic1(input)

			if err != nil {
				errCh <- err
			}

			outputCh <- output
		}()

		return outputCh, errCh
	}
}
