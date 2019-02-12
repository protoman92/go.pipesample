# go.pipesample

Request-response piping sample that demonstrates composition of functions.

We define a **Composable** as follows:

```go
type Composable = (interface{}) (interface{}, error)
```

Logic functions should all return **Composable**s. They generally should be in this form:

```go
type IDependency interface {
  TransformLogic(input string) (int, error)
}

// Logic performs some logic. Use DI here to keep input/output consistent.
func Logic(dependency IDependency) Composable {
  return func(input interface{}) (interface{}, error) {
    cast, ok := input.(string)

    if ok {
      return dependency.TransformLogic(cast)
    }

    return nil, errors.New("Cast error!")
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
output, err := Compose([]Composable{
  Logic1(dependency),
  Logic1ToLogic2Adapter(),
  Logic2(dependency),
  Logic2ToLogic3Adapter(),
  Logic3(dependency),
})("1")
```

And since **Compose** also returns a **Composable**, we can keep chaining them endlessly. The entire chain can then be put in a goroutine for async.
