Never initialize an empty map variable:

```go
var m map[string]string

// should be
var m = map[string]string{}
// or
var m = make(map[string]string)
```