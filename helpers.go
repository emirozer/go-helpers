package helpers

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	rnd "math/rand"
	"net"
	"os"
	"reflect"
	"regexp"
	"strings"
	"time"
	"unicode"
)

// AllUniqueWord will return a bool value based on
// the string containing only unique characters or not.
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

// StrPermutation checks if two strings have the same characters in the same amount
func StrPermutation(a string, b string) bool {
	ma := make(map[rune]int)
	mb := make(map[rune]int)

	for _, v := range a {
		ma[v]++
	}
	for _, v := range b {
		mb[v]++
	}
	return reflect.DeepEqual(ma, mb)
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

// RandomString returns a random string of size l
// Source: https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang
func RandomString(l int) string {
	const (
		letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		size    = int64(len(letters))
	)
	rnd.Seed(time.Now().UnixNano())
	bytes := make([]byte, l)
	for i := range bytes {
		bytes[i] = letters[rnd.Int63()%size]
	}
	return string(bytes)
}

// RandInt returns a random integer between the min and max you set
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

// UpperCaseFirst changes a string's first character to uppercase
func UpperCaseFirst(s string) string {
	phrase := []rune(s)
	phrase[0] = unicode.ToUpper(phrase[0])
	return string(phrase)
}

// UpperCaseWords changes all words' first character to uppercase
func UpperCaseWords(s string) string {
	var stack []string
	r := regexp.MustCompile("[^\\s]+")
	words := r.FindAllString(s, -1)

	for index := range words {
		first := strings.ToUpper(string(words[index][0]))
		rest := words[index][1:]
		stack = append(stack, first+rest)
	}

	return strings.Join(stack, " ")
}

// GetLocalIpv4 returns a string of the host machine local ipv4
func GetLocalIpv4() string {
	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			return fmt.Sprintf("%s", ipv4)
		}
	}
	return "localhost"
}

// RemoveDuplicatesFromIntSlice receives a slice of integers and iterates through them
// eliminating duplicate items in the slice.
func RemoveDuplicatesFromIntSlice(a []int) []int {
	result := []int{}
	seen := map[int]int{}
	for _, val := range a {
		if _, ok := seen[val]; !ok {
			result = append(result, val)
			seen[val] = val
		}
	}
	return result
}
