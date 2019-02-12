package pipesample

// ILogic3Dependency serves as dependency for logic 3.
type ILogic3Dependency interface {
	TransformLogic3(input bool) (string, error)
}

// Logic3 performs logic 3.
func Logic3(dependency ILogic3Dependency) Composable {
	return func(input interface{}) (interface{}, error) {
		cast, ok := input.(bool)

		if ok {
			return dependency.TransformLogic3(cast)
		}

		return nil, castError
	}
}
