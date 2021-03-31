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

## Benchmark

```
BenchmarkDirectlyGet-8   	1000000000	         0.286 ns/op
BenchmarkGobberGet-8     	100000000	        10.2 ns/op
BenchmarkDirectlySet-8   	1000000000	         0.285 ns/op
BenchmarkGobberSet-8     	19638064	        57.3 ns/op
```
