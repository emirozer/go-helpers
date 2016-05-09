package helpers

// StringInSlice verifies if a string is present in a slice
func StringInSlice(str string, slice []string) bool {
	for _, item := range slice {
		if item == str {
			return true
		}
	}
	return false
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
