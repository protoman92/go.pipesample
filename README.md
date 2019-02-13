# go.pipesample

[![Go Report Card](https://goreportcard.com/badge/github.com/protoman92/go.pipesample)](https://goreportcard.com/report/github.com/protoman92/go.pipesample)
[![Build Status](https://travis-ci.org/protoman92/go.pipesample.svg?branch=master)](https://travis-ci.org/protoman92/go.pipesample)
[![Coverage Status](https://coveralls.io/repos/github/protoman92/go.pipesample/badge.svg?branch=master)](https://coveralls.io/github/protoman92/go.pipesample?branch=master)

Input-output piping sample that demonstrates composition of functions.

We define a **Composable** as follows:

```go
type Composable = (interface{}) (interface{}, error)
```

Logic functions should all return **Composable**s. They generally should be in this form:

```go
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

    return nil, CastError("Logic 1", input)
  }
}
```

Since the output of one function may not match the input for another, in order to keep logic functions from changing, we also define adapters like so:

```go
func Logic1ToLogic2Adapter() Composable {
  return func(input interface{}) (interface{}, error) {
    cast, ok := input.(int)

    if ok {
      return uint(math.Abs(float64(cast)) + 1), nil
    }

    return nil, CastError("Logic 1 to Logic 2", input)
  }
}
```

Since logic functions and adapters all return **Composable**s, we can use a compose function to chain them:

```go
output, err := Compose(
  Logic1(dependency),
  Logic1ToLogic2Adapter(),
  Logic2(dependency),
  Logic2ToLogic3Adapter(),
  Logic3(dependency),
)("1")
```

And since **Compose** also returns a **Composable**, we can keep chaining them endlessly. The entire chain can then be put in a goroutine for async.

To go even more deeply, we can define **ComposableMapper** to wrap a base function with extra functionalities:

```go
// ComposableMapper represents a Composable converter.
type ComposableMapper = func(Composable) Composable
```

For example, we can define a tracer that tracks time spent on invoking a function:

```go
// ILogger does logging.
type ILogger interface {
  Log(event interface{})
}

// Trace does some tracing for a function with a specified name.
func Trace(logger ILogger, funcName string) ComposableMapper
  return func(composable Composable) Composable {
    return func(input interface{}) (interface{}, error) {
      startTime := time.Now()

      defer func() {
        elapsed := time.Now().Sub(startTime)
        logger.Log(fmt.Sprintf("%v took %v millis to run", funcName, elapsed))
      }()

      return composable(input)
    }
  }
}
```

Then in **Compose**:

```go
output, err := Compose(
  Trace(dependency, "Logic1")(Logic1(dependency)),
  Logic1ToLogic2Adapter(),
  Trace(dependency, "Logic2")(Logic2(dependency)),
  Logic2ToLogic3Adapter(),
  Trace(dependency, "Logic3")(Logic3(dependency)),
)("1")
```

We can also define **ComposeMapper** to chain **ComposableMapper** together like we did with **Composable**:

```go
// When
output, err := Compose(
  ComposeMapper(Trace(dependency, "Logic1"), commonMapper)(Logic1(dependency)),
  Logic1ToLogic2Adapter(),
  ComposeMapper(Trace(dependency, "Logic2"), commonMapper)(Logic2(dependency)),
  Logic2ToLogic3Adapter(),
  ComposeMapper(Trace(dependency, "Logic3"), commonMapper)(Logic3(dependency)),
)("1")
```
