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
		t.Error("Should have thrown cast error")
	}
}

func Test_ComposeFunctionsWithAdapters_ShouldSucceed(t *testing.T) {
	// Setup
	dependency := CreateDependency()

	// When
	output, err := Compose([]Composable{
		Trace(dependency, "Logic1")(Logic1(dependency)),
		Logic1ToLogic2Adapter(),
		Trace(dependency, "Logic2")(Logic2(dependency)),
		Logic2ToLogic3Adapter(),
		Trace(dependency, "Logic3")(Logic3(dependency)),
	})("1")

	// Then
	if err != nil {
		t.Errorf("Should not have thrown %v", err)
	}

	if output != "true" {
		t.Errorf("Should not have returned %v", output)
	}
}
