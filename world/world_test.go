package world

import "testing"

func TestGreet(t *testing.T) {
	result := Greet()
	if result != "Hello GitHub Actions" {
		t.Errorf("Greet() = %s; Expected Hello GitHub actions", result)
	}
}

func TestErrorExclude(t *testing.T) {
	tests := []int{-1, 0, 1}
	for i := 0; i < len(tests); i++ {
		err := ErrorExclude(tests[i])
		if err != nil {
			t.Log(err)
		}
	}
}

func TestErrorExcludeReturn(t *testing.T) {
	err := ErrorExcludeReturn(0)
	if err != nil {
		t.Log(err)
	}
}
