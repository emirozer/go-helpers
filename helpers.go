package helpers

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	rnd "math/rand"
	"os"
	"reflect"
	"strings"
	"time"
)

// AllUniqueWord will return a bool value based on if the word contains
// unique characters or not.
func AllUniqueWord(str string) bool {
	result := make(map[rune]int)
	for _, v := range str {
		result[v]++
	}
	for _, v := range result {
		if v > 1 {
			return false
		}
	}
	return true
}

// ReverseStr returns the string reversed rune-wise left to right
// Source: https://github.com/golang/example/tree/master/stringutil
func ReverseStr(s string) (result string) {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

//StrPermutation checks if two strings has the same elements in the same amount
func StrPermutation(a string, b string) bool {
	ma := make(map[string]int)
	mb := make(map[string]int)
	aSl := strings.Split(a, "")
	bSl := strings.Split(b, "")

	for _, v := range aSl {
		ma[v]++
	}
	for _, v := range bSl {
		mb[v]++
	}

	eq := reflect.DeepEqual(ma, mb)

	if eq {
		return true
	}

	return false

}

// ReplaceSpacesWSymb expects two strings, first arg being the string itself
// Second argument is the symbol to replace spaces
// e.g. s := "emir ozer" symb := "%20"  / Result will be "emir%20ozer"
func ReplaceSpacesWSymb(s string, symb string) (result string) {
	return strings.Join(strings.Fields(s), symb)
}

//PseudoUUID generates a uuid like string
func PseudoUUID() (uuid string) {

	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	uuid = fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

	return
}

//RandomString gets the desired length and returns a random string
func RandomString(l int) string {
	rnd.Seed(time.Now().UnixNano())
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(RandInt(65, 90))
	}
	return string(bytes)
}

//RandInt returns a random integer between the min and max you set
func RandInt(min int, max int) int {
	rnd.Seed(time.Now().UnixNano())
	return min + rnd.Intn(max-min)
}

// DiffSlices takes two integer slices and returns the diff between them
// e.g. DiffSlices(X, Y) will return the elements that are only in X
// If an integer is present in both slices but not the same number of time,
// it will be reflected in the result
func DiffSlices(X, Y []int) (ret []int) {
	m := make(map[int]int)

	for _, y := range Y {
		m[y]++
	}

	for _, x := range X {
		if m[x] > 0 {
			m[x]--
			continue
		}
		ret = append(ret, x)
	}

	return ret
}

// SHA1hash calculates a SHA1 given a string
func SHA1hash(cad string) string {
	data := sha1.Sum([]byte(cad))
	return hex.EncodeToString(data[:])
}

//FileExists verifies if a filepath exists
func FileExists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

// StringInSlice verifies if a string is present in a slice
func StringInSlice(str string, slice []string) bool {
	for _, item := range slice {
		if item == str {
			return true
		}
	}
	return false
}
