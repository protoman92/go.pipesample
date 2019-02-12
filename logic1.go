package pipesample

// ILogic1Dependency serves as dependency for logic 1.
type ILogic1Dependency interface {
	TransformLogic1(input string) (int, error)
}

// Logic1 performs logic 1.
func Logic1(dependency ILogic1Dependency) func(string) (int, error) {
	return func(input string) (int, error) {
		return dependency.TransformLogic1(input)
	}
}
