# Gobber

A Golang helper to access private struct field.

## Usage

```go
type Car struct {
	name string
}

//set private field
gobber := gobber.New(Car{})
ok := gobber.Set("name", "")

//get private field
namePtr := gobber.Get("name")
name := *(*string)(namePtr)
```
