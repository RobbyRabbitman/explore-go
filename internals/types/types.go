package types

import (
	"fmt"
)

// Basic Types
// bool
// string
// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr
// byte // alias for uint8
// rune // alias for int32 - represents a Unicode code point
// float32 float64
// complex64 complex128

// Custom Types
// struct

// Pointer Types
// *<Type>

// interface/any ???

// Array
// Slice

// Map

func PrintZeroValues() {
	var b bool
	var i int
	var str string
	var f float32
	var ui uint
	var anyDeclaration any // ???
	var emptyStruct struct{}
	var nonEmptyStruct struct {
		int
		bool
		string
	}
	var intP *int
	var anyP *any
	var emptyMap map[string]any
	// different types [3]<T> and [5]<T>
	var array3any [3]any
	var array5int [5]int
	var array5any [5]any
	var sliceAny []any
	var sliceInt []int
	fmt.Println("Zero Values")
	for _, v := range []any{b, i, str, f, ui, anyDeclaration, emptyStruct, nonEmptyStruct, intP, anyP, emptyMap, array3any, array5any, array5int, sliceAny, sliceInt} {
		fmt.Printf("Type: %T\nValue: %+v\nValue in Go-syntax: %#v\nvar x %T <=> var x := %#v\n\n", v, v, v, v, v)
	}
}
