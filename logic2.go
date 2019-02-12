package pipesample

// ILogic2Dependency serves as dependency for logic 2.
type ILogic2Dependency interface {
	TransformLogic2(input uint) (float32, error)
}

// Logic2 performs logic 2.
func Logic2(dependency ILogic2Dependency) Composable {
	return func(input interface{}) (interface{}, error) {
		cast, ok := input.(uint)

		if ok {
			return dependency.TransformLogic2(cast)
		}

		return nil, CastError("Logic 2", input)
	}
}
