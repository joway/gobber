package gobber_test

import (
	"gobber"
	"testing"
)

func newBenchStruct() testStruct {
	return testStruct{
		int: 1,
	}
}

func BenchmarkDirectlyGet(b *testing.B) {
	obj := newBenchStruct()
	for i := 0; i < b.N; i++ {
		_ = obj.int
	}
}

func BenchmarkGobberGet(b *testing.B) {
	robber := gobber.New(testStruct{})
	obj := newBenchStruct()
	for i := 0; i < b.N; i++ {
		ptr := robber.Get(&obj, "int")
		addr := (*int)(ptr)
		value := *addr
		if obj.int != value || &obj.int != addr {
			b.Fail()
		}
	}
}

func BenchmarkDirectlySet(b *testing.B) {
	obj := newBenchStruct()
	for i := 0; i < b.N; i++ {
		obj.int = i
	}
}

func BenchmarkGobberSet(b *testing.B) {
	robber := gobber.New(testStruct{})
	obj := newBenchStruct()
	for i := 0; i < b.N; i++ {
		success := robber.Set(&obj, "int", i)
		if !success || obj.int != i {
			b.Fail()
		}
	}
}
