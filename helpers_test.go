package helpers

import (
	// "fmt"
	"strings"
	"testing"
)

const (
	text     = "A collection of helper functions about anything and everything"
	textSHA1 = "037b993cade2a9493c8ac2f9766c6ec5889311f5"
)

var (
	sliceIntX     = []int{46, 77, 89, 43, 3, 33, 0, 10, 95}
	sliceIntY     = []int{5, 93, 54, 10, 48, 66, 97, 95, 46, 11, 8, 56, 72}
	sliceIntDiffX = []int{77, 89, 43, 3, 33, 0}
	sliceIntDiffY = []int{5, 93, 54, 48, 66, 97, 11, 8, 56, 72}
)

// --- --- --- Tests

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

// Tests for the SHA1hash function
func Test_SHA1hash(t *testing.T) {
	if result := SHA1hash(text); result != textSHA1 {
		t.Errorf("Expected result: %q, got %q instead\n", textSHA1, result)
	}
}

// Tests for the DiffSlices function
func Test_DiffSlices(t *testing.T) {
	result := DiffSlices(sliceIntX, sliceIntY)
	if len(sliceIntDiffX) != len(result) {
		t.Logf("Result size (%d) is different from the expected result size (%d)\n", len(sliceIntDiffX), len(result))
		t.FailNow()
	}
	for i, value := range result {
		if value != sliceIntDiffX[i] {
			t.Logf("Expected result at index %d: \"%d\", got \"%d\" instead\n", i, sliceIntDiffX[i], value)
			t.FailNow()
		}
	}
	result = DiffSlices(sliceIntY, sliceIntX)
	if len(sliceIntDiffY) != len(result) {
		t.Logf("Result size (%d) is different from the expected result size (%d)\n", len(sliceIntDiffY), len(result))
		t.FailNow()
	}
	for i, value := range result {
		if value != sliceIntDiffY[i] {
			t.Logf("Expected result at index %d: \"%d\", got \"%d\" instead\n", i, sliceIntDiffY[i], value)
			t.FailNow()
		}
	}
}

// --- --- --- Benchmarks

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

// Benchmark for the SHA1hash function
func Benchmark_SHA1hash(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SHA1hash(text)
	}
	b.StopTimer()
}

// Benchmark for the DiffSlices function
func Benchmark_DiffSlices(b *testing.B) {
	for n := 0; n < b.N; n++ {
		DiffSlices(sliceIntX, sliceIntY)
	}
	b.StopTimer()
}
