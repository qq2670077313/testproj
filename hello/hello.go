package hello

// Greet ... Greet GitHub Actions
func Greet() string {
	return "Hello GitHub Actions"
}

// Het ... make a test
// notest
func Het() string {
	return "Hello GitHub Actions"
}

// Get ... make a test
func Get() string {
	i := 10
	if true {
		// notest
		i++
	} else {
		// notest
		i--
		i++
	}
	i = i + 1
	return "Hello GitHub Actions"
}
