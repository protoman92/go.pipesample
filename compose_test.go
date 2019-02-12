package pipesample

import (
	"testing"
)

func Test_ComposeFunctionsWithoutAdapters_ShouldFail(t *testing.T) {
	// Setup
	dependency := CreateDependency()

	// When
	_, err := Compose([]Composable{
		Logic1(dependency),
		Logic2(dependency),
		Logic3(dependency),
	})("1")

	// Then
	if err == nil {
		t.Errorf("Should throw %v", castError)
	}
}
