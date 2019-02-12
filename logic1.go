package pipesample

// Logic1Dependency serves as dependency for logic 1.
type Logic1Dependency interface {
	Transform(input string) (int, error)
}

// Logic1Request serves as a request object for logic 1.
type Logic1Request struct {
	input string
	err   error
}

// Logic1Response serves as a response object for logic 1.
type Logic1Response struct {
	output int
	err    error
}

// Logic1 performs logic 1.
func Logic1(
	inputCh <-chan Logic1Request,
	dependency Logic1Dependency,
) <-chan Logic1Response {
	outputCh := make(chan Logic1Response)
	request := <-inputCh

	go func() {
		errorOutput := -1

		if request.err != nil {
			outputCh <- Logic1Response{output: errorOutput, err: request.err}
		}

		output, err := dependency.Transform(request.input)

		if err != nil {
			outputCh <- Logic1Response{output: errorOutput, err: err}
		}

		outputCh <- Logic1Response{output: output, err: nil}
	}()

	return outputCh
}
