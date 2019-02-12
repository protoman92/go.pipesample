package pipesample

// ILogic3Dependency serves as dependency for logic 3.
type ILogic3Dependency interface {
	TransformLogic3(input bool) (string, error)
}

// Logic3 performs logic 3.
func Logic3(dependency ILogic3Dependency) func(bool) (string, error) {
	return func(input bool) (string, error) {
		return dependency.TransformLogic3(input)
	}
}
