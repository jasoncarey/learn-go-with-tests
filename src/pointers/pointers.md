If a symbol (var, type, func, etc) starts with a lowercase symbol, then it is private *outside the pacakge it's defined in*.

When you call a function or method, the arguments are *copied*. In cases where we are manipulating the input variable, we need to point to it so we aren't referring to that copy

```go
// wrong, this won't point to the original balance
func (w Wallet) Deposit () {}

// correct
func (w *Wallet) Deposit() {}
```

Struct pointers are automatically dereferenced.

Linter: `go install github.com/kisielk/errcheck@latest`

Usage: `errcheck .`

If a function returns a pointer, check if it's `nil` or it might raise a `runtime exception` that the compiler won't help with