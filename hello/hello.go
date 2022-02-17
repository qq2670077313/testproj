package hello

// Greet ... Greet GitHub Actions
func Greet() string {
	return "Hello GitHub Actions"
}

// Het ... make a test
//nolint
func Het() string {
	return "Hello GitHub Actions"
}

// Get ... make a test
func Get() string {
	i := 10
	if true {
		i++
		//nolint
	} else {
		i--
		i += 1
	}
	return "Hello GitHub Actions"
}
