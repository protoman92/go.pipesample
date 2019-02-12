package pipesample

import "math"

// Logic1ToLogic2Adapter converts output from Logic1 to input for Logic2.
func Logic1ToLogic2Adapter() Composable {
	return func(input interface{}) (interface{}, error) {
		cast, ok := input.(int)

		if ok {
			return uint(math.Abs(float64(cast)) + 1), nil
		}

		return nil, CastError("Logic 1 to Logic 2", input)
	}
}

// Logic2ToLogic3Adapter converts output from Logic1 to input for Logic2.
func Logic2ToLogic3Adapter() Composable {
	return func(input interface{}) (interface{}, error) {
		cast, ok := input.(float32)

		if ok {
			return int(cast)%2 == 0, nil
		}

		return nil, CastError("Logic 2 to Logic 3", input)
	}
}
