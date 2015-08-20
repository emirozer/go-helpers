package helpers

import (
	"strings"
	"testing"
)

const (
	text = "A collection of helper functions about anything and everything"
)

// Tests for the StringInSlice function
func Test_StringInSlice(t *testing.T) {
	slice := strings.Split(text, " ")
	in := "helper"
	notIn := "potato"

	// Should return true
	if result := StringInSlice(in, slice); !result {
		t.Errorf("Expected result: true, got %t instead\n", result)
	}

	// Should return false
	if result := StringInSlice(notIn, slice); result {
		t.Errorf("Expected result: false, got %t instead\n", result)
	}
}

// Benchmark for the ReverseStr function
func Benchmark_ReverseStr(b *testing.B) {
	str := "this is the string to reverse"
	for n := 0; n < b.N; n++ {
		ReverseStr(str)
	}
	b.StopTimer()
}

// Benchmark for the AllUniqueWord function
func Benchmark_AllUniqueWord(b *testing.B) {
	str := "this is the string to test"
	for n := 0; n < b.N; n++ {
		AllUniqueWord(str)
	}
	b.StopTimer()
}

// Benchmark for the StringInSlice function
func Benchmark_StringInSlice(b *testing.B) {
	slice := strings.Split(text, " ")
	str := "potato"
	for n := 0; n < b.N; n++ {
		StringInSlice(str, slice)
	}
	b.StopTimer()
}
