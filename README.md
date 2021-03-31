# Gobber

A Golang helper to access private struct field.

## Usage

```go
type Car struct {
    name string
}

//init
gobber := gobber.New(Car{}) //Car{} just for gobber to setup struct metadata

//set private field
target := &Car{name: "old"}
ok := gobber.Set(target, "name", "new")

//get private field
namePtr := gobber.Get("name")
name := *(*string)(namePtr) //"new"
```

## How it works

```text
(address of the object) + (offset of the field in struct) == (address of the field)

//Get
*(*Type)(address of the field)

//Set
*(*Type)(address of the field) = *(*Type)(address of the new value)
```

## Benchmark

```
BenchmarkDirectlyGet-8   	1000000000	         0.286 ns/op
BenchmarkGobberGet-8     	100000000	        10.2 ns/op
BenchmarkDirectlySet-8   	1000000000	         0.285 ns/op
BenchmarkGobberSet-8     	19638064	        57.3 ns/op
```
