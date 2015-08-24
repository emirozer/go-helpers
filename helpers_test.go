package helpers

import (
	// "fmt"
	"strings"
	"testing"
)

const (
	separator         = "%20"
	text              = "A collection of helper functions     about anything   and  everything"
	textSHA1          = "e181303567270f6ee37444a21e568039d88b2267"
	textWithSeparator = "A%20collection%20of%20helper%20functions%20about%20anything%20and%20everything"
	textReversed      = "gnihtyreve  dna   gnihtyna tuoba     snoitcnuf repleh fo noitcelloc A"
	uniqueWord        = "abcdef ghijklmn"
	notUniqueWord     = "abcdef ghijkklmn"
	lowerWordInput    = "galatasaray"
	lowerWord         = "Galatasaray"
	upperWordsInput	  = "lorem ipsum dolor sit amet."
	upperWords 		  = "Lorem Ipsum Dolor Sit Amet."
)

var (
	sliceIntX     = []int{46, 77, 89, 43, 3, 33, 0, 10, 95}
	sliceIntY     = []int{5, 93, 54, 10, 48, 66, 97, 95, 46, 11, 8, 56, 72}
	sliceIntDiffX = []int{77, 89, 43, 3, 33, 0}
	sliceIntDiffY = []int{5, 93, 54, 48, 66, 97, 11, 8, 56, 72}
)

// --- --- --- Tests

// Tests for the AllUniqueWord function
func Test_AllUniqueWord(t *testing.T) {
	// Should return true
	if result := AllUniqueWord(uniqueWord); !result {
		t.Errorf("Expected result: true, got %t instead\n", result)
	}
	// Should return false
	if result := AllUniqueWord(notUniqueWord); result {
		t.Errorf("Expected result: false, got %t instead\n", result)
	}
}

// Tests for the ReverseStr function
func Test_ReverseStr(t *testing.T) {
	if result := ReverseStr(text); result != textReversed {
		t.Errorf("Expected result: %q, got %q instead\n", textReversed, result)
	}
}

// Tests for the StrPermutation function
func Test_StrPermutation(t *testing.T) {
	// Should return true
	if result := StrPermutation(text, text); !result {
		t.Errorf("Expected result: true, got %t instead\n", result)
	}
	// Should return false
	if result := StrPermutation(text, text+"some more text"); result {
		t.Errorf("Expected result: false, got %t instead\n", result)
	}
}

// Tests for the ReplaceSpacesWSymb function
func Test_ReplaceSpacesWSymb(t *testing.T) {
	if result := ReplaceSpacesWSymb(text, separator); result != textWithSeparator {
		t.Errorf("Expected result: %q, got %q instead\n", textWithSeparator, result)
	}
}

// TODO
// func Test_PseudoUUID(t *testing.T) {}

func Test_RandomString(t *testing.T) {
	size := 50
	var result string
	if result = RandomString(size); len(result) != size {
		t.Errorf("Expected result size: %d, got %d instead (result: %q)\n", size, len(result), result)
	}
	if RandomString(size) == result {
		t.Error("The strings returned are not random, got two time the same one\n")
	}
}

func Test_RandInt(t *testing.T) {
	result := RandInt(10, 30)

	if result < 10 || result > 30 {
		t.Error("The generated integer is out of min, max bounds")
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

// Tests for the SHA1hash function
func Test_SHA1hash(t *testing.T) {
	if result := SHA1hash(text); result != textSHA1 {
		t.Errorf("Expected result: %q, got %q instead\n", textSHA1, result)
	}
}

// Tests for the FileExists function
func Test_FileExists(t *testing.T) {
	if found := FileExists("./helpers.go"); !found {
		t.Errorf("Expected result: true, got %t instead\n", found)
	}
	if found := FileExists("./non_existent_file.go"); found {
		t.Errorf("Expected result: true, got %t instead\n", found)
	}
}

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

// Tests for the UpperCaseFirst function
func Test_UpperCaseFirst(t *testing.T) {
	if result := UpperCaseFirst(lowerWordInput); result != lowerWord {
		t.Errorf("Expected result: %q, got %q instead\n", lowerWord, result)
	}
}

// Tests for the UpperCaseWords function
func Test_UpperCaseWords(t *testing.T) {
	if result := UpperCaseWords(upperWordsInput); result != upperWords {
		t.Errorf("Expected result: %q, got %q instead\n", upperWords, result)
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

// Benchmark for the ReplaceSpacesWSymb function
func Benchmark_ReplaceSpacesWSymb(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ReplaceSpacesWSymb(text, separator)
	}
	b.StopTimer()
}

// Benchmark for the StrPermutation function
func Benchmark_StrPermutation(b *testing.B) {
	for n := 0; n < b.N; n++ {
		StrPermutation(text, text)
	}
	b.StopTimer()
}

// Benchmark for the StrPermutation function
func Benchmark_RandomString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RandomString(50)
	}
	b.StopTimer()
}

// Benchmark for the UpperCaseFirst function
func Benchmark_UpperCaseFirst(b *testing.B) {
	for n := 0; n < b.N; n++ {
		UpperCaseFirst(lowerWordInput)
	}
	b.StopTimer()
}

// Benchmark for the UpperCaseFirst function
func Benchmark_UpperCaseWords(b *testing.B) {
	for n := 0; n < b.N; n++ {
		UpperCaseFirst(upperWordsInput)
	}
	b.StopTimer()
}
