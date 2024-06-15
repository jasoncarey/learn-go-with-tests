Instead of creating methods of the same name but different signature (like in Java)
```go
func Area(circle Circle) float64 {}
func Area(rectangle Rectangle) float64 {}
```

We define methods on the structs (types)
```go
func (c Circle) Area() float64 {}
func (r Rectangle) Area() float64 {}
```

Interface resolution is implicit, so long as the type passed in matches what the interface is asking for, it will compile. That's why we can say
```go
type Shape interface {
    Area() float64
}
```
because all our shapes have an `Area()` method. The interface is *decoupled* from the concrete types. 