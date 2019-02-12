package pipesample

// Logic2Dependency serves as dependency for logic 2.
type Logic2Dependency interface {
	TransformLogic2(input uint) (float32, error)
}

// Logic2 performs logic 2.
func Logic2(dependency Logic2Dependency) func(
	inputCh <-chan uint,
) (<-chan float32, <-chan error) {
	return func(inputCh <-chan uint) (<-chan float32, <-chan error) {
		outputCh := make(chan float32)
		errCh := make(chan error)
		input := <-inputCh

		go func() {
			output, err := dependency.TransformLogic2(input)

			if err != nil {
				errCh <- err
			}

			outputCh <- output
		}()

		return outputCh, errCh
	}
}
