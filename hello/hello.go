package hello

// Greet ... Greet GitHub Actions
func Greet() string {
	i := 0
	if i == 1 {
		// notest
		i = 2
		return "111"
	} else {
		i = 3
	}
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
