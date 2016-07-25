package helpers

// Simple Panic on error
func Check(e error) {
	if e != nil {
		panic(e)
	}
}
