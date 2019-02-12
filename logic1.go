package pipesample

// ILogic1Dependency serves as dependency for logic 1.
type ILogic1Dependency interface {
	TransformLogic1(input string) (int, error)
}

// Logic1 performs logic 1.
func Logic1(dependency ILogic1Dependency) Composable {
	return func(input interface{}) (interface{}, error) {
		cast, ok := input.(string)

		if ok {
			return dependency.TransformLogic1(cast)
		}

		return nil, castError
	}
}
