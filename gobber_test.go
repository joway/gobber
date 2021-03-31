package gobber_test

import (
	"github.com/stretchr/testify/assert"
	"gobber"
	"testing"
	"unsafe"
)

type testStruct struct {
	string   string
	int      int
	pointer  *testStruct
	function func() int
}

type testcase struct {
	fieldName     string
	fieldValue    interface{}
	fieldNewValue interface{}
}

func initField(obj *testStruct, fieldName string, fieldValue interface{}) {
	switch fieldName {
	case "string":
		obj.string = fieldValue.(string)
	case "int":
		obj.int = fieldValue.(int)
	case "pointer":
		obj.pointer = fieldValue.(*testStruct)
	case "function":
		obj.function = fieldValue.(func() int)
	}
}

func equalField(t *testing.T, obj *testStruct, fieldName string, gotPtr unsafe.Pointer) {
	var exceptPtr unsafe.Pointer
	switch fieldName {
	case "string":
		exceptPtr = unsafe.Pointer(&obj.string)
	case "int":
		exceptPtr = unsafe.Pointer(&obj.int)
	case "pointer":
		exceptPtr = unsafe.Pointer(&obj.pointer)
	case "function":
		exceptPtr = unsafe.Pointer(&obj.function)
	}
	assert.Equal(t, exceptPtr, gotPtr)
}

func TestGobber(t *testing.T) {
	testcases := []testcase{
		{"string", "old", "new"},
		{"int", 1, -1},
		{"pointer", &testStruct{int: 1}, &testStruct{int: -1}},
		{"function", func() int { return 1 }, func() int { return -1 }},
	}

	robber := gobber.New(testStruct{})
	for _, tc := range testcases {
		obj := testStruct{}
		//init
		initField(&obj, tc.fieldName, tc.fieldValue)

		//get
		valuePtr := robber.Get(&obj, tc.fieldName)
		equalField(t, &obj, tc.fieldName, valuePtr)

		//set new value
		robber.Set(&obj, tc.fieldName, tc.fieldNewValue)

		//get new value
		valuePtr = robber.Get(&obj, tc.fieldName)
		equalField(t, &obj, tc.fieldName, valuePtr)
	}
}

func TestGobberSetFunc(t *testing.T) {
	robber := gobber.New(testStruct{})
	obj := testStruct{
		function: func() int {
			return 1
		},
	}
	robber.Set(&obj, "function", func() int { return -1 })
	assert.Equal(t, -1, obj.function())

	fptr := robber.Get(&obj, "function")
	function := *(*func() int)(fptr)
	assert.Equal(t, -1, function())
}

func TestGobberSetWrongType(t *testing.T) {
	robber := gobber.New(testStruct{})
	obj := testStruct{
		int: 1,
		function: func() int {
			return 1
		},
	}

	ok := robber.Set(&obj, "int", int64(1))
	assert.False(t, ok)
	ok = robber.Set(&obj, "function", func() int32 { return -1 })
	assert.False(t, ok)
}
