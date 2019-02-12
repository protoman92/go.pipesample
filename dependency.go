package pipesample

import "strconv"

// IDependency serves as dependency for the whole app.
type IDependency interface {
	ILogic1Dependency
	ILogic2Dependency
	ILogic3Dependency
}

type dependency struct{}

func (d *dependency) TransformLogic1(input string) (int, error) {
	return strconv.Atoi(input)
}

func (d *dependency) TransformLogic2(input uint) (float32, error) {
	return float32(input), nil
}

func (d *dependency) TransformLogic3(input bool) (string, error) {
	return strconv.FormatBool(input), nil
}

// CreateDependency creates a dependency instance.
func CreateDependency() IDependency {
	return &dependency{}
}
